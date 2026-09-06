package domain

type UserRepository interface {
	GetAll() ([]User, error)
	UpdateRole(id string, role string) (*User, error)
}

type UserService interface {
	GetAll() ([]User, error)
	UpdateRole(id string, role string) (*User, error)
}

var ValidRoles = []string{"super_admin", "admin", "teller", "anggota"}
