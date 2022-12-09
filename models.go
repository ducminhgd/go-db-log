package models

import (
	"time"

	"gorm.io/gorm"
)

type EventLog struct {
	ID         string    `gorm:"primarykey,column:id" json:"id"`
	EventType  string    `gorm:"column:event_type" json:"eventType"`
	ObjectType string    `gorm:"column:object_type" json:"objectType"`
	ObjectID   string    `gorm:"column:object_id" json:"objectID"`
	ActorType  string    `gorm:"column:actor_type" json:"actorType"`
	ActorID    string    `gorm:"column:actor_id" json:"actorID"`
	Data       string    `gorm:"column:data" json:"data"`
	Result     int32     `gorm:"column:result" json:"result"`
	Timestamp  time.Time `gorm:"column:timestamp" json:"timestamp"`
}

type QueryParams struct {
	Limit  int64
	Offset int64
	Result int32
	From   time.Time
	To     time.Time
}

// TableName returns the name of event logging table
func (EventLog) TableName() string {
	return "event_log"
}

type EventLogs []*EventLog

type Clause func(tx *gorm.DB)

func Create(db *gorm.DB, el *EventLog) (*EventLog, error) {
	err := db.Model(&EventLog{}).Create(el).Error
	return el, err
}

func Count(db *gorm.DB, params QueryParams) (int64, error) {
	var count int64
	tx := db.Model(&EventLogs{}).Count(&count)
	return count, tx.err
}
