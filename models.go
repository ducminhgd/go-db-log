package models

import "time"

type Timestamp time.Time

type EventLog struct {
	ID         string    `gorm:"primarykey,column:name"`
	EventType  string    `gorm:"column:event_type"`
	ObjectType string    `gorm:"column:object_type"`
	ObjectID   string    `gorm:"column:object_id"`
	ActorType  string    `gorm:"column:actor_type"`
	ActorID    string    `gorm:"column:actor_id"`
	Data       string    `gorm:"column:data"`
	Result     int32     `gorm:"column:result"`
	Timestamps Timestamp `gorm:"column:timestamps"`
}

// TableName returns the name of event logging table
func (EventLog) TableName() string {
	return "event_log"
}

type EventLogs []*EventLog
