package system

import (
	"go-web/model/system"
	"go-web/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")

	{
		baseRouter.POST("login", func(context *gin.Context) {
			context.JSON(http.StatusOK, "ok")
		})
		baseRouter.POST("register", func(context *gin.Context) {
			var form system.Register
			if err := context.ShouldBindJSON(&form); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	return baseRouter
}
