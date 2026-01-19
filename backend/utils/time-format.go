package utils

import "time"

func TimeFormat() string {
	return "2006-01-02 15:04:05"
}

func GetChinaTime() string {
	// 8*3600 表示 8小时的秒数
	cstZone := time.FixedZone("CST", 8*3600)
	// 2. 获取当前时间，并切换到该时区
	cstTime := time.Now().In(cstZone)

	return cstTime.Format(TimeFormat())
}
