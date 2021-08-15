package LogModels

import (
	"simple-api/config/dbadapter"
	"time"
)

//Logs models
type Logs struct {
	TrxId        string    `json:"trxid"`
	LoggingPoint string    `json:"logging_point"`
	ServiceName  string    `json:"service_name"`
	Payload      string    `json:"payload"`
	Datetime     time.Time `json:"datetime"`
}

//function to create connection to db
func OpenConection() dbadapter.Adapter {
	Adapter := dbadapter.Adapter{}.New()
	return Adapter
}

//function to insert to db
func (L *Logs) Create() error {
	adp := OpenConection()
	err := adp.Table.Create(L).Error
	adp.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}
