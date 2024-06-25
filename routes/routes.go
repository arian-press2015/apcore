package routes

import (
	"apcore/controllers"
	"apcore/messages"
	"apcore/middlewares"
	"apcore/response"
	"apcore/utils/fileupload"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewRoutes),
)

func NewRouter() *gin.Engine {
	return gin.Default()
}

type Routes struct {
	controllers       *controllers.Controllers
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware
	fileUploader      *fileupload.LocalFileUploader
}

func NewRoutes(
	controllers *controllers.Controllers,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	fileUploader *fileupload.LocalFileUploader,
) *Routes {
	return &Routes{
		controllers:       controllers,
		jwtAuthMiddleware: jwtAuthMiddleware,
		fileUploader:      fileUploader,
	}
}

func (r *Routes) SetupRoutes(router *gin.Engine) {
	PingRoutes(router, r.controllers.PingController)
	AuthRoutes(router, r.controllers.AuthController, r.jwtAuthMiddleware)
	UsersRoutes(router, r.controllers.UserController, r.jwtAuthMiddleware)
	RolesRoutes(router, r.controllers.RoleController, r.jwtAuthMiddleware)
	AdminAuthRoutes(router, r.controllers.AdminAuthController, r.jwtAuthMiddleware)
	CustomersRoutes(router, r.controllers.CustomerController, r.jwtAuthMiddleware)
	SwaggerRoutes(router)

	router.NoMethod(func(c *gin.Context) {
		response.Error(c, nil, messages.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
	})

	router.NoRoute(func(c *gin.Context) {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
	})

	router.Static("/public", "./public")

	router.POST("/fileupload", r.fileUploadHandler)
}

func (r *Routes) fileUploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("get form err: %s", err.Error()), http.StatusBadRequest)
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("open file err: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer fileContent.Close()

	uploadDir := "./public/images"

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("could not create upload directory: %v", err), http.StatusInternalServerError)
		return
	}

	filename, err := r.fileUploader.Save(fileContent, uploadDir, file.Filename)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("upload file err: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response.Success(c, gin.H{"filename": filename}, messages.MsgSuccessful, nil, http.StatusOK)
}
