package handler

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Education handles the creation of a new Porfolio
// @Summary 		Create Porfolio
// @Description 	Create page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Education  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/education/create [post]
func (h *Handler) CreateEducation(ctx *gin.Context) {
	education := pb.Education{}
	if err := ctx.BindJSON(&education); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", education.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid start date: %s", err.Error())})
		return
	}

	endDate, err := time.Parse("2006-01-02", education.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid end date: %s", err.Error())})
		return
	}

	education.StartDate = startDate.Format("2006-01-02")
	education.EndDate = endDate.Format("2006-01-02")

	_, err = h.Education.CreateEducation(ctx, &education)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// UpdateEducation 	handles the creation of a new Education
// @Summary 		Update Education
// @Description 	Update page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Education ID"
// @Param   		Update  body    pb.Education  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/education/update/{id} [put]
func (h *Handler) UpdateEducation(ctx *gin.Context) {
	arr := pb.Education{}
	id := ctx.Param("id")
	arr.Id = id
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Education.UpdateEducation(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteEducation 	handles the creation of a new Education
// @Summary 		Delete Education
// @Description 	Delete page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Education ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/education/delete/{id} [delete]
func (h *Handler) DeleteEducation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Education.DeleteEducation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllEducation 	handles the creation of a new Education
// @Summary 		GetAll Education
// @Description 	GetAll page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Education true  "Query parameter"
// @Success 		200  {object}  pb.GetAllEducations  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/education/getall [get]
func (h *Handler) GetAllEducation(ctx *gin.Context) {
	Education := &pb.Education{}
	Education.Id = ctx.Query("id")
	Education.UserId = ctx.Query("user_id")
	Education.Institution = ctx.Query("institution")
	Education.Degree = ctx.Query("degree")
	Education.FieldOfStudy = ctx.Query("field_of_study")
	Education.StartDate = ctx.Query("start_date")
	Education.EndDate = ctx.Query("end_date")

	res, err := h.Education.GetAllEducation(ctx, Education)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdEducation 	handles the creation of a new Education
// @Summary 		GetById Education
// @Description 	GetById page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Education ID"
// @Success 		200 {object}  pb.Education   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/education/getbyid/{id} [get]
func (h *Handler) GetByIdEducation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Education.GetByIdEducation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByUserIdEducation handles retrieving Educations by User ID
// @Summary 		Get Education by User ID
// @Description 	Retrieve educatioin by user ID
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		user_id    path    string  true  "User ID"
// @Success 		200 {array}  pb.Education  "Get Education by User ID Successful"
// @Failure 		400 {string}  string 		"User ID is required"
// @Failure 		404 {string}  string 		"Educations not found"
// @Failure 		500 {string}  string 		"Error while retrieving educations"
// @Router 			/education/byuser/{user_id} [get]
func (h *Handler) GetByUserIdEducation(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	if user_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	education := &pb.Education{}
	res, err := h.Education.GetAllEducation(ctx, education)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving educations"})
		return
	}

	// Filter education by user_id
	var userEducations []*pb.Education
	for _, exp := range res.Educations {
		if exp.UserId == user_id {
			userEducations = append(userEducations, exp)
		}
	}

	if len(userEducations) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Educations not found"})
		return
	}

	ctx.JSON(http.StatusOK, pb.GetAllEducations{Educations: userEducations})
}