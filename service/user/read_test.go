package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/mars-terminal/mock-tests/entities/user"
	"github.com/mars-terminal/mock-tests/service"
)

func Test_CreateUserTest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := service.NewMockUserService(mockCtrl)

	testUser := &user.User{
		Id:       1,
		Name:     "Rust",
		Password: "password",
	}

	mockUserService.EXPECT().Create(gomock.Any(), "Rust", "password").Return(testUser, nil).Times(1)

	create, err := mockUserService.Create(context.Background(), "Rust", "password")
	if err != nil {
		return
	}

	require.Equal(t, create, testUser)
}
