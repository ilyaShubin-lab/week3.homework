package v1

import (
	"testing"

	"boilerplates/payment/internal/service/mocks"
	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	suite.Suite

	paymentService *mocks.MockPaymentService
	api            *api
}

func (s *APISuite) SetupTest() {
	s.paymentService = mocks.NewMockPaymentService(s.T())
	s.api = NewAPI(s.paymentService)
}

func TestServiceSuite(t *testing.T) { suite.Run(t, new(APISuite)) }
