package execution

import (
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/mongo"
	"kp-runner/config"
	"kp-runner/log"
	"kp-runner/model"
	"kp-runner/server/client"
	"kp-runner/server/golink"
	"sync"
	"time"
)

type RequestTimeData struct {
	PlanId  string           `json:"planId"`
	SceneId string           `json:"sceneId"`
	Apis    []RequestTimeApi `json:"apis"`
}

type RequestTimeApi struct {
	ApiName   string  `json:"apiName"`
	Threshold float64 `json:"threshold"`
	Actual    float64 `json:"actual"`
	Custom    int     `json:"custom"`
}

func GetRequestTime(esClient *elastic.Client, requestTimeData *RequestTimeData) {

}

// ExecutionRTModel 响应时间模式
func ExecutionRTModel(statusCh chan bool, plan *model.Plan, ch chan *model.ResultDataMsg, wg *sync.WaitGroup, requestCollection, responseCollection *mongo.Collection) {
	defer close(ch)
	// 定义一个chan, 从es中获取当前错误率与阈值分别是多少
	requestTimeData := new(RequestTimeData)
	// 连接es，并查询当前错误率为多少，并将其放入到chan中
	err, esClient := client.NewEsClient(config.Config["esHost"].(string))
	if err != nil {
		return
	}
	go GetRequestTime(esClient, requestTimeData)
	startTime := time.Now().Unix()
	// preConcurrent 是为了回退,此功能后续开发
	//preConcurrent := startConcurrent
	startConcurrent := plan.ConfigTask.TestModel.RTTest.StartConcurrent
	length := plan.ConfigTask.TestModel.RTTest.Length
	maxConcurrent := plan.ConfigTask.TestModel.RTTest.MaxConcurrent
	lengthDuration := plan.ConfigTask.TestModel.RTTest.LengthDuration
	stableDuration := plan.ConfigTask.TestModel.RTTest.StableDuration
	timeUp := plan.ConfigTask.TestModel.RTTest.TimeUp
	eventList := plan.Scene.EventList
	concurrent := startConcurrent
	// 只要开始时间+持续时长大于当前时间就继续循环
	for startTime+lengthDuration > time.Now().Unix() {
		_, status := model.QueryPlanStatus(plan.PlanID + ":" + plan.Scene.SceneId + ":" + "status")
		if status == "false" {
			return
		}

		for i := int64(0); i < concurrent; i++ {
			wg.Add(1)
			go func(i, concurrent int64) {
				if plan.Variable == nil {
					plan.Variable = new(sync.Map)
				}
				globalVariable := plan.Variable
				golink.DisposeScene(eventList, ch, plan, globalVariable, wg, requestCollection, responseCollection, i, concurrent)
				wg.Done()
			}(i, concurrent)
			// 如果设置了启动并发时长
			if timeUp != 0 && (startConcurrent/timeUp)%i == 0 && i != 0 {
				time.Sleep(1 * time.Second)
			}
		}
		if concurrent == maxConcurrent && lengthDuration == stableDuration && startTime+int64(lengthDuration) >= time.Now().Unix() {
			log.Logger.Info("计划", plan.PlanName, "结束")
			return
		}
		// 如果当前并发数小于最大并发数，
		if concurrent <= maxConcurrent {
			// 从开始时间算起，加上持续时长。如果大于现在是的时间，说明已经运行了持续时长的时间，那么就要将开始时间的值，变为现在的时间值
			if startTime+lengthDuration >= time.Now().Unix() {
				startTime = time.Now().Unix()
				//preConcurrent = concurrent
				if concurrent+length <= maxConcurrent {
					concurrent = concurrent + length
				} else {
					concurrent = maxConcurrent
					lengthDuration = stableDuration
				}
				apis := requestTimeData.Apis
				for _, api := range apis {
					if api.Threshold < api.Actual {
						log.Logger.Info(api.ApiName, "接口：在", concurrent, "并发时", api.Custom, "线响应时间", api.Actual, "大于所设阈值", api.Threshold)
						log.Logger.Info("计划", plan.PlanName, "结束")
						return

					}
				}

			}
		}

	}

}
