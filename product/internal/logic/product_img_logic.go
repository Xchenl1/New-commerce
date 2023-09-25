package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/product/internal/dao"
	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"
	"E-commerce_system/serializer"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductImgLogic {
	return &ProductImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductImgLogic) ProductImg(req *types.ProductImgRequest) (resp *types.Response, err error) {
	resp = new(types.Response)
	var productimg []*model.ProductImg
	err = dao.ListProductImg(l.svcCtx.DB, req.Id, &productimg)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildProductImgs(productimg, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port))
	return
}
