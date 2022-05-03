package domain

type Contact struct {
	UserID      string      `json:"user_id" binding:"required"`
	ContactType ContactType `json:"contact_type"`
	Value       string      `json:"value" binding:"required"`
}
