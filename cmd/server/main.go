package main

import (
	"GenealogyManagementSystem/internal/application/event"
	"GenealogyManagementSystem/internal/application/family_relation"
	"GenealogyManagementSystem/internal/application/media"
	"GenealogyManagementSystem/internal/application/person"
	"GenealogyManagementSystem/internal/application/relationship"
	eventDomain "GenealogyManagementSystem/internal/domain/event"
	familyDomain "GenealogyManagementSystem/internal/domain/family_relation"
	mediaDomain "GenealogyManagementSystem/internal/domain/media"
	personDomain "GenealogyManagementSystem/internal/domain/person"
	relationDomain "GenealogyManagementSystem/internal/domain/relationship"
	"GenealogyManagementSystem/internal/infrastructure/persistence/sqlite"
	"GenealogyManagementSystem/internal/infrastructure/web"
	"fmt"
	"log"
)

// @title 族谱管理系统 API
// @version 1.0
// @description 族谱管理系统的RESTful API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	fmt.Println("Genealogy Management System Server")
	log.Println("Server starting...")

	// 初始化数据库连接
	err := sqlite.InitDB("genealogy.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer sqlite.CloseDB()

	// 创建仓储实例
	personRepo := sqlite.NewPersonRepository(sqlite.DB)
	relationshipRepo := sqlite.NewRelationshipRepository(sqlite.DB)
	familyRelationRepo := sqlite.NewFamilyRelationRepository(sqlite.DB)
	eventRepo := sqlite.NewEventRepository(sqlite.DB)
	mediaRepo := sqlite.NewMediaRepository(sqlite.DB)

	// 创建领域服务实例
	personService := personDomain.NewService(personRepo)
	relationshipService := relationDomain.NewService(relationshipRepo)
	familyRelationService := familyDomain.NewService(familyRelationRepo)
	eventService := eventDomain.NewService(eventRepo)
	mediaService := mediaDomain.NewService(mediaRepo)

	// 创建应用服务实例
	personAppService := person.NewApplicationService(personService)
	relationshipAppService := relationship.NewApplicationService(relationshipService)
	familyRelationAppService := family_relation.NewApplicationService(familyRelationService)
	eventAppService := event.NewApplicationService(eventService)
	mediaAppService := media.NewApplicationService(mediaService)

	// 设置路由
	router := web.SetupRouter(personAppService, relationshipAppService, familyRelationAppService, eventAppService, mediaAppService)

	// 启动服务器
	port := "8080"
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
