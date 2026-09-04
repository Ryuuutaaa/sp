package repository

import "sp-backend/internal/domain"

type MemberRepo struct {
	gql *GQLClient
}

func NewMemberRepo(gql *GQLClient) *MemberRepo {
	return &MemberRepo{gql: gql}
}

func (r *MemberRepo) GetAll() ([]domain.Member, error) {
	var resp struct {
		Members []domain.Member `json:"members"`
	}
	err := r.gql.Run(`query { members { id memberNumber name nik address phone occupation ktpPhotoUrl selfieKtpPhotoUrl verificationStatus status joinDate } }`, nil, &resp)
	return resp.Members, err
}

func (r *MemberRepo) GetByID(id string) (*domain.Member, error) {
	var resp struct {
		Member *domain.Member `json:"member"`
	}
	err := r.gql.Run(`query($id: ID!) { member(id: $id) { id memberNumber name nik address phone occupation ktpPhotoUrl selfieKtpPhotoUrl verificationStatus verifiedBy verifiedAt status joinDate } }`, map[string]interface{}{"id": id}, &resp)
	return resp.Member, err
}

func (r *MemberRepo) Create(input domain.CreateMemberInput) (*domain.Member, error) {
	var resp struct {
		CreateMember *domain.Member `json:"createMember"`
	}
	err := r.gql.Run(`mutation($memberNumber: String!, $name: String!, $nik: String!, $address: String!, $phone: String!, $occupation: String!, $ktpPhotoUrl: String!, $selfieKtpPhotoUrl: String!) {
		createMember(memberNumber: $memberNumber, name: $name, nik: $nik, address: $address, phone: $phone, occupation: $occupation, ktpPhotoUrl: $ktpPhotoUrl, selfieKtpPhotoUrl: $selfieKtpPhotoUrl) {
			id memberNumber name nik address phone occupation ktpPhotoUrl selfieKtpPhotoUrl verificationStatus status
		}
	}`, map[string]interface{}{
		"memberNumber":      input.MemberNumber,
		"name":              input.Name,
		"nik":               input.NIK,
		"address":           input.Address,
		"phone":             input.Phone,
		"occupation":        input.Occupation,
		"ktpPhotoUrl":       input.KTPPhotoURL,
		"selfieKtpPhotoUrl": input.SelfieKTPPhotoURL,
	}, &resp)
	return resp.CreateMember, err
}

func (r *MemberRepo) Update(id string, input domain.UpdateMemberInput) (*domain.Member, error) {
	vars := map[string]interface{}{"id": id}
	if input.Name != nil {
		vars["name"] = *input.Name
	}
	if input.Address != nil {
		vars["address"] = *input.Address
	}
	if input.Phone != nil {
		vars["phone"] = *input.Phone
	}
	if input.Occupation != nil {
		vars["occupation"] = *input.Occupation
	}
	var resp struct {
		UpdateMember *domain.Member `json:"updateMember"`
	}
	err := r.gql.Run(`mutation($id: ID!, $name: String, $address: String, $phone: String, $occupation: String) {
		updateMember(id: $id, name: $name, address: $address, phone: $phone, occupation: $occupation) {
			id memberNumber name nik address phone occupation verificationStatus status
		}
	}`, vars, &resp)
	return resp.UpdateMember, err
}

func (r *MemberRepo) UpdateStatus(id string, status string) (*domain.Member, error) {
	var resp struct {
		UpdateMemberStatus *domain.Member `json:"updateMemberStatus"`
	}
	err := r.gql.Run(`mutation($id: ID!, $status: String!) {
		updateMemberStatus(id: $id, status: $status) { id status }
	}`, map[string]interface{}{"id": id, "status": status}, &resp)
	return resp.UpdateMemberStatus, err
}

func (r *MemberRepo) Verify(id string, status string, verifiedBy string) (*domain.Member, error) {
	var resp struct {
		UpdateMemberVerification *domain.Member `json:"updateMemberVerification"`
	}
	err := r.gql.Run(`mutation($id: ID!, $status: String!, $verifiedBy: ID!) {
		updateMemberVerification(id: $id, status: $status, verifiedBy: $verifiedBy) { id verificationStatus verifiedBy }
	}`, map[string]interface{}{"id": id, "status": status, "verifiedBy": verifiedBy}, &resp)
	return resp.UpdateMemberVerification, err
}
