package web

import (
	"GenealogyManagementSystem/internal/application/event"
	"GenealogyManagementSystem/internal/application/family_relation"
	"GenealogyManagementSystem/internal/application/media"
	"GenealogyManagementSystem/internal/application/person"
	"GenealogyManagementSystem/internal/application/relationship"
	"GenealogyManagementSystem/internal/infrastructure/web/handler"

	_ "GenealogyManagementSystem/cmd/server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 配置路由
func SetupRouter(
	personAppService *person.ApplicationService,
	relationshipAppService *relationship.ApplicationService,
	familyRelationAppService *family_relation.ApplicationService,
	eventAppService *event.ApplicationService,
	mediaAppService *media.ApplicationService,
) *gin.Engine {
	// 创建gin引擎
	r := gin.Default()

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	}))

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("swagger")))

	// 创建处理器实例
	personHandler := handler.NewPersonHandler(personAppService)
	relationshipHandler := handler.NewRelationshipHandler(relationshipAppService)
	familyRelationHandler := handler.NewFamilyRelationHandler(familyRelationAppService)
	eventHandler := handler.NewEventHandler(eventAppService)
	mediaHandler := handler.NewMediaHandler(mediaAppService)

	// API路由组
	api := r.Group("/api")
	{
		// 人员相关路由
		personGroup := api.Group("/person")
		{
			personGroup.POST("/create", personHandler.Create)
			personGroup.PUT("/update", personHandler.Update)
			personGroup.DELETE("/delete", personHandler.Delete)
			personGroup.GET("/getById", personHandler.GetByID)
			personGroup.GET("/getByUuid", personHandler.GetByUUID)
			personGroup.GET("/list", personHandler.List)
			personGroup.GET("/search", personHandler.Search)
		}

		// 亲子关系相关路由
		relationshipGroup := api.Group("/relationship")
		{
			relationshipGroup.POST("/createParentChild", relationshipHandler.CreateParentChild)
			relationshipGroup.DELETE("/deleteParentChild", relationshipHandler.DeleteParentChild)
			relationshipGroup.GET("/getParentChild", relationshipHandler.GetParentChild)
			relationshipGroup.GET("/getChildrenByParentId", relationshipHandler.GetChildrenByParentID)
			relationshipGroup.GET("/getParentsByChildId", relationshipHandler.GetParentsByChildID)
		}

		// 家庭关系相关路由
		familyRelationGroup := api.Group("/familyRelation")
		{
			familyRelationGroup.POST("/create", familyRelationHandler.Create)
			familyRelationGroup.PUT("/update", familyRelationHandler.Update)
			familyRelationGroup.DELETE("/delete", familyRelationHandler.Delete)
			familyRelationGroup.GET("/getById", familyRelationHandler.GetByID)
			familyRelationGroup.GET("/getByPersonId", familyRelationHandler.GetByPersonID)
			familyRelationGroup.GET("/getBySpouseId", familyRelationHandler.GetBySpouseID)
			familyRelationGroup.GET("/getByFamilyId", familyRelationHandler.GetByFamilyID)
		}

		// 事件相关路由
		eventGroup := api.Group("/event")
		{
			eventGroup.POST("/create", eventHandler.Create)
			eventGroup.PUT("/update", eventHandler.Update)
			eventGroup.DELETE("/delete", eventHandler.Delete)
			eventGroup.GET("/getById", eventHandler.GetByID)
			eventGroup.GET("/getByPersonId", eventHandler.GetByPersonID)
			eventGroup.GET("/getByType", eventHandler.GetByType)
			eventGroup.GET("/getImportant", eventHandler.GetImportant)
		}

		// 多媒体相关路由
		mediaGroup := api.Group("/media")
		{
			mediaGroup.POST("/create", mediaHandler.Create)
			mediaGroup.PUT("/update", mediaHandler.Update)
			mediaGroup.DELETE("/delete", mediaHandler.Delete)
			mediaGroup.GET("/getById", mediaHandler.GetByID)
			mediaGroup.GET("/getByPersonId", mediaHandler.GetByPersonID)
			mediaGroup.GET("/getByEventId", mediaHandler.GetByEventID)
			mediaGroup.GET("/getByType", mediaHandler.GetByType)
			mediaGroup.GET("/getPublic", mediaHandler.GetPublic)
		}
	}

	return r
}
