package authdto

type RegisterResponse struct {
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)" json:"status"`
	// Token  string `gorm:"type: varchar(255)" json:"token"`
}

type LoginResponse struct {
	Id     int    `gorm:"type: int" json:"-"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
	Status string `gorm:"type: varchar(50)" json:"status"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)" json:"status"`
}
