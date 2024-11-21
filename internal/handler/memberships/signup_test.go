package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "example@mail.com",
					Username: "example",
					Password: "secret",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "example@mail.com",
					Username: "example",
					Password: "secret",
				}).Return(errors.New("username or email exists"))
			},
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			c := gin.New()
			h := &Handler{
				e:       c,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/sign-up`
			req := memberships.SignUpRequest{
				Email:    "example@mail.com",
				Username: "example",
				Password: "secret",
			}

			val, err := json.Marshal(req)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			newReq, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.e.ServeHTTP(w, newReq)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
