package model

type User struct {
	ID       int    `json:"-" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" validate:"required|minLen:6|maxLen:16" gorm:"comment:用户名称"`
	Password string `json:"password" validate:"required|minLen:6|maxLen:16" gorm:"comment:用户密码"`
	TimeRecorder
}
