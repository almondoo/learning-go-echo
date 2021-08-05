package main

import (
	"fmt"
	"learning-go-echo/application/usecase"
	"learning-go-echo/db"
	"learning-go-echo/domain/service"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/infrastructure/storedb"
	"learning-go-echo/interface/fileupload"
	"learning-go-echo/interface/handler"
	"learning-go-echo/interface/middleware"
	"learning-go-echo/interface/validation"
	"os"

	env "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	err := env.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	e := echo.New()

	//- CustomValidation
	e.Validator = validation.NewValidator()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(middleware.CSRFMiddleware())
	e.Use(middleware.CORSMiddleware())

	// authClient, err := store.NewDataStoreDB(os.Getenv("PROJECT_ID"))
	authClient := storedb.NewRedisDB()
	newAuth := auth.NewRedisAuth(authClient)

	//- Repository
	repositories, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer repositories.Close()
	if os.Getenv("ENV") == "local" {
		repositories.Seeder()
	}

	//- validator
	dbValidator := validation.NewValidatorWithDB(repositories.DB)
	e.Use(middleware.CustomContextMiddleware(dbValidator))

	//- Newトークン
	userToken := auth.NewUserToken()
	adminToken := auth.NewAdminToken()
	artistToken := auth.NewArtistToken()
	group := e.Group("")
	routing := handler.NewRouting(group, newAuth, userToken, adminToken, artistToken)

	//- 共通Routing
	routing.InitCommonRouting()

	//- User関連のRouting
	userService := service.NewUserService(repositories.User, newAuth, userToken)
	userUsecase := usecase.NewUserUsecase(userService, newAuth, userToken)
	userHandler := handler.NewUserHandler(userUsecase)
	routing.InitAuthUserRouting(userHandler)

	//- User関連のRouting
	adminService := service.NewAdminService(repositories.Admin, newAuth, adminToken)
	authAdminUsecase := usecase.NewAdminUsecase(adminService, newAuth, adminToken)
	adminHandler := handler.NewAdminHandler(authAdminUsecase)
	routing.InitAuthAdminRouting(adminHandler)

	fileupload := fileupload.NewFileUpload()
	cropHandler := handler.NewCropHandler(fileupload)
	routing.InitCropRouting(cropHandler)

	productService := service.NewProductService(repositories.Product, repositories.Artist)
	productImageService := service.NewProductImageService(repositories.ProductImage)
	productUsecase := usecase.NewProductUsecase(productService, productImageService)
	productHandler := handler.NewProductHandler(productUsecase)
	routing.InitProductRouting(productHandler)

	roleService := service.NewRoleService(repositories.Role)
	roleUsecase := usecase.NewRoleUsecase(roleService)
	roleHandler := handler.NewRoleHandler(roleUsecase)
	routing.InitRoleRouting(roleHandler)

	articleService := service.NewArticleService(repositories.Article)
	articleUsecase := usecase.NewArticleUsecase(articleService)
	articleHandler := handler.NewArticleHandler(articleUsecase)
	routing.InitArticleRouting(articleHandler)

	artistService := service.NewArtistService(repositories.Artist, newAuth, artistToken)
	artistUsecase := usecase.NewArtistUsecase(artistService, newAuth, artistToken)
	artistHandler := handler.NewArtistHandler(artistUsecase)
	routing.InitArtistRouting(artistHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
