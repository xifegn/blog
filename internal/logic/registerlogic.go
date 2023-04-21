package logic

import (
	"blog/bkmodel/dao/model"
	"blog/common/cryptx"
	"context"
	"google.golang.org/grpc/status"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	table := l.svcCtx.Bkmodel.User
	_, err := table.WithContext(l.ctx).Where(table.Mobile.Eq(req.Mobile)).Debug().First()
	if err == nil {
		return nil, status.Error(500, "用户已存在")
	}
	newUser := model.User{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
	}
	err = table.WithContext(l.ctx).Create(&newUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.RegisterResp{
		Name:   req.Name,
		Mobile: req.Mobile,
	}, nil
}
