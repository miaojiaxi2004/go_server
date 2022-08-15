package system

import "time"

type CommunityIndex struct {
	ID           uint      `gorm:"primarykey"`
	Titlepic     string    `json:"titlepic"`
	Status       string    `json:"status"`
	UserID       int       `json:"user_id"`
	CreateDate   time.Time // 创建时间
	SupportCount int       `json:"support_count"`
	Title        string    `json:"title"`
	Text         string    `json:"text"`
}

func (CommunityIndex) TableName() string {
	return "api_communityindex"
}
