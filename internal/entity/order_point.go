package entity

import "time"

type OrderPoint struct {
	Id               int64
	RegistrationDate time.Time
	City             string
}
