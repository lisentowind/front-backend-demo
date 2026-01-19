package utils

import (
	"math/rand"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func RandomUser(id int) User {
	names := []string{"张三", "李四", "王五", "赵六", "钱七", "孙八"}

	return User{
		Id:    id,
		Name:  names[rand.Intn(len(names))],
		Age:   rand.Intn(30) + 18,
		Email: "user" + strconv.Itoa(id) + "@test.com",
	}
}
