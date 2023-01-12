package impl

import (
	"context"
	"fmt"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/defeng-hub/Go-Storage/oss"
)

var _ oss.Service = (*AliYunImpl)(nil)

type AliYunImpl struct {
}

func (a AliYunImpl) CreateFile(ctx context.Context, file *model.File) (*model.File, error) {
	fmt.Printf("xxxxxxxxxxxxxxxxxxxxx")
	//TODO implement me
	panic("implement me")
}

func (a AliYunImpl) DeleteFile(ctx context.Context, file *model.File) (*model.File, error) {
	//TODO implement me
	panic("implement me")
}
