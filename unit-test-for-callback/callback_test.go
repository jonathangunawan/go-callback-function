package main

import (
	"fmt"
	"testing"

	"github.com/jonathangunawan/go-callback-function/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain_Process(t *testing.T) {
	type args struct {
		num1     int
		num2     int
		callback func(int, int) (int, error)
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: args{
				num1: 2,
				num2: 5,
				callback: func(i1, i2 int) (int, error) {
					return i1 + i2, nil
				},
			},
			want:    7,
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				num1: 2,
				num2: 5,
				callback: func(i1, i2 int) (int, error) {
					return 0, fmt.Errorf("error")
				},
			},
			want:    0,
			wantErr: fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Main{}

			got, err := m.Process(tt.args.num1, tt.args.num2, tt.args.callback)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestSub_AnotherProcess(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: args{
				num1: 2,
				num2: 5,
			},
			want:    7,
			wantErr: nil,
		},
		{
			name: "success",
			args: args{
				num1: 10,
				num2: 5,
			},
			want:    0,
			wantErr: fmt.Errorf("more than 10"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sub{}

			got, err := s.AnotherProcess(tt.args.num1, tt.args.num2)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUsecase_SomeService(t *testing.T) {
	type args struct {
		data1 int
		data2 int
	}
	tests := []struct {
		name    string
		args    args
		mock    func(*mocks.MainItf, *mocks.SubItf)
		wantErr error
	}{
		{
			name: "success",
			args: args{
				data1: 2,
				data2: 5,
			},
			mock: func(mi *mocks.MainItf, si *mocks.SubItf) {
				si.On("AnotherProcess", 2, 5).Return(7, nil)
				mi.On("Process",
					2,
					5,
					mock.MatchedBy(func(fn func(int, int) (int, error)) bool {
						fn(2, 5)
						return true
					}),
				).Return(7, nil)
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				data1: 10,
				data2: 5,
			},
			mock: func(mi *mocks.MainItf, si *mocks.SubItf) {
				si.On("AnotherProcess", 10, 5).Return(0, fmt.Errorf("more than 10"))
				mi.On("Process",
					10,
					5,
					mock.MatchedBy(func(fn func(int, int) (int, error)) bool {
						_, err := fn(10, 5)
						return err.Error() == "more than 10"
					}),
				).Return(7, fmt.Errorf("more than 10")).Once()
			},
			wantErr: fmt.Errorf("wrong"),
		},
		{
			name: "error 2",
			args: args{
				data1: 10,
				data2: 5,
			},
			mock: func(mi *mocks.MainItf, si *mocks.SubItf) {
				si.On("AnotherProcess", 10, 5).Return(0, fmt.Errorf("more than 10"))
				mi.On("Process",
					10,
					5,
					mock.MatchedBy(func(fn func(int, int) (int, error)) bool {
						_, err := fn(10, 5)
						return err != nil
					}),
				).Return(7, fmt.Errorf("more than 10"))
			},
			wantErr: fmt.Errorf("wrong"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := mocks.NewMainItf(t)
			sm := mocks.NewSubItf(t)

			u := Usecase{
				m: mm,
				s: sm,
			}

			tt.mock(mm, sm)

			err := u.SomeService(tt.args.data1, tt.args.data2)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
