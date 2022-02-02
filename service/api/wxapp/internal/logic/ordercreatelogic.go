package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/order"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/product"
	"google.golang.org/grpc/status"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCreateLogic {
	return OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req types.OrderCreateRequest) (resp *types.OrderCreateResponse, err error) {

	//res, err := l.svcCtx.OrderRpc.Create(l.ctx, &orderclient.CreateRequest{
	//	Uid:    req.Uid,
	//	Pid:    req.Pid,
	//	Amount: req.Amount,
	//	Status: req.Status,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &types.CreateResponse{
	//	Id: res.Id,
	//}, nil
	// 获取 OrderRpc BuildTarget

	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 获取 ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://etcd:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/orderclient.Order/Create",
			orderRpcBusiServer+"/orderclient.Order/CreateRevert",
			&order.CreateRequest{
				Uid:    req.Uid,
				Pid:    req.Pid,
				Amount: req.Amount,
				Status: 0,
			}).
		Add(productRpcBusiServer+"/productclient.Product/DecrStock",
			productRpcBusiServer+"/productclient.Product/DecrStockRevert",
			&product.DecrStockRequest{
				Id:  req.Pid,
				Num: 1,
			})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.OrderCreateResponse{}, nil
}

//提示：SagaGrpc.Add 方法第一个参数 action 是微服务 grpc 访问的方法路径，这个方法路径需要分别去以下文件中寻找。
//mall/service/order/rpc/order/order.pb.go
//mall/service/product/rpc/product/product.pb.go
//按关键字 Invoke 搜索即可找到。
