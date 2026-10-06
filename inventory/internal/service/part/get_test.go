package part

import (
	"boilerplates/inventory/internal/model"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestGetSuccess() {
	partUUID := "part-1"
	expected := model.Part{
		UUID:  partUUID,
		Name:  "Супер пупер двигатель",
		Price: 22233,
	}

	s.partRepository.EXPECT().
		Get(mock.Anything, partUUID).
		Return(expected, nil)

	part, err := s.service.Get(s.T().Context(), partUUID)

	s.Require().NoError(err)
	s.Require().Equal(expected, part)
}
