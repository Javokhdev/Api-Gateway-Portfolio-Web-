package handler

import (
	"net/http"
	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// GetUserPortfolio handles retrieving the full user portfolio by User ID
// @Summary      Get User Portfolio
// @Description  Retrieve user portfolio by user ID
// @Tags         Portfolio
// @Accept       json
// @Produce      json
// @Param        user_id    path    string  true  "User ID"
// @Success      200 {object} pb.UserPortfolio   "Get User Portfolio Successful"
// @Failure      400 {string} string             "User ID is required"
// @Failure      500 {string} string             "Error while retrieving portfolio"
// @Router       /portfolio/{user_id} [get]
func (h *Handler) GetUserPortfolio(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Initialize the portfolio response
	portfolio := &pb.UserPortfolio{}

	// Retrieve Projects
	projectReq := &pb.Project{UserId: userID}
	projectRes, err := h.Project.GetAllProject(ctx, projectReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving projects"})
		return
	}
	portfolio.Projects = projectRes.Projects

	// Retrieve Skills
	skillReq := &pb.Skill{UserId: userID}
	skillRes, err := h.Skill.GetAllSkill(ctx, skillReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving skills"})
		return
	}
	portfolio.Skills = skillRes.Skills

	// Retrieve Experiences
	experienceReq := &pb.Experience{}
	experienceRes, err := h.Experience.GetAllExperience(ctx, experienceReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving experiences"})
		return
	}
	// Filter experiences by user ID
	for _, exp := range experienceRes.Experiences {
		if exp.UserId == userID {
			portfolio.Experiences = append(portfolio.Experiences, exp)
		}
	}

	// Retrieve Education
	educationReq := &pb.Education{}
	educationRes, err := h.Education.GetAllEducation(ctx, educationReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving education"})
		return
	}
	// Filter education by user ID
	for _, edu := range educationRes.Educations {
		if edu.UserId == userID {
			portfolio.Educations = append(portfolio.Educations, edu)
		}
	}

	// Return the aggregated portfolio
	ctx.JSON(http.StatusOK, portfolio)
}
