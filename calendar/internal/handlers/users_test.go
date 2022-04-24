package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	todo "http/internal/models"
	"http/internal/repository"
	service_mocks "http/internal/repository/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Login(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockStorage, user todo.User)
	user := todo.User{
		Id:       1,
		Username: "test",
		Password: "12345",
		Email:    "test@gmail.com",
		Timezone: "Europe/Kiev",
		Token:    "token",
	}
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            todo.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Wrong password",
			inputBody: `{"username":"test","password":"123456"}`,
			inputUser: user,
			mockBehavior: func(r *service_mocks.MockStorage, user todo.User) {
				r.EXPECT().GetUserByUsername("test").Return(user, nil)
			},
			expectedStatusCode:   403,
			expectedResponseBody: "wrong password",
		},
		{
			name:      "Ok",
			inputBody: `{"username":"servitpol","password":"12345"}`,
			inputUser: todo.User{
				Username: "servitpol",
				Password: "12345",
			},
			mockBehavior: func(r *service_mocks.MockStorage, user todo.User) {
				r.EXPECT().GetUserByUsername("servitpol").Return(user, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"token"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockStorage(c)
			test.mockBehavior(repo, test.inputUser)

			services := *repository.NewStorage(repo)
			handler := Handler{services}

			r := http.HandlerFunc(handler.Login)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/login",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)

		})
	}
}
