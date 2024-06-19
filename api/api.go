package api

import (
	"github.com/Javokhdev/Portfolio-Api-Gateway/api/handler"
	_"github.com/Javokhdev/Portfolio-Api-Gateway/docs"


	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Portfolio
// @version 1.0
// @description Portfolio
// @host localhost:8080
// @BasePath /
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	u := r.Group("/skill")
	u.POST("/create", h.CreateSkill)
	u.PUT("/update/:id", h.UpdateSkill)
	u.DELETE("/delete/:id", h.DeleteSkill)
	u.GET("/getall", h.GetAllSkill)
	u.GET("/getbyid/:id", h.GetByIdSkill)
	u.GET("/byuser/:user_id", h.GetByUserIdSkill)



	p := r.Group("/project")
	p.POST("/create", h.CreateProject)
	p.PUT("/update/:id", h.UpdateProject)
	p.DELETE("/delete/:id", h.DeleteProject)
	p.GET("/getall", h.GetAllProject)
	p.GET("/getbyid/:id", h.GetByIdProject)
	p.GET("/byuser", h.GetByUserIdProject)
	p.GET("/search", h.SearchProjects)



	e := r.Group("/education")
	e.POST("/create", h.CreateEducation)
	e.PUT("/update/:id", h.UpdateEducation)
	e.DELETE("/delete/:id", h.DeleteEducation)
	e.GET("/getall", h.GetAllEducation)
	e.GET("/getbyid/:id", h.GetByIdEducation)
	e.GET("/byuser/:user_id", h.GetByUserIdEducation)



	v := r.Group("/experience")
	v.POST("/create", h.CreateExperience)
	v.PUT("/update/:id", h.UpdateExperience)
	v.DELETE("/delete/:id", h.DeleteExperience)
	v.GET("/getall", h.GetAllExperience)
	v.GET("/getbyid/:id", h.GetByIdExperience)
	v.GET("/byuser/:user_id", h.GetByUserIdExperience)



	return r
}
