package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gonote/db"
	"net/http"
	"net/url"
	"strconv"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
}

func (student *Student) toString() string {
	return "Name:" + student.Name + " Age:" + strconv.Itoa(student.Age) + " Height:" + strconv.FormatFloat(float64(student.Height), 'f', 2, 32)
}

// GetStudentInfo 通过学生id获取学生信息
func GetStudentInfo(studentId string) Student {
	studentInfo, err := db.RedisClient.HGetAll("student:" + studentId).Result()
	db.HandleRedisError(err)
	fmt.Println(studentInfo)
	stud := Student{}
	for field, value := range studentInfo {
		if field == "Name" {
			stud.Name = value
		} else if field == "Age" {
			stud.Age, _ = strconv.Atoi(value)
		} else if field == "Height" {
			height, _ := strconv.ParseFloat(value, 10)
			stud.Height = float32(height)
		}
	}
	return stud
}

func GetName(ctx *gin.Context) {
	param := ctx.Query("student_id")
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "student_id is empty")
		return
	}
	studentId, err := url.QueryUnescape(param)
	if err != nil {
		ctx.String(http.StatusBadRequest, "student_id is error")
		return
	}

	// 通过学生id获取学生信息
	std := GetStudentInfo(studentId)
	ctx.JSON(http.StatusOK, std)
}
