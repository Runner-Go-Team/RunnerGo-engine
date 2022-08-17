package model

// Plan 计划结构体
type Plan struct {
	PlanID     string      `json:"planId"`   // 计划id
	PlanName   string      `json:"planName"` // 计划名称
	ReportId   string      `json:"reportId"` // 报告名称
	ReportName string      `json:"reportName"`
	ConfigTask *ConfigTask `json:"configTask"` // 任务配置
	Variable   *Variable   `json:"variable"`   // 全局变量
	Scene      *Scene      `json:"scene"`      // 场景列表
}