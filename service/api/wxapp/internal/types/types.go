// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Gender   int64  `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type UserInfoResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type ProductCreateRequest struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}

type ProductCreateResponse struct {
	Id int64 `json:"id"`
}

type ProductUpdateRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Desc   string `json:"desc,optional"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount,optional"`
	Status int64  `json:"status,optional"`
}

type ProductUpdateResponse struct {
}

type ProductRemoveRequest struct {
	Id int64 `json:"id"`
}

type ProductRemoveResponse struct {
}

type ProductDetailRequest struct {
	Id int64 `json:"id"`
}

type ProductDetailResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Stock  int64  `json:"stock"`
	Amount int64  `json:"amount"`
	Status int64  `json:"status"`
}

type CartsItem struct {
	Id        int64  `json:"id"`
	GoodsID   int64  `json:"goodsId"`
	GoodsName string `json:"goodsName"`
	Num       int64  `json:"num"`
}

type CartsListResponse struct {
	List []CartsItem `json:"list"`
}

type CartsAddRequest struct {
	GoodsID int64 `json:"goodsId"`
	Num     int64 `json:"num"`
}

type CartsAddResponse struct {
	Id string `json:"id"`
}

type CartsEditRequest struct {
	Id   int64 `json:"id"`
	Type int64 `json:"type"`
}

type CartsEditResponse struct {
}

type CartsDelRequest struct {
	Idx int64 `json:"idx"`
}

type CartsDelResponse struct {
}

type OrderCreateRequest struct {
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type OrderCreateResponse struct {
	Id int64 `json:"id"`
}

type OrderUpdateRequest struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid,optional"`
	Pid    int64 `json:"pid,optional"`
	Amount int64 `json:"amount,optional"`
	Status int64 `json:"status,optional"`
}

type OrderUpdateResponse struct {
}

type OrderRemoveRequest struct {
	Id int64 `json:"id"`
}

type OrderRemoveResponse struct {
}

type OrderDetailRequest struct {
	Id int64 `json:"id"`
}

type OrderDetailResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type OrderListRequest struct {
	Uid int64 `json:"uid"`
}

type OrderListResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type PayCreateRequest struct {
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
}

type PayCreateResponse struct {
	Id int64 `json:"id"`
}

type PayDetailRequest struct {
	Id int64 `json:"id"`
}

type PayDetailResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
	Source int64 `json:"source"`
	Status int64 `json:"status"`
}

type PayCallbackRequest struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Oid    int64 `json:"oid"`
	Amount int64 `json:"amount"`
	Source int64 `json:"source"`
	Status int64 `json:"status"`
}

type PayCallbackResponse struct {
}

type WxPayCallBackReq struct {
}

type WxPayCallBackResp struct {
	ReturnCode string `json:"return_code"`
}
