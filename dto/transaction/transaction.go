package transactiondto

type TransactionRequest struct {
	AccountNumber int    `json:"account_number" form:"account_number" grom:"type: int"`
	ProofTransfer string `json:"proof_transfer" form:"proof_transfer" gorm:"type varchar(50)"`
}

type TransactionResponse struct {
	ID            int    `json:"id"`
	AccountNumber int    `json:"account_number"`
	ProofTransfer string `json:"proof_transfer"`
	Status        string `json:"status"`
	UserID        int    `json:"user_id"`
}
