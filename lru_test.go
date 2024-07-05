package lru

import (
	"container/list"
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		{
			name: "Init new LRU cache struct",
			args: args{cap: 3},
			want: NewLRUCache(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLRUCache(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRU_Add(t *testing.T) {
	type fields struct {
		Queue    *list.List
		Items    map[string]*Node
		Capacity int
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Succesful adding value",
			fields: fields{
				Queue:    list.New(),
				Items:    make(map[string]*Node),
				Capacity: 1,
			},
			args: args{key: "key", value: "value"},
			want: true,
		},
		{
			name: "Dublicate key",
			fields: fields{
				Queue: list.New(),
				Items: map[string]*Node{
					"key": {
						Data:   "value",
						KeyPrt: nil,
					},
				},
				Capacity: 1},
			args: args{key: "key", value: "value"},
			want: false,
		},
		{
			name: "Capacity overflow",
			fields: fields{
				Queue: list.New(),
				Items: map[string]*Node{
					"old": {
						Data:   "old",
						KeyPrt: nil,
					},
				},
				Capacity: 1,
			},
			args: args{key: "key", value: "value"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LRU{
				Queue:    tt.fields.Queue,
				Items:    tt.fields.Items,
				Capacity: tt.fields.Capacity,
			}
			if got := c.Add(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("LRU.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRU_Get(t *testing.T) {
	type fields struct {
		Capacity int
	}
	type args struct {
		key string
	}
	type initValue struct {
		key   string
		value string
	}
	tests := []struct {
		name       string
		fields     fields
		initValues []initValue
		args       args
		wantValue  string
		wantOk     bool
	}{
		{
			name: "Value is exists",
			fields: fields{
				Capacity: 2,
			},
			initValues: []initValue{
				{
					key:   "key1",
					value: "value1",
				},
				{
					key:   "key2",
					value: "value2",
				},
			},
			args: args{
				key: "key1",
			},
			wantValue: "value1",
			wantOk:    true,
		},
		{
			name: "Value if not exists",
			fields: fields{
				Capacity: 1,
			},
			initValues: []initValue{
				{
					key:   "key1",
					value: "value1",
				},
			},
			args: args{
				key: "key2",
			},
			wantValue: "",
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewLRUCache(tt.fields.Capacity)
			for _, initValue := range tt.initValues {
				c.Add(initValue.key, initValue.value)
			}

			gotValue, gotOk := c.Get(tt.args.key)
			if gotValue != tt.wantValue {
				t.Errorf("LRU.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("LRU.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestLRU_Remove(t *testing.T) {
	type fields struct {
		Queue    *list.List
		Items    map[string]*Node
		Capacity int
	}
	type initValue struct {
		key   string
		value string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		initValues []initValue
		args       args
		wantOk     bool
	}{
		{
			name: "Deleting key is exitsts",
			fields: fields{
				Capacity: 1,
			},
			initValues: []initValue{
				{
					key:   "key1",
					value: "value1",
				},
			},
			args: args{
				key: "key1",
			},
			wantOk: true,
		},
		{
			name: "Deleting key is not exists",
			fields: fields{
				Capacity: 1,
			},
			initValues: []initValue{
				{},
			},
			args: args{
				key: "key1",
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewLRUCache(tt.fields.Capacity)
			for _, initValue := range tt.initValues {
				c.Add(initValue.key, initValue.value)
			}

			if gotOk := c.Remove(tt.args.key); gotOk != tt.wantOk {
				t.Errorf("LRU.Remove() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
