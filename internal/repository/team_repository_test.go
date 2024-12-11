package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TeamRepositoryTestSuite struct {
	suite.Suite
	DB       *gorm.DB
	TeamRepo TeamRepository
}

func (suite *TeamRepositoryTestSuite) SetupTest() {
	// Create an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}
	suite.DB = db

	// Auto-migrate the Team entity
	suite.DB.AutoMigrate(&entity.Team{})

	// Initialize the TeamRepository
	suite.TeamRepo = NewTeamRepository(suite.DB)
}

func (suite *TeamRepositoryTestSuite) TearDownTest() {
	// Clean up the database
	suite.DB.Exec("DELETE FROM teams")
}

func (suite *TeamRepositoryTestSuite) TestCreateTeam() {
	team := &entity.Team{Name: "Team A", HubID: 1}

	// Create a team
	err := suite.TeamRepo.Create(team)

	// Assert no error and the team is saved
	assert.NoError(suite.T(), err)

	// Fetch the team by ID
	fetchedTeam, err := suite.TeamRepo.FindByID(team.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Team A", fetchedTeam.Name)
}

func (suite *TeamRepositoryTestSuite) TestFindTeamsByHubID() {
	team1 := &entity.Team{Name: "Team A", HubID: 1}
	team2 := &entity.Team{Name: "Team B", HubID: 1}

	// Create teams
	suite.TeamRepo.Create(team1)
	suite.TeamRepo.Create(team2)

	// Find teams by HubID
	teams, err := suite.TeamRepo.FindByHubID(1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), teams, 2)
}

func TestTeamRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TeamRepositoryTestSuite))
}
