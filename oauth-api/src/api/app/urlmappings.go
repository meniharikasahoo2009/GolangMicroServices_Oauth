package app

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meniharikasahoo2009/GolangMicroServices_Oauth/oauth-api/src/api/controller/oauth_controller"
)

func mapUrls() {
	subRouterAuthenticated := router.Group("/api/v1/ArticleName", gin.BasicAuth(gin.Accounts{
		"admin": "adminpass",
	}))
	subRouterAuthenticated.GET("/:Name", controller.GetArticles)
}
