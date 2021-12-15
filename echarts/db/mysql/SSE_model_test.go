package mysql

import (
	"wakever/charts/db/mysql/SSE"

	"github.com/kelseyhightower/confd/log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user *mysql.SSE
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{user: &mysql.SSE{SSE: "test"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
	}{
		{
			name: "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := GetAllSSEs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, user := range gotUsers {
				log.Info("user: %v", user)
			}

		})
	}
}

func TestUpdateUserNameById(t *testing.T) {
	type args struct {
		SSE string
		userId   int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				SSE: "test",
				userId:   10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserNameById(tt.args.SSE, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserNameById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByNameAndPassword(t *testing.T) {
	type args struct {
		SSE     string
		Id      int64
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
	}{
		{
			name: "test",
			args: args{
				SSE:     "Klein",
				Id:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//tt.args.Id
			gotUser, err := GetUserBySSEAndId(tt.args.SSE,tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByNameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Info("user %v", gotUser)
		})
	}
}