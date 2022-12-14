package models

import "time"

type Transaction struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int          `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User          UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountNumber int          `json:"account_number"`
	ProofTransfer string       `json:"prrof_transfer"`
	Remaining     int          `json:"remaining_active"`
	Status        string       `json:"status"`
	Amount        int          `json:"ammount"`
	Subscription  string       `json:"subscription"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
