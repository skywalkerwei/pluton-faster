package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/types"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/order"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/product"
	"google.golang.org/grpc/status"

<<<<<<< HEAD:service/api/wxapp/internal/logic/ordercreatelogic.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
<<<<<<< HEAD:service/api/wxapp/internal/logic/ordercreatelogic.go
=======
=======
	"github.com/zeromicro/go-zero/core/logx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/order/api/internal/logic/createlogic.go
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/order/api/internal/logic/createlogic.go
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.CreateRequest) (resp *types.CreateResponse, err error) {
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

	return &types.CreateResponse{}, nil
}

//提示：SagaGrpc.Add 方法第一个参数 action 是微服务 grpc 访问的方法路径，这个方法路径需要分别去以下文件中寻找。
//mall/service/order/rpc/order/order.pb.go
//mall/service/product/rpc/product/product.pb.go
//按关键字 Invoke 搜索即可找到。