package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/huangzhuo492008824/go-gin-example/docs"
	"github.com/huangzhuo492008824/go-gin-example/middleware/jwt"
	"github.com/huangzhuo492008824/go-gin-example/pkg/setting"
	"github.com/huangzhuo492008824/go-gin-example/pkg/upload"
	"github.com/huangzhuo492008824/go-gin-example/routers/api"
	v1 "github.com/huangzhuo492008824/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/api/v1/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT)
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
