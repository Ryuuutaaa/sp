package storage

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"sp-backend/internal/config"
	"time"
)

type R2Storage struct {
	cfg *config.Config
}

func NewR2Storage(cfg *config.Config) *R2Storage {
	return &R2Storage{cfg: cfg}
}

func (s *R2Storage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	// Stub uploader R2 (bisa dikembangkan dengan AWS S3 SDK v2 compatible)
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s/%d%s", folder, time.Now().UnixNano(), ext)
	publicURL := fmt.Sprintf("%s/%s", s.cfg.R2PublicURL, filename)
	return publicURL, nil
}
