package models

type Singer struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Old         int    `json:"old" form:"old" gorm:"type: int"`
	Category    string `json:"category" form:"category" gorm:"type: varchar(255)"`
	StartCareer int    `json:"start_career" form:"start_career" gorm:"type: int"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
}
