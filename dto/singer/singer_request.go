package singerdto

type SingerRequest struct {
	Title       string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(50)"`
	Old         int    `json:"old" form:"old" gorm:"type: int"`
	Category    string `json:"category" form:"category" gorm:"type: varvhar(255)"`
	StartCareer int    `json:"start_career" form:"start_career" gorm:"type: varchar(255)"`
}

type UpdateSingerRequest struct {
	Title       string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(255)"`
	Old         int    `json:"old" form:"old" gorm:"type: int"`
	Category    string `json:"category" form:"category" gorm:"type: varvhar(255)"`
	StartCareer int    `json:"start_career" form:"start_career" gorm:"type: varchar(255)"`
}

type SingerResponse struct {
	Title       string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" gorm:"type varchar(225)"`
	Old         int    `json:"old" form:"old" gorm:"type: int"`
	Category    string `json:"category" form:"category" gorm:"type: varvhar(255)"`
	StartCareer int    `json:"start_career" form:"start_career" gorm:"type: varchar(255)"`
}
