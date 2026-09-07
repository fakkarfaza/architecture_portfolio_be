package model

type ContactMessage struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Subject   *string `json:"subject,omitempty"`
	Message   string  `json:"message"`
	CreatedAt *string `json:"created_at,omitempty"`
}
