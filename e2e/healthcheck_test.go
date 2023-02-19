package e2e

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type End2EndSuite struct {
	suite.Suite
}

func TestEnd2EndSuite(t *testing.T) {
	suite.Run(t, new(End2EndSuite))
}

func (s *End2EndSuite) TestHealthcheck() {
	c := http.Client{}
	r, err := c.Get("http://localhost:3000/healthcheck")

	s.NoError(err)
	s.Equal(http.StatusOK, r.StatusCode)
}
