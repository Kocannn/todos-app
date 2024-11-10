package domain

import "time"

type Todo struct {
	Id          int       `json:"id" gorm:"primary_key;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:datetime;not null;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:datetime;not null;autoUpdateTime"`
	UserId      int       `json:"user_id"`
	User        User      `json:"user" gorm:"foreignKey:UserId"`
}
