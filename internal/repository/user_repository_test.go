package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hub_management_service/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DB       *gorm.DB
	UserRepo UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	// Create an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}
	suite.DB = db

	// Auto-migrate the User entity
	suite.DB.AutoMigrate(&entity.User{})

	// Initialize the UserRepository
	suite.UserRepo = NewUserRepository(suite.DB)
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// Clean up the database
	suite.DB.Exec("DELETE FROM users")
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := &entity.User{Name: "User 1", TeamID: 1}

	// Create a user
	err := suite.UserRepo.Create(user)

	// Assert no error and the user is saved
	assert.NoError(suite.T(), err)

	// Fetch the user by ID
	fetchedUser, err := suite.UserRepo.FindByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "User 1", fetchedUser.Name)
}

func (suite *UserRepositoryTestSuite) TestFindUserByTeamID() {
	user1 := &entity.User{Name: "User 1", TeamID: 1}
	user2 := &entity.User{Name: "User 2", TeamID: 1}

	// Create users
	suite.UserRepo.Create(user1)
	suite.UserRepo.Create(user2)

	// Find users by TeamID
	users, err := suite.UserRepo.FindUserByTeamID(1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 2)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
