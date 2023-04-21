package logic

import (
	"blog/common/cryptx"
	"blog/common/jwtx"
	"context"
	"google.golang.org/grpc/status"
	"time"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	table := l.svcCtx.Bkmodel.User
	user, err := table.WithContext(l.ctx).Where(table.Mobile.Eq(req.Mobile)).Debug().First()

	if err != nil {
		return nil, status.Error(100, "用户不存在")
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if password != user.Password {
		return nil, status.Error(500, "密码错误")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, int64(user.ID))
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
