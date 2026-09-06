package repository

import "sp-backend/internal/domain"

type UserRepo struct {
	gql *GQLClient
}

func NewUserRepo(gql *GQLClient) *UserRepo {
	return &UserRepo{gql: gql}
}

func (r *UserRepo) GetAll() ([]domain.User, error) {
	var resp struct {
		Users []domain.User `json:"users"`
	}
	err := r.gql.Run(`query { users { id memberId email role status } }`, nil, &resp)
	return resp.Users, err
}

func (r *UserRepo) UpdateRole(id string, role string) (*domain.User, error) {
	var resp struct {
		UpdateUserRole *domain.User `json:"updateUserRole"`
	}
	err := r.gql.Run(`mutation($id: ID!, $role: String!) {
		updateUserRole(id: $id, role: $role) { id memberId email role status }
	}`, map[string]interface{}{"id": id, "role": role}, &resp)
	return resp.UpdateUserRole, err
}
