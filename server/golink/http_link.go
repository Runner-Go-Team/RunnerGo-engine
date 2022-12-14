// Package golink 连接
package golink

import (
	"RunnerGo-engine/model"
	"RunnerGo-engine/server/client"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// HttpSend 发送http请求
func HttpSend(event model.Event, api model.Api, configuration *model.Configuration, requestCollection *mongo.Collection) (bool, int64, uint64, float64, float64, string) {
	var (
		isSucceed     = true
		errCode       = model.NoError
		receivedBytes = float64(0)
		errMsg        = ""
	)

	resp, req, requestTime, sendBytes, err, str := client.HTTPRequest(api.Method, api.Request.URL, api.Request.Body, api.Request.Query,
		api.Request.Header, api.Request.Auth, api.Timeout)
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	defer fasthttp.ReleaseRequest(req)
	var regex []map[string]interface{}
	if api.Regex != nil {
		for _, regular := range api.Regex {
			if regular.IsChecked == -1 {
				continue
			}
			reg := make(map[string]interface{})
			value := regular.Extract(string(resp.Body()), configuration)
			reg[regular.Var] = value
			regex = append(regex, reg)
		}
	}
	if err != nil {
		isSucceed = false
		errMsg = err.Error()
	}
	if resp.StatusCode() != 200 {
		isSucceed = false
		errMsg = string(resp.Body())
	}
	var assertionMsgList []model.AssertionMsg
	// 断言验证
	if api.Assert != nil {
		var assertionMsg = model.AssertionMsg{}
		var (
			code    = int64(10000)
			succeed = true
			msg     = ""
		)
		for _, v := range api.Assert {
			if v.IsChecked == -1 {
				continue
			}
			code, succeed, msg = v.VerifyAssertionText(resp)
			if succeed != true {
				errCode = code
				isSucceed = succeed
				errMsg = msg
			}
			assertionMsg.Code = code
			assertionMsg.IsSucceed = succeed
			assertionMsg.Msg = msg
			assertionMsgList = append(assertionMsgList, assertionMsg)
		}
	}
	// 接收到的字节长度
	//contentLength = uint(resp.Header.ContentLength())

	receivedBytes = float64(resp.Header.ContentLength()) / 1024
	if receivedBytes <= 0 {
		receivedBytes = float64(len(resp.Body())) / 1024
	}
	// 开启debug模式后，将请求响应信息写入到mongodb中
	if api.Debug != "" && api.Debug != "stop" {
		switch api.Debug {
		case model.All:
			debugMsg := make(map[string]interface{})
			debugMsg["uuid"] = api.Uuid.String()
			debugMsg["event_id"] = event.Id
			debugMsg["api_id"] = api.TargetId
			debugMsg["api_name"] = api.Name
			debugMsg["type"] = model.RequestType
			debugMsg["request_time"] = requestTime / uint64(time.Millisecond)
			debugMsg["request_code"] = resp.StatusCode()
			debugMsg["request_header"] = req.Header.String()
			debugMsg["response_body"] = string(resp.Body())
			if string(req.Body()) != "" {
				debugMsg["request_body"] = string(req.Body())
			} else {
				debugMsg["request_body"] = str
			}
			if string(resp.Body()) == "" && errMsg != "" {
				debugMsg["response_body"] = errMsg
			}

			debugMsg["response_header"] = resp.Header.String()

			debugMsg["response_bytes"] = receivedBytes
			if err != nil {
				debugMsg["response_body"] = err.Error()
			} else {

			}
			switch isSucceed {
			case false:
				debugMsg["status"] = model.Failed
			case true:
				debugMsg["status"] = model.Success
			}

			debugMsg["next_list"] = event.NextList

			if api.Assert != nil {
				debugMsg["assertion"] = assertionMsgList
			}
			if api.Regex != nil {
				debugMsg["regex"] = regex
			}
			if requestCollection != nil {
				debugMsg["report_id"] = event.ReportId
				model.Insert(requestCollection, debugMsg)
			}
		case model.OnlySuccess:
			if isSucceed == true {
				debugMsg := make(map[string]interface{})
				debugMsg["uuid"] = api.Uuid.String()
				debugMsg["event_id"] = event.Id
				debugMsg["api_id"] = api.TargetId
				debugMsg["api_name"] = api.Name
				debugMsg["type"] = model.RequestType
				debugMsg["request_time"] = requestTime / uint64(time.Millisecond)
				debugMsg["request_code"] = resp.StatusCode()
				debugMsg["request_header"] = req.Header.String()
				if string(req.Body()) != "" {
					debugMsg["request_body"] = string(req.Body())
				} else {
					debugMsg["request_body"] = str
				}
				debugMsg["response_header"] = resp.Header.String()
				if string(resp.Body()) == "" && errMsg != "" {
					debugMsg["response_body"] = errMsg
				}
				debugMsg["response_bytes"] = receivedBytes
				debugMsg["status"] = model.Success
				debugMsg["next_list"] = event.NextList
				if err != nil {
					debugMsg["response_body"] = err.Error()
				}
				if api.Assert != nil {
					debugMsg["assertion"] = assertionMsgList
				}
				if api.Regex != nil {
					debugMsg["regex"] = regex
				}
				if requestCollection != nil {
					debugMsg["report_id"] = event.ReportId
					model.Insert(requestCollection, debugMsg)
				}
			}

		case model.OnlyError:
			if isSucceed == false {
				debugMsg := make(map[string]interface{})
				debugMsg["uuid"] = api.Uuid.String()
				debugMsg["event_id"] = event.Id
				debugMsg["api_id"] = api.TargetId
				debugMsg["api_name"] = api.Name
				debugMsg["type"] = model.RequestType
				debugMsg["request_time"] = requestTime / uint64(time.Millisecond)
				debugMsg["request_code"] = resp.StatusCode()
				debugMsg["request_header"] = req.Header.String()
				if string(req.Body()) != "" {
					debugMsg["request_body"] = string(req.Body())
				} else {
					debugMsg["request_body"] = str
				}
				debugMsg["response_header"] = resp.Header.String()
				if string(resp.Body()) == "" && errMsg != "" {
					debugMsg["response_body"] = errMsg
				}

				debugMsg["response_bytes"] = receivedBytes
				debugMsg["status"] = model.Failed
				debugMsg["next_list"] = event.NextList
				if api.Assert != nil {
					debugMsg["assertion"] = assertionMsgList
				}
				if api.Regex != nil {
					debugMsg["regex"] = regex
				}
				if requestCollection != nil {
					debugMsg["report_id"] = event.ReportId
					model.Insert(requestCollection, debugMsg)
				}
			}

		}
	}
	return isSucceed, errCode, requestTime, sendBytes, receivedBytes, errMsg
}
