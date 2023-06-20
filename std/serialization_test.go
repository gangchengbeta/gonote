package std

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

// 单元测试
var (
	s = Student{"张宝磊", 58, false}
	c = Class{
		Id:       "1(1)班",
		Students: []Student{s, s, s},
	}
)

// TestSerialJSON 单元测试
func TestSerialJSON(t *testing.T) {
	marshal, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class
	err = json.Unmarshal(marshal, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}

	fmt.Println("json 没毛病")
}

// TestSerialSonic 单元测试
func TestSerialSonic(t *testing.T) {
	marshal, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c2 Class
	err = sonic.Unmarshal(marshal, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}

	fmt.Println("sonic 没毛病")
}

// TestSerialJSON 基准测试
func BenchmarkSerialJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshal, _ := json.Marshal(c)
		var c2 Class
		_ = json.Unmarshal(marshal, &c2)
	}
}

// BenchmarkSerialSonic 基准测试
func BenchmarkSerialSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshal, _ := sonic.Marshal(c)
		var c2 Class
		_ = sonic.Unmarshal(marshal, &c2)
	}
}
