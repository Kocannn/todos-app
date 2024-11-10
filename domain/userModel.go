package domain

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;not null;autoUpdateTime"`
	Todos     []Todo    `json:"todos"`
}
