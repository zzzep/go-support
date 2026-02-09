package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/zzzep/go-support/bdd"
)

// Example interface for mock
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
	Delete(id int) error
}

type User struct {
	ID    int
	Name  string
	Email string
}

// Repository mock
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Save(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestExampleWithMocks demonstrates how to use mocks with the new Given, When, Then types
func TestExampleWithMocks(t *testing.T) {
	scenario := bdd.NewScenario()

	// Define steps using the new types
	// The mock will be created and configured within the Given steps
	scenario.Given(IHaveAUserRepository)
	scenario.Given(AUserExistsInRepository) // Automatically "AND"

	scenario.When(ISearchUserByID)
	scenario.When(IUpdateUser) // Automatically "AND"

	scenario.Then(UserWasFound)
	scenario.Then(UserWasUpdated)          // Automatically "AND"
	scenario.Then(MockExpectationsWereMet) // Automatically "AND"

	bdd.NewBDD(t).AddScenario(*scenario).Run()
}

// TestExampleMockWithAssertions demonstrates mock usage with assertions
func TestExampleMockWithAssertions(t *testing.T) {
	scenario := bdd.NewScenario()

	// The mock will be created and configured within the Given steps
	scenario.Given(IHaveAUserRepository)
	scenario.When(IDeleteUser)
	scenario.Then(UserWasDeleted)
	scenario.Then(MockExpectationsWereMet)

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// Example steps using mocks

func IHaveAUserRepository(sc bdd.Scenario, t *testing.T) {
	// Creates and adds the mock to the scenario within the Given step
	// Uses the interface as key: (*UserRepository)(nil)
	mockRepo := &MockUserRepository{}
	sc.AddMock((*UserRepository)(nil), mockRepo)
}

func AUserExistsInRepository(sc bdd.Scenario, t *testing.T) {
	// Gets the mock using the interface as key
	mockRepo := sc.Mocks[(*UserRepository)(nil)].(*MockUserRepository)

	// Configures the mock expectation
	user := &User{ID: 1, Name: "John", Email: "john@example.com"}
	mockRepo.On("FindByID", 1).Return(user, nil).Once()
}

func ISearchUserByID(sc bdd.Scenario, t *testing.T) {
	// Gets the mock using the interface as key
	mockRepo := sc.Mocks[(*UserRepository)(nil)].(*MockUserRepository)

	// Executes the action that uses the mock
	user, err := mockRepo.FindByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.Name)
}

func IUpdateUser(sc bdd.Scenario, t *testing.T) {
	// Gets the mock using the interface as key
	mockRepo := sc.Mocks[(*UserRepository)(nil)].(*MockUserRepository)

	// Configures expectation for update
	updatedUser := &User{ID: 1, Name: "John Smith", Email: "john.smith@example.com"}
	mockRepo.On("Save", updatedUser).Return(nil).Once()

	// Executes the update
	err := mockRepo.Save(updatedUser)
	assert.NoError(t, err)
}

func IDeleteUser(sc bdd.Scenario, t *testing.T) {
	// Gets the mock using the interface as key
	mockRepo := sc.Mocks[(*UserRepository)(nil)].(*MockUserRepository)

	// Configures expectation for deletion within the When step
	mockRepo.On("Delete", 1).Return(nil).Once()

	// Executes the deletion
	err := mockRepo.Delete(1)
	assert.NoError(t, err)
}

func UserWasFound(sc bdd.Scenario, t *testing.T) {
	// Verification already done in the When step
	assert.True(t, true)
}

func UserWasUpdated(sc bdd.Scenario, t *testing.T) {
	// Verification already done in the When step
	assert.True(t, true)
}

func UserWasDeleted(sc bdd.Scenario, t *testing.T) {
	// Verification already done in the When step
	assert.True(t, true)
}

func MockExpectationsWereMet(sc bdd.Scenario, t *testing.T) {
	// Gets the mock using the interface as key
	mockRepo := sc.Mocks[(*UserRepository)(nil)].(*MockUserRepository)

	// Verifies that all mock expectations were met
	mockRepo.AssertExpectations(t)
}
