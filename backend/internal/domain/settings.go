package domain

type SettingRepository interface {
	GetAll() ([]Setting, error)
	GetByKey(key string) (*Setting, error)
	Upsert(input UpsertSettingInput) (*Setting, error)
}

type SettingService interface {
	GetAll() ([]Setting, error)
	GetByKey(key string) (*Setting, error)
	Upsert(input UpsertSettingInput) (*Setting, error)
}

type UpsertSettingInput struct {
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
}
