package models

type Event struct {
	ID               int    `gorm:"column:id; primary_key; not null" json:"id"`
	Title            string `gorm:"column:title; not null" json:"title"`
	ShortDescription string `gorm:"column:short_description; not null" json:"short_description"`
	Description      string `gorm:"column:description; not null" json:"description"`
	Location         string `gorm:"column:location; not null" json:"location"`
	Date             string `gorm:"column:date; not null" json:"date"`
	Time             string `gorm:"column:time; not null" json:"time"`
	IsFeatured       bool   `gorm:"column:is_featured; not null" json:"is_featured"`
	CreatedBy        int    `gorm:"column:created_by; not null" json:"created_by"`
	User             User   `gorm:"foreignKey:CreatedBy; references:ID"`
	BaseModel
}

type EventUser struct {
	ID      int   `gorm:"column:id; primary_key; not null" json:"id"`
	EventID int   `gorm:"column:event_id; not null" json:"event_id"`
	Event   Event `gorm:"foreignKey:EventID; references:ID"`
	UserID  int   `gorm:"column:user_id; not null" json:"user_id"`
	User    User  `gorm:"foreignKey:UserID; references:ID"`
	BaseModel
}
