package utils

import (
	"math/rand"
	"strconv"

	"go-backend-demo/internal/model"
)

func RandomUser(id int) model.User {
	names := []string{"张三", "李四", "王五", "赵六", "钱七", "孙八"}

	return model.User{
		Id:         id,
		Name:       names[rand.Intn(len(names))],
		Age:        rand.Intn(30) + 18,
		Email:      "user" + strconv.Itoa(id) + "@test.com",
		CreateTime: GetChinaTime(),
	}
}
