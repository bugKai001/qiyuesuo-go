package request

import (
	"encoding/json"
	"github.com/bugKai001/qiyuesuo/http"
	"github.com/bugKai001/qiyuesuo/model"
)

type ContractSignUserRequest struct {
	Param *model.UserSignParam
}

func (obj ContractSignUserRequest) GetUrl() string {
	return "/v2/contract/personalsign"
}

func (obj ContractSignUserRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewPostHttpParameter()
	if obj.Param != nil {
		jsonBytes, _ := json.Marshal(obj.Param)
		parameter.SetJsonParamer(string(jsonBytes))
	}
	return parameter
}
