package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		dataLen int
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.dataLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		Top     int
		DataLen int
		Data    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Top:     tt.fields.Top,
				DataLen: tt.fields.DataLen,
				Data:    tt.fields.Data,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsFull(t *testing.T) {
	type fields struct {
		Top     int
		DataLen int
		Data    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Top:     tt.fields.Top,
				DataLen: tt.fields.DataLen,
				Data:    tt.fields.Data,
			}
			if got := s.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		Top     int
		DataLen int
		Data    []int
	}
	tests := []struct {
		name     string
		fields   fields
		wantData int
		wantOk   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Top:     tt.fields.Top,
				DataLen: tt.fields.DataLen,
				Data:    tt.fields.Data,
			}
			gotData, gotOk := s.Pop()
			if gotData != tt.wantData {
				t.Errorf("Pop() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Pop() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		Top     int
		DataLen int
		Data    []int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Top:     tt.fields.Top,
				DataLen: tt.fields.DataLen,
				Data:    tt.fields.Data,
			}
			if got := s.Push(tt.args.data); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}
