package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain/mocks"
	"gitlab.com/alfred_soegiarto/training-clean-arch/user/usecase"
	"testing"
)

var testUser1 = domain.User{
	Id:     "1",
	Name:       "test1",
	Email:    "test1@gmail.com",
	Password:   "asdasd",
}

var testUser2 = domain.User{
	Id:     "2",
	Name:       "test2",
	Email:    "test2@gmail.com",
	Password:   "asdasd",
}


func TestUserUsecase_FetchAll(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockItemRepo := new(mocks.ItemRepository)
	// setup the repo mock
	mockUserRepo.On("FetchAll").Return([]domain.User{testUser1, testUser2}, nil)

	// Initiate and call the tested function
	testUserUsecase := usecase.NewUserUsecase(mockUserRepo, mockItemRepo)
	res, _ := testUserUsecase.FetchAll()

	// Assert Behaviour
	mockUserRepo.AssertExpectations(t)

	// Assert Data
	assert.Equal(t, "1", res[0].Id)
	assert.Equal(t, "test1", res[0].Name)
	assert.Equal(t, "test1@gmail.com", res[0].Email)
	assert.Equal(t, "asdasd", res[0].Password)

	assert.Equal(t, "2", res[1].Id)
	assert.Equal(t, "test2", res[1].Name)
	assert.Equal(t, "test2@gmail.com", res[1].Email)
	assert.Equal(t, "asdasd", res[1].Password)
}

