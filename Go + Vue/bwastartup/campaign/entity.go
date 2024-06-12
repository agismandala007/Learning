package campaign

import "time"

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	ShortDescribe  string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmmount    int
	CurrentAmount  int
	Slug           string
	CreateAt       time.Time
	UpdateAt       time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string
	isPrimary  int
	CreateAt   time.Time
	UpdateAt   time.Time
}
