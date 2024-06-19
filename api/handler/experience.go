package handler

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Experience handles the creation of a new Experience
// @Summary 		Create Experience
// @Description 	Create page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Experience  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/experience/create [post]
func (h *Handler) CreateExperience(ctx *gin.Context) {
	var arr pb.Experience
	if err := ctx.BindJSON(&arr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", arr.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid start date: %s", err.Error())})
		return
	}

	endDate, err := time.Parse("2006-01-02", arr.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid end date: %s", err.Error())})
		return
	}

	arr.StartDate = startDate.Format("2006-01-02")
	arr.EndDate = endDate.Format("2006-01-02")

	_, err = h.Experience.CreateExperience(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateExperience 	handles the creation of a new Experience
// @Summary 		Update Experience
// @Description 	Update page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Experience ID"
// @Param   		Update  body    pb.Experience  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/experience/update/{id} [put]
func (h *Handler) UpdateExperience(ctx *gin.Context) {
	id := ctx.Param("id")
	arr := &pb.Experience{Id: id}
	if err := ctx.BindJSON(arr); err != nil {
		panic(err)
	}
	_, err := h.Experience.UpdateExperience(ctx, arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, "Success!!!")
}

// DeleteExperience 	handles the creation of a new Experience
// @Summary 		Delete Experience
// @Description 	Delete page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Experience ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/experience/delete/{id} [delete]
func (h *Handler) DeleteExperience(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Experience.DeleteExperience(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllExperience 	handles the creation of a new Experience
// @Summary 		GetAll Experience
// @Description 	GetAll page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Experience true  "Query parameter"
// @Success 		200  {object}  pb.GetAllExperiences  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/experience/getall [get]
func (h *Handler) GetAllExperience(ctx *gin.Context) {
	Experience := &pb.Experience{}
	// restoran.Address = ctx.Param("restoran")
	// Experience.Name = ctx.Param("name")

	res, err := h.Experience.GetAllExperience(ctx, Experience)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdExperience 	handles the creation of a new Experience
// @Summary 		GetById Experience
// @Description 	GetById page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Experience ID"
// @Success 		200 {object}  pb.Experience   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/experience/getbyid/{id} [get]
func (h *Handler) GetByIdExperience(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Experience.GetByIdExperience(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}


// GetByUserIdExperience handles retrieving Experiences by User ID
// @Summary 		Get Experience by User ID
// @Description 	Retrieve experience by user ID
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		user_id    path    string  true  "User ID"
// @Success 		200 {array}  pb.Experience   "Get Experience by User ID Successful"
// @Failure 		400 {string}  string 		"User ID is required"
// @Failure 		404 {string}  string 		"Experiences not found"
// @Failure 		500 {string}  string 		"Error while retrieving experiences"
// @Router 			/experience/byuser/{user_id} [get]
func (h *Handler) GetByUserIdExperience(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	if user_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	experience := &pb.Experience{}
	res, err := h.Experience.GetAllExperience(ctx, experience)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving experiences"})
		return
	}

	// Filter experiences by user_id
	var userExperiences []*pb.Experience
	for _, exp := range res.Experiences {
		if exp.UserId == user_id {
			userExperiences = append(userExperiences, exp)
		}
	}

	if len(userExperiences) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Experiences not found"})
		return
	}

	ctx.JSON(http.StatusOK, pb.GetAllExperiences{Experiences: userExperiences})
}