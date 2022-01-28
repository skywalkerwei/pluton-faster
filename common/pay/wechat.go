package pay

//func WxMini() {
//	client, err := wechat.NewClientV3("", "", "", "")
//	if err != nil {
//
//		return
//	}
//	fmt.Println(client)
//}

//func WxNotify(context *gin.Context) {
//	notifyReq, err := wechat.V3ParseNotify(context.Request)
//	if err != nil {
//		return
//	}
//
//	fmt.Println("WxNotify", notifyReq)
//	//err = notifyReq.VerifySignByPK(wxPublicKey)
//	//if err != nil {
//	//	xlog.Error(err)
//	//	return
//	//}
//	//result, err := notifyReq.DecryptCipherText(apiV3Key)
//
//	//c.JSON(http.StatusOK, &wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
//
//}
