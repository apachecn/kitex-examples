// Code generated by Kitex v0.2.1. DO NOT EDIT.

package shopservice

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return shopServiceServiceInfo
}

var shopServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ShopService"
	handlerType := (*shop.ShopService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SettleShop":      kitex.NewMethodInfo(settleShopHandler, newShopServiceSettleShopArgs, newShopServiceSettleShopResult, false),
		"GetShopIdByName": kitex.NewMethodInfo(getShopIdByNameHandler, newShopServiceGetShopIdByNameArgs, newShopServiceGetShopIdByNameResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "shop",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.2.1",
		Extra:           extra,
	}
	return svcInfo
}

func settleShopHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*shop.ShopServiceSettleShopArgs)
	realResult := result.(*shop.ShopServiceSettleShopResult)
	success, err := handler.(shop.ShopService).SettleShop(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopServiceSettleShopArgs() interface{} {
	return shop.NewShopServiceSettleShopArgs()
}

func newShopServiceSettleShopResult() interface{} {
	return shop.NewShopServiceSettleShopResult()
}

func getShopIdByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*shop.ShopServiceGetShopIdByNameArgs)
	realResult := result.(*shop.ShopServiceGetShopIdByNameResult)
	success, err := handler.(shop.ShopService).GetShopIdByName(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShopServiceGetShopIdByNameArgs() interface{} {
	return shop.NewShopServiceGetShopIdByNameArgs()
}

func newShopServiceGetShopIdByNameResult() interface{} {
	return shop.NewShopServiceGetShopIdByNameResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SettleShop(ctx context.Context, req *shop.SettleShopReq) (r *shop.SettleShopResp, err error) {
	var _args shop.ShopServiceSettleShopArgs
	_args.Req = req
	var _result shop.ShopServiceSettleShopResult
	if err = p.c.Call(ctx, "SettleShop", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetShopIdByName(ctx context.Context, req *shop.GetShopIdByNameReq) (r *shop.GetShopIdByNameResp, err error) {
	var _args shop.ShopServiceGetShopIdByNameArgs
	_args.Req = req
	var _result shop.ShopServiceGetShopIdByNameResult
	if err = p.c.Call(ctx, "GetShopIdByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
