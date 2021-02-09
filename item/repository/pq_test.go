package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/alfred_soegiarto/training-clean-arch/domain"
	itemRepo "gitlab.com/alfred_soegiarto/training-clean-arch/item/repository"
	"log"
	"regexp"
	"testing"
)

var testModelItem = &domain.Item{
	Id: "1",
	User_id: "1",
	App_name: "Zoom",
	App_email: "zoomEmail@gmail.com",
	App_password: "zoomPass",
}

var testModelItem2 = &domain.Item{
	Id: "2",
	User_id: "1",
	App_name: "Google",
	App_email: "googleEmail@gmail.com",
	App_password: "googlePass",
}

var testUserId = 1

func NewItemRepoTesting(dbConn *sql.DB) domain.ItemRepository{
	newRepo := itemRepo.NewPostgreItemRepository(dbConn)
	return newRepo
}

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error opening db stub: %s", err)
	}
	return db, mock
}

func TestFetchAllItem(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewItemRepoTesting(db)

	query := "SELECT id, user_id, app_name, app_email, app_password FROM items WHERE $1"
	rows := sqlmock.NewRows([]string{"id", "User_id", "app_ame", "app_email", "app_password"}).
		AddRow(testModelItem.Id, testModelItem.User_id, testModelItem.App_name, testModelItem.App_email, testModelItem.App_password)
	rows.AddRow(testModelItem.Id, testModelItem2.User_id, testModelItem2.App_name, testModelItem2.App_email, testModelItem2.App_password)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testUserId).WillReturnRows(rows)

	res, err := repo.FetchAll(int64(testUserId))
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, testModelItem.App_name, res[0].App_name)
	assert.Equal(t, testModelItem2.App_name, res[1].App_name)
}

func TestGetItemByName(t *testing.T) {
	db, mock := NewDBMock()
	repo := NewItemRepoTesting(db)

	query := `SELECT items.id, user_id, app_name, app_email, app_password 
								FROM items JOIN users 
								ON items.user_id = users.id
								WHERE users.id = $1 
								AND LOWER(items.app_name) = LOWER($2)`

	rows := sqlmock.NewRows([]string{"id", "user_id", "app_name", "app_email", "app_password"}).
		AddRow(testModelItem.Id, testModelItem.User_id, testModelItem.App_name, testModelItem.App_email, testModelItem.App_password)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testUserId, testModelItem.App_name).WillReturnRows(rows)

	res, err := repo.GetByName(int64(testUserId), testModelItem.App_name)
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, testModelItem.App_email, res.App_email)
}
