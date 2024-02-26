package logic

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/linehk/go-microservices-blogger/errcode"
	"github.com/linehk/go-microservices-blogger/service/blog/rpc/blog"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/internal/svc"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/model"
	"github.com/linehk/go-microservices-blogger/service/user/rpc/user"

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
	appUserModel, err := l.svcCtx.AppUserModel.FindOneByUuid(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.UserNotExist))
		return nil, errcode.Wrap(errcode.UserNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	var userResp user.User
	err = copier.Copy(&userResp, appUserModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Convert))
		return nil, errcode.Wrap(errcode.Convert)
	}
	userResp.Id = appUserModel.Uuid
	userResp.Kind = "blogger#user"

	localeModel, err := l.svcCtx.LocaleModel.FindOneByAppUserUuid(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		l.Error(errcode.Msg(errcode.LocaleNotExist))
		return nil, errcode.Wrap(errcode.LocaleNotExist)
	}
	if err != nil {
		l.Error(errcode.Msg(errcode.Database))
		return nil, errcode.Wrap(errcode.Database)
	}

	err = copier.Copy(&userResp.Locale, localeModel)
	if err != nil {
		l.Error(errcode.Msg(errcode.Convert))
		return nil, errcode.Wrap(errcode.Convert)
	}

	listByUserReq := &blog.ListByUserReq{
		UserId: in.GetUserId(),
	}
	listByUserResp, err := l.svcCtx.BlogService.ListByUser(l.ctx, listByUserReq)
	if err != nil {
		l.Error(errcode.Msg(errcode.Service))
		return nil, errcode.Wrap(errcode.Service)
	}
	for i, blogItem := range listByUserResp.GetItems() {
		userResp.Blogs[i].SelfLink = blogItem.GetSelfLink()
	}

	return &userResp, nil
}
