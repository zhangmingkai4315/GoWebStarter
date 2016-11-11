package controllers

import (
	"net/http"
	"webapp/common"
)

type  Version struct{
	Version string `json:version`
}

func GetApiVersionHandler(w http.ResponseWriter, r *http.Request) {

	response:=Version{"v1.0"}
	common.DisplayAppData(w,response)
	return
}