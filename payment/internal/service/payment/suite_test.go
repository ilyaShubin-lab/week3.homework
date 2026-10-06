package payment

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.service = NewService()
}

func TestServiceSuite(t *testing.T) { suite.Run(t, new(ServiceSuite)) }
