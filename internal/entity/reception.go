package entity

import "time"

type ReceptionStatus string

const (
	ReceptionStatusInProgress ReceptionStatus = "in_progress"
	ReceptionStatusClosed     ReceptionStatus = "closed"
)

type Reception struct {
	Id            int64
	ReceptionDate time.Time
	OrderPoint    OrderPoint
	Products      []Product
	Status        ReceptionStatus
}
