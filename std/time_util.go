package std

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"          // 代表年月日
	TIME = "2006-01-02 15:04:05" // 代表年月日时分秒
)

// 时间常见操作
func PackageTime() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now.Unix())

	fmt.Println()
	//休眠
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 2)
	}
	fmt.Println()
	// 两个时刻之差
	t1 := time.Now()
	diff := t1.Sub(now)
	fmt.Println(diff.Milliseconds())            // 获取时间差毫秒数
	fmt.Println(time.Since(now).Milliseconds()) // 获取时间差

	// 计算时间的加法
	// time + duration = time
	d := 2 * time.Second
	t2 := t1.Add(d)
	fmt.Println(t2.Unix())

	// 时间转换
	duration, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println(duration)

	// 时间的格式化

	fmt.Println(now.Format(DATE)) //2023-06-20
	s := now.Format(TIME)
	fmt.Println(s) //2023-06-20 17:14:34

	// 时间的解析
	// 构造时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t3, _ := time.ParseInLocation(TIME, s, loc)
	fmt.Println(t3.Unix())

}
