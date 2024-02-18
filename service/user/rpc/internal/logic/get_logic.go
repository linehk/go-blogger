package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-blogger/service/user/rpc/model"
	"github.com/linehk/go-blogger/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *user.GetReq) (*user.User, error) {
	appUser, err := l.svcCtx.AppUserModel.FindOneByUuid(l.ctx, in.UserId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, fmt.Errorf("FindOneByUuid NotFound err: %v", err)
	}
	if appUser == nil {
		return nil, fmt.Errorf("FindOneByUuid NotExists err: %v", err)
	}
	var respUser user.User
	err = copier.Copy(&respUser, appUser)
	if err != nil {
		return nil, fmt.Errorf("copier.Copy err: %v", err)
	}
	return &respUser, nil
}
