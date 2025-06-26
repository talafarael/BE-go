package user_postgres_test

import (
	"database/sql"
	"gin/internal/user/user_repository/user_postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: sqldb,
			}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestGetUserByEamil(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	repo := user_postgres.NewUserRepo(db)

	// user
	userRows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "John Doe", "john@example.com", "pass")
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email=\$1 ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs("john@example.com", 1).
		WillReturnRows(userRows)

	// vacancy
	vacancyRows := sqlmock.NewRows([]string{"id", "user_id", "title"}).
		AddRow(1, 1, "Test Vacancy")
	mock.ExpectQuery(`SELECT \* FROM "vacancies" WHERE "vacancies"\."user_id" = \$1`).
		WithArgs(1).
		WillReturnRows(vacancyRows)

	user, err := repo.GetUserByEmail("john@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}

//func TestGetUserByID(t *testing.T) {
//	sqlDB, db, mock := DbMock(t)
//	defer sqlDB.Close()
//
//	repo := user_postgres.NewUserRepo(db)
//
//	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
//		AddRow(1, "John Doe", "john@example.com", "pass")
//
//	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."id" = \$1 ORDER BY "users"\."id" LIMIT \$2`).
//		WithArgs(1, 1).
//		WillReturnRows(rows)
//
//	user, err := repo.GetUserByID(1)
//
//	assert.NoError(t, err)
//	assert.Equal(t, "John Doe", user.Name)
//	assert.Equal(t, "john@example.com", user.Email)
//}

//func TestCreateUser_Success(t *testing.T) {
//	sqlDB, db, mock := DbMock(t)
//	defer sqlDB.Close()
//
//	repo := user_postgres.NewUserRepo(db)
//	userDto := &auth_dto.RegisterDto{
//		Name:     "John",
//		Email:    "john@example.com",
//		Password: "hashed-password",
//	}
//
//	mock.ExpectBegin()
//	mock.ExpectQuery(`INSERT INTO "users"`).
//		WithArgs(userDto.Name, userDto.Email, userDto.Password).
//		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
//	mock.ExpectCommit()
//	createdUser, err := repo.CreateUser(userDto)
//
//	assert.NoError(t, err)
//	assert.Equal(t, userDto.Name, createdUser.Name)
//	assert.Equal(t, userDto.Email, createdUser.Email)
//	assert.Equal(t, userDto.Password, createdUser.Password)
//}
//
//func TestCreateUser_Failure(t *testing.T) {
//	sqlDB, db, mock := DbMock(t)
//	defer sqlDB.Close()
//	repo := user_postgres.NewUserRepo(db)
//
//	userDto := &auth_dto.RegisterDto{
//		Name:     "Jane",
//		Email:    "jane@example.com",
//		Password: "hashed-password",
//	}
//
//	mock.ExpectBegin()
//	mock.ExpectQuery(`INSERT INTO "users"`).
//		WithArgs(userDto.Name, userDto.Email, userDto.Password).
//		WillReturnError(sqlmock.ErrCancelled)
//	mock.ExpectRollback()
//
//	_, err := repo.CreateUser(userDto)
//
//	assert.Error(t, err)
//	assert.Equal(t, response_error.ErrUserAlredy, err)
//}
