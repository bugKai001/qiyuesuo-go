package request

import (
	"github.com/bugKai001/qiyuesuo/http"
)

type OpenCompleteReportDownloadRequest struct {
	ProveId string `json:"proveId,omitempty"`
}

func (obj OpenCompleteReportDownloadRequest) GetUrl() string {
	return "/chain/evidence/complete/report/download"
}

func (obj OpenCompleteReportDownloadRequest) GetHttpParameter() *http.HttpParameter {
	parameter := http.NewGetHttpParameter()
	parameter.AddParam("proveId", obj.ProveId)
	return parameter
}
