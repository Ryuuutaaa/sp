package repository

import "sp-backend/internal/domain"

type SettingRepo struct {
	gql *GQLClient
}

func NewSettingRepo(gql *GQLClient) *SettingRepo {
	return &SettingRepo{gql: gql}
}

func (r *SettingRepo) GetAll() ([]domain.Setting, error) {
	var resp struct {
		Settings []domain.Setting `json:"settings"`
	}
	err := r.gql.Run(`query { settings { id key value description updatedAt } }`, nil, &resp)
	return resp.Settings, err
}

func (r *SettingRepo) GetByKey(key string) (*domain.Setting, error) {
	var resp struct {
		Setting *domain.Setting `json:"setting"`
	}
	err := r.gql.Run(`query($key: String!) { setting(key: $key) { id key value description updatedAt } }`, map[string]interface{}{"key": key}, &resp)
	return resp.Setting, err
}

func (r *SettingRepo) Upsert(input domain.UpsertSettingInput) (*domain.Setting, error) {
	var resp struct {
		UpsertSetting *domain.Setting `json:"upsertSetting"`
	}
	vars := map[string]interface{}{"key": input.Key, "value": input.Value}
	if input.Description != nil {
		vars["description"] = *input.Description
	}
	err := r.gql.Run(`mutation($key: String!, $value: String!, $description: String) {
		upsertSetting(key: $key, value: $value, description: $description) { id key value description }
	}`, vars, &resp)
	return resp.UpsertSetting, err
}
