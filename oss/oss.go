package oss

import (
	"context"
	"github.com/defeng-hub/Go-Storage/model"
)

type Service interface {
	CreateFile(ctx context.Context, file *model.File) (*model.File, error)
	DeleteFile(ctx context.Context, file *model.File) (*model.File, error)
}
