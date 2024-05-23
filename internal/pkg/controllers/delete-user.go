package controllers

import "github.com/gin-gonic/gin"

type DeleteUserRequest struct {
	Id uint `json:"id"`
}

type DeleteUserResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (s *Server) DeleteUser(c *gin.Context) {

}
