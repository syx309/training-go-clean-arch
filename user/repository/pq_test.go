package repository_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	userRepo "gitlab.com/alfred_soegiarto/training-clean-arch/user/repository"
	"log"
	"regexp"
	"strconv"
	"testing"
)

var testModel = &domain.User{
	Id:     "1",
	Name:       "test1",
	Email:    "test1@gmail.com",
	Password:   "asdasd",
}

var testModel2 = &domain.User{
	Id:     "2",
	Name:       "test2",
	Email:    "test2@gmail.com",
	Password:   "asdasd",
}

var testInsertModel = &domain.NewUser{
	Name:       "test3",
	Email:    	"test3@gmail.com",
	Password:   "asdasd",
}

func NewUserRepoTesting(dbConn *sql.DB) domain.UserRepository{
	newRepo := userRepo.NewPostgreUserRepository(dbConn)
	return newRepo
}

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error opening db stub: %s", err)
	}
	return db, mock
}

func TestFetchAll(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewUserRepoTesting(db)
	query := "SELECT name, email, password FROM users"
	rows := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow(testModel.Name,testModel.Email,testModel.Password)
	rows.AddRow(testModel2.Name, testModel2.Email, testModel2.Password)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	res, err := repo.FetchAll()
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, testModel.Name, res[0].Name)
	assert.Equal(t, testModel2.Name, res[1].Name)
}

func TestGetById(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewUserRepoTesting(db)

	query := "SELECT id, name, email, password FROM users WHERE id = $1"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(testModel.Id, testModel.Name, testModel.Email, testModel.Password)
	id, err := strconv.ParseInt(testModel.Id, 10, 64)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnRows(rows)

	res, err := repo.GetById(id)
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, testModel.Name, res.Name)
}

func TestInsert(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewUserRepoTesting(db)

	query := fmt.Sprintf("INSERT INTO users (name, email, password) VALUES ('%s', '%s', '%s')", testInsertModel.Name, testInsertModel.Email, testInsertModel.Password)

	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.Insert(testInsertModel)
	assert.NoError(t, err)
}



