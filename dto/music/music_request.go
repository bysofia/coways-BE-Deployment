package musicdto

type MusicRequest struct {
	Title     string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(255)"`
	Year      int    `json:"year" form:"year" gorm:"type: int"`
	SingerID  int    `json:"singer_id" form:"singer_id" gorm:"type: int"`
	MusicFile string `json:"music_file" form:"music_file" gorm:"type: varchar(255)"`
}

type UpdateMusicRequest struct {
	Title     string `json:"title" form:"title" gorm:"type: varchar:(255)"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(255)"`
	Year      int    `json:"year" form:"year" gorm:"type: int"`
	SingerID  int    `json:"singer_id" form:"singer_id" gorm:"type: int"`
	MusicFile string `json:"music_file" form:"music_file" gorm:"type: varchar(255)"`
}

type MusicResponse struct {
	Title     string `json:"title" form:"title" gorm:"type: varchar:(255)"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(255)"`
	Year      int    `json:"year" form:"year" gorm:"type: int"`
	SingerID  int    `json:"singer_id" form:"singer_id" gorm:"type: int"`
	MusicFile string `json:"music_file" form:"music_file" gorm:"type: varchar(255)"`
}
