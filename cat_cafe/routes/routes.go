package routes

import (
    "cat_cafe/controllers"
    "github.com/gin-gonic/gin"
)

// ✅ กำหนด API Path
func SetupRoutes(router *gin.Engine, catController *controllers.CatController) {
    router.POST("/cats", catController.CreateCat)
    router.GET("/cats", catController.GetCats)
}
