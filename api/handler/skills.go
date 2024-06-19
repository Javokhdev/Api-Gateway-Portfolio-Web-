package handler

import (
	"net/http"

	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Skill handles the creation of a new Porfolio
// @Summary 		Create Porfolio
// @Description 	Create page
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Skill  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/skill/create [post]
func (h *Handler) CreateSkill(ctx *gin.Context) {
	arr := pb.Skill{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Skill.CreateSkill(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateSkill 	handles the creation of a new Skill
// @Summary 		Update Skill
// @Description 	Update page
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Skill ID"
// @Param   		Update  body    pb.Skill  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/skill/update/{id} [put]
func (h *Handler) UpdateSkill(ctx *gin.Context) {
	skill := pb.Skill{}
	if err := ctx.BindJSON(&skill); err != nil {
		panic(err)
	}
	skill.Id = ctx.Param("id")
	_, err := h.Skill.UpdateSkill(ctx, &skill)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, "Success!!!")
}

// DeleteSkill 	handles the creation of a new Skill
// @Summary 		Delete Skill
// @Description 	Delete page
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Skill ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/skill/delete/{id} [delete]
func (h *Handler) DeleteSkill(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Skill.DeleteSkill(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllSkill 	handles the creation of a new Skill
// @Summary 		GetAll Skill
// @Description 	GetAll page
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Skill true  "Query parameter"
// @Success 		200  {object}  pb.GetAllSkills  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/skill/getall [get]
func (h *Handler) GetAllSkill(ctx *gin.Context) {
	skill := &pb.Skill{}
	// restoran.Address = ctx.Param("restoran")
	skill.Name = ctx.Param("name")

	res, err := h.Skill.GetAllSkill(ctx, skill)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdSkill 	handles the creation of a new Skill
// @Summary 		GetById Skill
// @Description 	GetById page
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Skill ID"
// @Success 		200 {object}  pb.Skill   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/skill/getbyid/{id} [get]
func (h *Handler) GetByIdSkill(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Skill.GetByIdSkill(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByUserIdSkill handles retrieving Skills by User ID
// @Summary 		Get Skills by User ID
// @Description 	Retrieve skills by user ID
// @Tags 			Skill
// @Accept  		json
// @Produce  		json
// @Param     		user_id    path    string  true  "User ID"
// @Success 		200 {object}  pb.GetAllSkills   "Get Skills by User ID Successful"
// @Failure 		400 {string}  string 			"User ID is required"
// @Failure 		404 {string}  string 			"Skills not found"
// @Failure 		500 {string}  string 			"Error while retrieving skills"
// @Router 			/skill/byuser/{user_id} [get]
func (h *Handler) GetByUserIdSkill(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	if user_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	skill := &pb.Skill{}
	res, err := h.Skill.GetAllSkill(ctx, skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving skills"})
		return
	}

	// Filter skills by user_id
	var userSkills []*pb.Skill
	for _, skill := range res.Skills {
		if skill.UserId == user_id {
			userSkills = append(userSkills, skill)
		}
	}

	if len(userSkills) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Skills not found"})
		return
	}

	ctx.JSON(http.StatusOK, pb.GetAllSkills{Skills: userSkills})
}

