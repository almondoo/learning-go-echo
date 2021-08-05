package handler

import (
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Routing interface {
	InitCommonRouting()
	InitAuthUserRouting(UserHanlder)
	InitAuthAdminRouting(AdminHandler)
	InitCropRouting(CropHandler)
	InitProductRouting(ProductHandler)
	InitRoleRouting(RoleHandler)
	InitArticleRouting(ArticleHandler)
	InitArtistRouting(ArtistHanlder)
}

type routing struct {
	e           *echo.Group
	auth        auth.AuthInterface
	userToken   auth.TokenInterface
	adminToken  auth.TokenInterface
	artistToken auth.TokenInterface
}

const (
	AuthorTypeUser   = "user"
	AuthorTypeArtist = "artist"
	AuthorTypeAdmin  = "admin"
)

func NewRouting(e *echo.Group, auth auth.AuthInterface, userToken auth.TokenInterface, adminToken auth.TokenInterface, artistToken auth.TokenInterface) Routing {
	return &routing{e: e, auth: auth, userToken: userToken, adminToken: adminToken, artistToken: artistToken}
}

func (r *routing) InitCommonRouting() {
	r.e.GET("/api/first", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, map[string]string{
			"csrf": c.Get("csrf").(string),
		})
	}))
}

//- ユーザー関連処理
func (r *routing) InitAuthUserRouting(userHandler UserHanlder) {
	//- グループ
	base := r.e.Group("/api/user")

	//- ログイン処理
	base.POST("/login", userHandler.Login())

	//- アカウント作成
	base.POST("/register", userHandler.Register())

	//- jwtが必要なルート
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.userToken, AuthorTypeUser))

	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "check")
	}))

	jwt.POST("/logout", userHandler.Logout())

	jwt.POST("/edit", userHandler.Edit())
}

//- ユーザー関連処理
func (r *routing) InitAuthAdminRouting(adminHandler AdminHandler) {
	//- グループ
	base := r.e.Group("/api/manage/admin")

	//- ログイン
	base.POST("/login", adminHandler.Login())

	//- jwtが必要なルート
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.adminToken, AuthorTypeAdmin))

	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "check")
	}))

	//- アカウント作成
	jwt.POST("/register", adminHandler.Register())

	//- ログアウト
	jwt.POST("/logout", adminHandler.Logout())
}

func (r *routing) InitCropRouting(cropHandler CropHandler) {
	//- グループ
	base := r.e.Group("/api/crop")

	//-
	base.POST("", cropHandler.Crop())
}

func (r *routing) InitProductRouting(productHandler ProductHandler) {
	//- グループ
	base := r.e.Group("/api/product")

	//- 作品一覧取得
	base.GET("/list", productHandler.GetProductList())

	//- アーティストと作品の一覧取得
	base.GET("/artist/list/:artist", productHandler.GetProductArtistList())

	//- jwtが必要なルート
	// jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.artistToken))

	//- 作品詳細取得
	base.GET("/detail/:id", productHandler.GetProductDetail())

	//- 作品登録
	base.POST("/detail", productHandler.CreateProduct())

	//- 作品編集
	base.PUT("/detail/:id", productHandler.Edit())

	//- 作品削除
	// base.DELETE("/detail/:id",productHandler.Delete())
}

func (r *routing) InitArticleRouting(articleHandler ArticleHandler) {
	//- グループ
	base := r.e.Group("/api/article")

	//- 作品一覧取得
	base.GET("/list", articleHandler.GetArticleList())

	//- 作品一覧取得 アーティスト
	// base.GET("/artist/list/:artist", articleHandler.GetArticleArtistList())

	jwtGroup := base.Group("", middleware.AuthMiddleware(r.auth, r.artistToken, AuthorTypeArtist))

	//- 作品詳細取得
	base.GET("/detail/:id", articleHandler.GetArticleDetail())

	//- 作品登録
	jwtGroup.POST("/detail", articleHandler.CreateArticle())

	//- 作品編集
	jwtGroup.PUT("/detail/:id", articleHandler.EditArticle())

	//- 作品削除
	// base.DELETE("/detail/:id",articleHandler.Delete())
}

func (r *routing) InitArtistRouting(artistHandler ArtistHanlder) {
	//- グループ
	base := r.e.Group("/api/artist")

	//- ログイン処理
	base.POST("/login", artistHandler.Login())

	//- アカウント作成
	base.POST("/register", artistHandler.Register())

	//- jwtが必要なルート
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.artistToken, AuthorTypeArtist))

	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "check")
	}))

	jwt.POST("/logout", artistHandler.Logout())

	jwt.POST("/edit", artistHandler.Edit())

}

func (r *routing) InitRoleRouting(roleHandler RoleHandler) {
	//- グループ
	base := r.e.Group("/api/role")

	//-
	base.POST("", roleHandler.Get())
}
