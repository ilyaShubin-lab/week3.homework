package part

import (
	"testing"

	"boilerplates/inventory/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	partRepository *mocks.MockPartRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.partRepository = mocks.NewMockPartRepository(s.T())

	s.service = NewService(
		s.partRepository,
	)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
