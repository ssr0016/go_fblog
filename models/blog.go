package models

type Blog struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      string `json:"userID"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
}
