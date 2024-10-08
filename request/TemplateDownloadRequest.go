package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type TemplateDownloadRequest struct {
	// 模板ID
	TemplateId string `json:"templateId,omitempty"`
}

func (obj TemplateDownloadRequest) GetUrl() string {
	return "/v2/template/download"
}

func (obj TemplateDownloadRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("templateId", obj.TemplateId)
	return parameter
}
