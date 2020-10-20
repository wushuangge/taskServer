package controller

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	_struct "taskdash/app/struct"
	"time"
)

var (
	mutex        sync.Mutex
	mapHeartbeat map[string]_struct.ServiceTimer
)

const (
	tickerRatio  = 30
	counterRatio = 3
)

func InitTicker() {
	mapHeartbeat = make(map[string]_struct.ServiceTimer)
	ticker := time.NewTicker(tickerRatio * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				mutex.Lock()
				for k, v := range mapHeartbeat {
					if v.Enable == true {
						if v.Counter > 0 {
							mapHeartbeat[k] = fillTimer(v.Enable, v.Counter-1)
						} else {
							ServiceUpdate(k, false)
							mapHeartbeat[k] = fillTimer(false, v.Counter)
							fmt.Println("timeout!!!", "id is ", k)
						}
					}
				}
				mutex.Unlock()
			}
		}
	}()
}

func HeartBeat(data []byte) {
	heartBeat := _struct.HeartBeat{}
	err := json.Unmarshal(data, &heartBeat)
	if err != nil {
		log.Error(err)
		return
	}
	id := GetMd5String(heartBeat.Network + heartBeat.Address)
	mutex.Lock()
	if mapHeartbeat[id].Enable == false {
		ServiceUpdate(id, true)
	}
	mapHeartbeat[id] = fillTimer(true, counterRatio)
	mutex.Unlock()
}

func fillTimer(enable bool, counter int) _struct.ServiceTimer {
	serviceTimer := _struct.ServiceTimer{
		Enable:  enable,
		Counter: counter,
	}
	return serviceTimer
}
