package domain

import "time"

type Member struct {
	ID                 string     `json:"id"`
	MemberNumber       string     `json:"memberNumber"`
	Name               string     `json:"name"`
	NIK                string     `json:"nik"`
	Address            string     `json:"address"`
	Phone              string     `json:"phone"`
	Occupation         string     `json:"occupation"`
	KTPPhotoURL        string     `json:"ktpPhotoUrl"`
	SelfieKTPPhotoURL  string     `json:"selfieKtpPhotoUrl"`
	VerificationStatus string     `json:"verificationStatus"`
	VerifiedBy         *string    `json:"verifiedBy"`
	VerifiedAt         *time.Time `json:"verifiedAt"`
	JoinDate           string     `json:"joinDate"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
}

type MemberRepository interface {
	GetAll() ([]Member, error)
	GetByID(id string) (*Member, error)
	Create(input CreateMemberInput) (*Member, error)
	Update(id string, input UpdateMemberInput) (*Member, error)
	UpdateStatus(id string, status string) (*Member, error)
	Verify(id string, status string, verifiedBy string) (*Member, error)
}

type MemberService interface {
	GetAll() ([]Member, error)
	GetByID(id string) (*Member, error)
	Register(input CreateMemberInput) (*Member, error)
	Update(id string, input UpdateMemberInput) (*Member, error)
	Deactivate(id string) (*Member, error)
	Verify(id string, status string, verifiedBy string) (*Member, error)
}

type CreateMemberInput struct {
	MemberNumber      string `json:"memberNumber"`
	Name              string `json:"name"`
	NIK               string `json:"nik"`
	Address           string `json:"address"`
	Phone             string `json:"phone"`
	Occupation        string `json:"occupation"`
	KTPPhotoURL       string `json:"ktpPhotoUrl"`
	SelfieKTPPhotoURL string `json:"selfieKtpPhotoUrl"`
}

type UpdateMemberInput struct {
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	Phone      *string `json:"phone"`
	Occupation *string `json:"occupation"`
}
