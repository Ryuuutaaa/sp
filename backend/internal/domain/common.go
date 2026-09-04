package domain

type Setting struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type User struct {
	ID       string  `json:"id"`
	MemberID *string `json:"memberId"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Status   string  `json:"status"`
}
