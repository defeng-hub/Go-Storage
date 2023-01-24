package impl

import (
	"context"
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/defeng-hub/Go-Storage/oss"
)

var _ oss.Service = (*LocalStorageImpl)(nil)

type LocalStorageImpl struct {
}

func (l LocalStorageImpl) CreateFile(ctx context.Context, file *model.File) (*model.File, error) {
	file.Url = *common.RunUrl + "/upload/" + file.Link
	return file, nil
}

func (l LocalStorageImpl) DeleteFile(ctx context.Context, file *model.File) (*model.File, error) {
	return file, nil
}
