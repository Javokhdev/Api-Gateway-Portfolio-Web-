package handler

import (
	"net/http"
	"strings"

	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Project handles the creation of a new Project
// @Summary 		Create Project
// @Description 	Create page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Project  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/project/create [post]
func (h *Handler) CreateProject(ctx *gin.Context) {
	arr := pb.Project{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Project.CreateProject(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateProject 	handles the creation of a new Project
// @Summary 		Update Project
// @Description 	Update page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Project ID"
// @Param   		Update  body    pb.Project  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/project/update/{id} [put]
func (h *Handler) UpdateProject(ctx *gin.Context) {
	arr := pb.Project{}
	id := ctx.Param("id")
	arr.Id = id
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Project.UpdateProject(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteProject 	handles the creation of a new Project
// @Summary 		Delete Project
// @Description 	Delete page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Project ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/project/delete/{id} [delete]
func (h *Handler) DeleteProject(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Project.DeleteProject(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllProject 	handles the creation of a new Project
// @Summary 		GetAll Project
// @Description 	GetAll page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Project true  "Query parameter"
// @Success 		200  {object}  pb.GetAllProjects  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/project/getall [get]
func (h *Handler) GetAllProject(ctx *gin.Context) {
	Project := &pb.Project{}
	Project.Id = ctx.Query("id")
	Project.UserId = ctx.Query("user_id")
	Project.Title = ctx.Query("title")
	Project.Description = ctx.Query("description")
	Project.Url = ctx.Query("url")

	res, err := h.Project.GetAllProject(ctx, Project)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdProject 	handles the creation of a new Project
// @Summary 		GetById Project
// @Description 	GetById page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Project ID"
// @Success 		200 {object}  pb.Project   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/project/getbyid/{id} [get]
func (h *Handler) GetByIdProject(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Project.GetByIdProject(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByUserIdProject handles retrieving Projects by User ID
// @Summary 		Get Projects by User ID
// @Description 	Retrieve projects by user ID
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		user_id    query   string  true  "User ID"
// @Success 		200 {object}  pb.GetAllProjects  	"Get Projects by User ID Successful"
// @Failure 		400 {string}  string 		"User ID is required"
// @Failure 		404 {string}  string 		"User not found"
// @Failure 		500 {string}  string 		"Error while retrieving projects"
// @Router 			/project/byuser [get]
func (h *Handler) GetByUserIdProject(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	if user_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	Project := &pb.Project{}
	res, err := h.Project.GetAllProject(ctx, Project)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving projects"})
		return
	}

	// Filter projects by user_id
	var userProjects []*pb.Project
	for _, project := range res.Projects {
		if project.UserId == user_id {
			userProjects = append(userProjects, project)
		}
	}

	if len(userProjects) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, pb.GetAllProjects{Projects: userProjects})
}

// SearchProjects handles searching projects by name or description
// @Summary      Search Projects
// @Description  Search projects by name or description
// @Tags         Project
// @Accept       json
// @Produce      json
// @Param        query   query    string  true  "Search Query"
// @Success      200     {object} pb.GetAllProjects   "Search Successful"
// @Failure      400     {string} string              "Search query is required"
// @Failure      500     {string} string              "Error while searching projects"
// @Router       /project/search [get]
func (h *Handler) SearchProjects(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Call GetAllProject to get all projects
	allProjects, err := h.Project.GetAllProject(ctx, &pb.Project{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving projects"})
		return
	}

	// Filter projects based on the search query
	var filteredProjects []*pb.Project
	for _, project := range allProjects.Projects {
		if strings.Contains(strings.ToLower(project.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(project.Description), strings.ToLower(query)) {
			filteredProjects = append(filteredProjects, project)
		}
	}

	// Prepare the response
	response := &pb.GetAllProjects{Projects: filteredProjects}

	ctx.JSON(http.StatusOK, response)
}
