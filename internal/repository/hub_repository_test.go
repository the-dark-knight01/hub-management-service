package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HubRepositoryTestSuite struct {
	suite.Suite
	DB      *gorm.DB
	HubRepo HubRepository
}

func (suite *HubRepositoryTestSuite) SetupTest() {
	// Create an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}
	suite.DB = db

	// Auto-migrate the Hub entity
	suite.DB.AutoMigrate(&entity.Hub{}, &entity.Team{}) // Add all related entities here

	// Initialize the HubRepository
	suite.HubRepo = NewHubRepository(suite.DB)
}

func (suite *HubRepositoryTestSuite) TearDownTest() {
	// Clean up the database
	suite.DB.Exec("DELETE FROM hubs")
}

func (suite *HubRepositoryTestSuite) TestCreateHub() {
	hub := &entity.Hub{Name: "Test Hub"}

	// Create a hub
	err := suite.HubRepo.Create(hub)

	// Assert no error and the hub is saved
	assert.NoError(suite.T(), err)

	// Fetch the hub by ID
	fetchedHub, err := suite.HubRepo.FindByID(hub.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Test Hub", fetchedHub.Name)
}

func (suite *HubRepositoryTestSuite) TestFindAllHubs() {
	hub1 := &entity.Hub{Name: "Hub 1"}
	hub2 := &entity.Hub{Name: "Hub 2"}

	// Create hubs
	suite.HubRepo.Create(hub1)
	suite.HubRepo.Create(hub2)

	// Fetch all hubs
	hubs, err := suite.HubRepo.FindAll()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), hubs, 2)
}

func (suite *HubRepositoryTestSuite) TestSearchHubByName() {
	hub := &entity.Hub{Name: "Test Hub"}

	// Create hub
	suite.HubRepo.Create(hub)

	// Search for hubs by name
	hubs, err := suite.HubRepo.SearchByName("Test")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), hubs, 1)
	assert.Equal(suite.T(), "Test Hub", hubs[0].Name)
}

func TestHubRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(HubRepositoryTestSuite))
}
