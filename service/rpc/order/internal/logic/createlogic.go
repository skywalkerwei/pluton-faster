package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/user"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"

	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err != nil {
			return fmt.Errorf("用户不存在")
		}

		newOrder := model.Order{
			Uid:    in.Uid,
			Pid:    in.Pid,
			Amount: in.Amount,
			Status: 0,
		}
		// 创建订单
		_, err = l.svcCtx.OrderModel.TxInsert(tx, &newOrder)
		if err != nil {
			return fmt.Errorf("订单创建失败")
		}

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil

}
