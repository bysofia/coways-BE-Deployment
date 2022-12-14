package models

type Music struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Title     string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year      int    `json:"year" form:"year" gorm:"type: varchar(255)"`
	SingerID  int    `json:"singer_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Singer    Singer `json:"singer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MusicFile string `json:"music_file" form:"music_file" gorm:"type: varchar(255)"`
}
