// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"E-commerce_system/user/internal/middleware"
	"net/http"

	"E-commerce_system/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{middleware.NewCorsMiddleware().CorsHandle},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: UserLoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: UserRegistrHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware.JwtAuth, middleware.NewCorsMiddleware().CorsHandle},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/:id",
					Handler: UserGetidHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: UserUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/image",
					Handler: UserImageUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/sendemail",
					Handler: UserSendEmailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/validemail",
					Handler: UserValidEmailHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
