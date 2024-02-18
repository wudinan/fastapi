// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi/internal/model"
)

type (
	IModelAgent interface {
		// 根据模型代理ID获取模型代理信息
		GetModelAgent(ctx context.Context, id string) (*model.ModelAgent, error)
		// 模型代理列表
		List(ctx context.Context, ids []string) ([]*model.ModelAgent, error)
		// 根据模型代理ID获取密钥列表
		GetModelAgentKeys(ctx context.Context, id string) ([]*model.Key, error)
		// 挑选模型代理
		PickModelAgent(ctx context.Context, m *model.Model) (modelAgent *model.ModelAgent, err error)
		// 移除模型代理
		RemoveModelAgent(ctx context.Context, m *model.Model, modelAgent *model.ModelAgent)
		// 记录错误模型代理
		RecordErrorModelAgent(ctx context.Context, m *model.Model, modelAgent *model.ModelAgent)
		// 挑选模型代理密钥
		PickModelAgentKey(ctx context.Context, modelAgent *model.ModelAgent) (key *model.Key, err error)
		// 移除模型代理密钥
		RemoveModelAgentKey(ctx context.Context, modelAgent *model.ModelAgent, key *model.Key)
		// 记录错误模型代理密钥
		RecordErrorModelAgentKey(ctx context.Context, modelAgent *model.ModelAgent, key *model.Key)
	}
)

var (
	localModelAgent IModelAgent
)

func ModelAgent() IModelAgent {
	if localModelAgent == nil {
		panic("implement not found for interface IModelAgent, forgot register?")
	}
	return localModelAgent
}

func RegisterModelAgent(i IModelAgent) {
	localModelAgent = i
}