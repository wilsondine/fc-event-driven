package create_account_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yansb/full-cycle-course/walletcore/internal/entity"
	"github.com/yansb/full-cycle-course/walletcore/internal/use_case/create_account"
	"github.com/yansb/full-cycle-course/walletcore/internal/use_case/mocks"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := create_account.NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := create_account.CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
