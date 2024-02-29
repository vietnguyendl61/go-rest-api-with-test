package e2e_test

import (
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestHappyHealthCheck() {
	c := http.Client{}

	response, _ := c.Get("http://localhost:8001/health-check")

	s.Equal(http.StatusOK, response.StatusCode)

	expectJson := `{"status": "OK", "message":""}`
	body, _ := io.ReadAll(response.Body)
	s.JSONEq(expectJson, string(body))
}
