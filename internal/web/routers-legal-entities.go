package web

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krisch/crm-backend/internal/legalentities"
	"github.com/sirupsen/logrus"
)

// RegisterLegalEntityRoutes initializes the routes for legal entity operations.
func RegisterLegalEntityRoutes(router *gin.Engine, service *legalentities.Service) {
	group := router.Group("/federation/legal-entities")

	group.GET("/", getAllLegalEntitiesHandler(service))
	group.POST("/", createLegalEntityHandler(service))
	group.PUT("/:id", updateLegalEntityHandler(service))
	group.DELETE("/:id", deleteLegalEntityHandler(service))

}

func getAllLegalEntitiesHandler(service *legalentities.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		entities, err := service.GetAllLegalEntities(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to get all legal entities")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, entities)
	}
}

func createLegalEntityHandler(service *legalentities.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var entity legalentities.LegalEntity
		if err := c.ShouldBindJSON(&entity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		if err := service.CreateLegalEntity(ctx, entity); err != nil {
			logrus.WithError(err).Error("Failed to create legal entity")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, entity)
	}
}

func updateLegalEntityHandler(service *legalentities.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var entity legalentities.LegalEntity
		if err := c.ShouldBindJSON(&entity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		entity.ID = id
		if err := service.UpdateLegalEntity(ctx, id, entity); err != nil {
			logrus.WithError(err).Error("Failed to update legal entity")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, entity)
	}
}

func deleteLegalEntityHandler(service *legalentities.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ctx := c.Request.Context()
		if err := service.DeleteLegalEntity(ctx, id); err != nil {
			logrus.WithError(err).Error("Failed to delete legal entity")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}




