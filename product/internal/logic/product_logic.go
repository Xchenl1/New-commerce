package logic

import (
	"E-commerce_system/code"
	"E-commerce_system/model"
	"E-commerce_system/product/internal/dao"
	"E-commerce_system/serializer"
	"E-commerce_system/service"
	"E-commerce_system/tools"
	"context"
	"gorm.io/gorm"
	"strconv"
	"sync"

	"E-commerce_system/product/internal/svc"
	"E-commerce_system/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductLogic) Product(req *types.ProductCreateRequest) (resp *types.Response, err error) {
	//todo:核心是先将用户输入的商品信息更新到product表中，然后再将图片一张张存入product_img表中  将第一张设置为首页  将信息返回给用户
	file, _ := l.svcCtx.File[0].Open() //第一张图片做为商品的首页
	userid, name := tools.VerityToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.JwtAuth.Jwtauth)
	resp = new(types.Response)
	var boss model.User
	err = dao.LookUserInfo(l.svcCtx.DB, userid, name, &boss)
	if err != nil {
		resp.Status = code.ErrorExistUser
		resp.Msg = code.GetMsg(code.ErrorExistUser)
		resp.Data = err
		return
	}
	path, err := service.UploadProductToLocalStatic(file, userid, name) //加载到本地路径
	if err != nil {
		resp.Status = code.ErrorProductImgUpload
		resp.Msg = code.GetMsg(code.ErrorProductImgUpload)
		return
	}
	product := model.Product{
		Model:         gorm.Model{},
		Name:          req.Name,
		CategoryId:    req.CategoryId,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		Pricediscount: req.DiscountPrice,
		OnSale:        true,
		Num:           req.Num,
		BossId:        boss.ID,
		BossName:      boss.NickName,
		BossAvatar:    boss.Avatar,
	}
	err = dao.CreateProduct(l.svcCtx.DB, &product)
	if err != nil {
		resp.Status = code.Error
		resp.Msg = code.GetMsg(code.Error)
		resp.Data = err
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(l.svcCtx.File))
	for index, file := range l.svcCtx.File {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		path, err = service.UploadProductToLocalStatic(tmp, userid, name+num)
		if err != nil {
			resp.Status = code.ErrorProductImgUpload
			resp.Msg = code.GetMsg(code.ErrorProductImgUpload)
			resp.Data = err
			return
		}
		var productimg = model.ProductImg{
			ProductId: req.CategoryId,
			ImgPath:   path,
		}
		err = dao.CreateProductImg(l.svcCtx.DB, &productimg)
		if err != nil {
			resp.Status = code.ErrorProductImgUpload
			resp.Msg = code.GetMsg(code.ErrorProductImgUpload)
			resp.Data = err
			return
		}
		wg.Done()
	}
	wg.Wait()
	resp.Status = code.Success
	resp.Msg = code.GetMsg(code.Success)
	resp.Data = serializer.BuildProduct(product, l.svcCtx.Config.Host, strconv.Itoa(l.svcCtx.Config.Port), l.ctx, l.svcCtx.Redis)
	return
}
