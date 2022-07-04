package handler

import (
	"base-app/datamodels"
	"base-app/driver"
	"base-app/pkg/dto"
	"base-app/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type capabilityHandler struct {
	Logger zap.Logger
}

func NewCapabilityHandler() *capabilityHandler {

	return &capabilityHandler{
		Logger: *utils.Logger,
	}
}

func (handler *capabilityHandler) HandleCapabilityCreation(contxt *gin.Context) {

	fmt.Println(contxt.Request.Body)

	var requestBody dto.Capability
	if err := contxt.ShouldBindJSON(&requestBody); err != nil {

		contxt.AbortWithStatus(400)
		handler.Logger.Error(err.Error())
		return
	}
	// handler.Logger.Info("request", zap.Any("body", requestBody))
	collection := driver.DbClient.Database("api-management").Collection("capability")
	dbctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	capabilityModel := datamodels.Capability{
		CapabilityName: requestBody.CapabilityName,
		Description:    requestBody.Description,
		Version:        requestBody.Version,
		TenantName:     requestBody.TenantName,
		ContextPath:    requestBody.ContextPath,
		HealthCheckURL: requestBody.HealthCheckURL,
		OpenAPISpecs:   requestBody.OpenAPISpecs,
		Category:       requestBody.Category,
		EnableCORS:     requestBody.EnableCors,
		Labels:         requestBody.Labels,
		Authorization:  requestBody.Authorization,
		Owner:          requestBody.Owner,
		Consumers:      requestBody.Consumers,
	}

	res, insertErr := collection.InsertOne(dbctx, capabilityModel)
	handler.Logger.Info("db res", zap.Any("res", res), zap.Any("err", insertErr))

	filter := bson.D{{"context_path", capabilityModel.ContextPath}}
	var findRes *datamodels.Capability

	if findErr := collection.FindOne(dbctx, filter).Decode(&findRes); findErr != nil {
		log.Println(findErr)
		contxt.AbortWithStatus(500)
		return
	}
	// var target []datamodels.Capability
	// marshErr := bson.Unmarshal([]byte(findRes), &target)
	fmt.Println(findRes)
	contxt.Status(200)
}
