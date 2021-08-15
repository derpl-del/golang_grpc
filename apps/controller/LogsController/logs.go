package LogsController

import (
	"simple-api/apps/models/LogModels"
	"time"
)

//Pipeline to Logging
func CreatePoint(trxid string, point string, service string, payload string) {
	go func(trxid string, point string, service string, payload string) {
		logs := LogModels.Logs{
			TrxId:        trxid,
			LoggingPoint: point,
			ServiceName:  service,
			Payload:      payload,
			Datetime:     time.Now(),
		}
		logs.Create()
	}(trxid, point, service, payload)
}
