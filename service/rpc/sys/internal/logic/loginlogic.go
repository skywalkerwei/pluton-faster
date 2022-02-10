package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sys"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sys.LoginReq) (*sys.LoginResp, error) {
	//userInfo, err := l.svcCtx.UserModel.FindOneByName(in.UserName)
	//
	//switch err {
	//case nil:
	//case model.ErrNotFound:
	//	logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", in.UserName, err.Error())
	//	return nil, status.Error(100, "用户不存在")
	//default:
	//	logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", in.UserName, err.Error())
	//	return nil, err
	//}
	//password := cryptx.PasswordEncrypt(userInfo.Salt, in.Password)
	//if password != userInfo.Password {
	//	return nil, status.Error(100, "密码错误")
	//}
	//resp := &sys.LoginResp{
	//	Status:           userInfo.Status,
	//	CurrentAuthority: "admin",
	//	Id:               userInfo.Id,
	//	UserName:         userInfo.Name,
	//}
	return &sys.LoginResp{}, nil
}
