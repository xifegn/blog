package logic

import (
	"blog/internal/svc"
	"blog/internal/types"
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	table := l.svcCtx.Bkmodel.User
	user, err := table.WithContext(l.ctx).Where(table.ID.Eq(int32(uid))).First()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.UserInfoResp{
		Id:   int64(user.ID),
		Name: user.Name,
	}, nil
}
