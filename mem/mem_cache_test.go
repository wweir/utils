package mem_test

import (
	"errors"
	"log"
	"testing"

	"github.com/wweir/utils/mem"
)

type testType struct {
	Test string
}

func (t *testType) GetOne(key interface{}) (interface{}, error) {
	var str string
	var ok bool
	if str, ok = key.(string); !ok {
		return nil, errors.New("type not match")
	}
	log.Println("from db")
	return &testType{str}, nil
}

func TestRemember(t *testing.T) {
	dst := &testType{}
	type args struct {
		dst mem.CacheType
		key interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		"from db",
		args{dst, "mock_key"},
		false,
	}, {
		"from cache",
		args{dst, "mock_key"},
		false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mem.Remember(tt.args.dst, tt.args.key); (err != nil) != tt.wantErr || dst.Test != "mock_key" {
				t.Errorf("Remember() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	dst := &testType{}
	if err := mem.Remember(dst, "mock_key"); err != nil {
		t.Errorf(err.Error())
	}

	mem.Delete(dst, "mock_key")

	if err := mem.Remember(dst, "mock_key"); err == nil {
		t.Errorf(err.Error())
	}
}
