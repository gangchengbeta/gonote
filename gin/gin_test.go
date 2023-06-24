package gin

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetStudentInfo(t *testing.T) {
	student := GetStudentInfo("123")
	fmt.Println(student.toString())
	if student.Name != "coder4j" || student.Age != 18 || student.Height != 180.04 {
		t.Error("获取学生信息失败")
	}
}

func TestGetName(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/get_name?student_id=123")
	if err != nil {
		fmt.Println(err)
		t.Error("请求失败")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error("响应流关闭失败")
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("响应流读取失败")
	}
	fmt.Println(string(bytes))
}
