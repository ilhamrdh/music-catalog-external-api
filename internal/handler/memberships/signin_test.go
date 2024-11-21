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
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       response.Response
		wantErr            bool
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			mockFn: func() {
				mockSvc.EXPECT().SignIn(memberships.SignInRequest{
					Email:    "example@mail.com",
					Password: "secret",
				}).Return("accessToken", nil)
			},
			expectedBody: response.Response{
				Status:  http.StatusOK,
				Message: "Login successfully",
				Data: memberships.SignInResponse{
					AccessToken: "accessToken",
				},
			},
			wantErr: false,
		},
		{
			name:               "failed",
			expectedStatusCode: 400,
			mockFn: func() {
				mockSvc.EXPECT().SignIn(memberships.SignInRequest{
					Email:    "example@mail.com",
					Password: "secret",
				}).Return("", errors.New("invalid credentials"))
			},
			expectedBody: response.Response{
				Status: http.StatusBadRequest,
				Error:  "invalid credentials",
			},
			wantErr: true,
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

			var actual response.Response
			err = json.Unmarshal(w.Body.Bytes(), &actual)
			assert.NoError(t, err)

			if !tt.wantErr {
				assert.Equal(t, tt.expectedBody.Status, actual.Status)
				assert.Equal(t, tt.expectedBody.Message, actual.Message)

				expectedData := tt.expectedBody.Data.(memberships.SignInResponse)
				actualData := memberships.SignInResponse{}
				dataBytes, _ := json.Marshal(actual.Data)
				err = json.Unmarshal(dataBytes, &actualData)
				assert.NoError(t, err)
				assert.Equal(t, expectedData, actualData)
			} else {
				assert.Equal(t, tt.expectedBody.Status, actual.Status)
				assert.Equal(t, tt.expectedBody.Error, actual.Error)
			}
		})
	}
}
