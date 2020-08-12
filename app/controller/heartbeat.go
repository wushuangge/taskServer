package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	_struct "taskdash/app/struct"
)

var ServiceHeartBeat map[string]int64

func HeartBeat(data []byte) {
	heartBeat:=_struct.HeartBeat{}
	err:=json.Unmarshal(data, &heartBeat)
	if err != nil{
		log.Error(err)
		return
	}
	ServiceHeartBeat[heartBeat.URL] = heartBeat.Time
}