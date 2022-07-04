package handler

import "github.com/gin-gonic/gin"

// Handler interface
type Handler interface {
	HandleGreetings(context *gin.Context)
	HandleCapabilityCreation(context *gin.Context)
}
