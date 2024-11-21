package memberships

import (
	"testing"

	"github.com/ilhamrdh/music-catalog-external-api/internal/configs"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "example@gmail.com",
					Username: "usernametest",
					Password: "secret",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when GetUser",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "example@gmail.com",
					Username: "usernametest",
					Password: "secret",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when Signup",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "example@gmail.com",
					Username: "usernametest",
					Password: "secret",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_SignIn(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	type args struct {
		request memberships.SignInRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.SignInRequest{
					Email:    "test@mail.com",
					Password: "secret",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@mail.com",
					Password: "$2a$10$oHSoKSdP.cl6/5gr1rFMTeidsixABYrI8zJznj1E8bG.pX7jf11am",
					Username: "example",
				}, nil)
			},
		},
		{
			name: "error",
			args: args{
				request: memberships.SignInRequest{
					Email:    "test@mail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				request: memberships.SignInRequest{
					Email:    "test@mail.com",
					Password: "salah",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@mail.com",
					Password: "$2a$10$oHSoKSdP.cl6/5gr1rFMTeidsixABYrI8zJznj1E8bG.pX7jf11am",
					Username: "example",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "secret",
					},
				},
				repository: mockRepo,
			}
			got, err := s.SignIn(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
