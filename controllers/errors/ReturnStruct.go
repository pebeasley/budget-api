package errors

import "time"

type ReturnStruct struct {
	Status  string
	Time    time.Time
	Message string
	Data    string
}
