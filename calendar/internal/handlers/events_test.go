package handlers

import (
	"bytes"
	todo "calendar/internal/models"
	"calendar/internal/repository"
	service_mocks "calendar/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateEvent(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockStorage, event todo.Event)
	event := todo.Event{
		Uid:           1,
		Title:         "test",
		Description:   "qwerty",
		TimestampFrom: 1641124800,
		TimestampTo:   1641297600,
	}
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
		inputEvent           todo.Event
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Create event",
			inputBody:   `{"title":"test", "description":"qwerty", "time_from": "02 Jan 22 15:00", "time_to":"04 Jan 22 15:00"}`,
			inputEvent:  event,
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(r *service_mocks.MockStorage, event todo.Event) {
				r.EXPECT().GetUserByToken("token").Return(user, nil)
				r.EXPECT().CreateEvent(event).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockStorage(c)
			test.mockBehavior(repo, test.inputEvent)

			services := *repository.NewStorage(repo)
			handler := Handler{services}

			r := http.HandlerFunc(handler.CreateEvent)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "api/events",
				bytes.NewBufferString(test.inputBody))
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)

		})
	}
}
