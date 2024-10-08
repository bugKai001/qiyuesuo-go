package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type SealAutoCreateRequest struct {
	TenantName    string               `json:"tenantName,omitempty"`
	Name          string               `json:"name,omitempty"`
	SealImageInfo *model.SealImageInfo `json:"sealImageInfo,omitempty"`
	Users         []*model.User        `json:"users,omitempty"`
}

func (obj SealAutoCreateRequest) GetUrl() string {
	return "/v2/seal/autocreate"
}

func (obj SealAutoCreateRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	jsonBytes, _ := json.Marshal(obj)
	parameter.SetJsonParamer(string(jsonBytes))
	return parameter
}
