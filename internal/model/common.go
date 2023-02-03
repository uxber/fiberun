package model

import "time"

type TimeRecorder struct {
	Created time.Time `json:"-" gorm:"autoCreateTime;comment:创建时间"`
	Updated time.Time `json:"-" gorm:"autoUpdateTime;comment:更新时间"`
}
