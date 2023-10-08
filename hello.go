package main

import (
	"example/user/hello/users"
	"fmt"
)

func main() {
	cratedUsers := users.CreateUser()
	fmt.Println(cratedUsers)
}
