package e2e

import (
	"github.com/dan-kc/go-rest-api/packages/initializers"
	"github.com/dan-kc/go-rest-api/packages/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"testing"
)

type GetPostSuite struct {
	suite.Suite
}

func TestGetPostSuite(t *testing.T) {
	suite.Run(t, new(GetPostSuite))
}

func (s *GetPostSuite) TestGetNonExistingPost() {
	c := http.Client{}
	r, err := c.Get("http://localhost:3000/post/999999999999")

	s.NoError(err)
	s.Equal(http.StatusNotFound, r.StatusCode)

	body, err2 := io.ReadAll(r.Body)
	s.NoError(err2)
	s.JSONEq(`{"message":"No post exists with an ID of 999999999999"}`, string(body))
}

func (s *GetPostSuite) TestGetExistingPost() {
	c := http.Client{}

	post := models.Post{
		Model: gorm.Model{ID: 18},
		Title: "e2e Test",
		Body:  "GetPOSTTT",
	}

	dsn := os.Getenv("DB_URL")
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.Post{})
	DB.Create(&post)

	r, err := c.Get("http://localhost:3000/post/18")

	s.NoError(err)
	s.Equal(http.StatusOK, r.StatusCode)

	body, err2 := io.ReadAll(r.Body)
	s.NoError(err2)
	s.JSONEq(`{"message":"No post exists with an ID of 999999999999"}`, string(body))

	initializers.DB.Unscoped().Delete(&post)
}
