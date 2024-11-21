package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignIn(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       memberships.SignInResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignIn(memberships.SignInRequest{
					Email:    "example@mail.com",
					Password: "secret",
				}).Return("accessToken", nil)
			},
			expectedStatusCode: 200,
			expectedBody: memberships.SignInResponse{
				AccessToken: "accessToken",
			},
			wantErr: false,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignIn(memberships.SignInRequest{
					Email:    "example@mail.com",
					Password: "secret",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 400,
			expectedBody:       memberships.SignInResponse{},
			wantErr:            true,
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
			endpoint := `/memberships/sign-in`

			req := memberships.SignInRequest{
				Email:    "example@mail.com",
				Password: "secret",
			}

			val, err := json.Marshal(req)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			newReq, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.e.ServeHTTP(w, newReq)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := memberships.SignInResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			} else {

			}

		})
	}
}
