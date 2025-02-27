package controllers

import (
	"cat_cafe/models"
	"cat_cafe/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CatController struct {
	Repo *repository.CatRepository
}

// Controller ใหม่ และเชื่อมกับ Repository
func NewCatController(repo *repository.CatRepository) *CatController {
	return &CatController{Repo: repo}
}

// รับข้อมูลจาก API `/cats` และบันทึกลง MongoDB
func (ctrl *CatController) CreateCat(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.Repo.CreateCat(cat)
	c.JSON(http.StatusOK, gin.H{"message": "Cat added"})
}

// ดึงข้อมูลแมวทั้งหมดจาก MongoDB
func (ctrl *CatController) GetCats(c *gin.Context) {
	cats, _ := ctrl.Repo.GetCats()
	c.JSON(http.StatusOK, cats)
}
