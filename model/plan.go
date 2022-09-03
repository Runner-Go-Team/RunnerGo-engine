package model

import "sync"

// Plan 计划结构体
type Plan struct {
	PlanId     string      `json:"planId" bson:"planId"`     // 计划id
	PlanName   string      `json:"planName" bson:"planName"` // 计划名称
	ReportId   string      `json:"reportId" bson:"reportId"` // 报告名称
	ReportName string      `json:"reportName" bson:"reportName"`
	ConfigTask *ConfigTask `json:"configTask" bson:"configTask"` // 任务配置
	Variable   *sync.Map   `json:"variable" bson:"variable"`     // 全局变量
	Scene      *Scene      `json:"scene" bson:"scene"`           // 场景
}

// Group 分组
type Group struct {
	Group    *Group    `json:"groups" bson:"groups"`
	Variable *sync.Map `json:"variable"` // 全局变量
	Scene    *Scene    `json:"scene"`    // 场景
}
