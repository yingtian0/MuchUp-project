package rest

import (
	"MuchUp/app/internal/domain/entity"
	mock_usecase "MuchUp/app/internal/domain/usecase/mocks"
	logger "MuchUp/app/pkg/logger"
	"MuchUp/app/utils"
	"encoding/json"

	"io"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUsecase := mock_usecase.NewMockUserUsecase(ctrl)
	mockMessageUsecase := mock_usecase.NewMockMessageUsecase(ctrl)
	appLogger := logger.NewLogger()

	handler := NewHandler(mockUserUsecase, mockMessageUsecase, appLogger)

	reqbody := `{"name":"testname","email":"test@example.com","password":"testpasswordhash"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(reqbody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockUserUsecase.EXPECT().
		CreateUser(gomock.Any()).
		Return(&entity.User{
			ID:       "mock_ID",
			NickName: "testname",
			Email:    utils.StringPtr("test@example.com"),
		}, nil)

	handler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response Response
	bodyBytes, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(bodyBytes, &response)
	require.NoError(t, err)

	responseData, ok := response.Data.(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "test@example.com", responseData["Email"])
}
