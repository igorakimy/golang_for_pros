package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// Получить идентификатор текущего пользователя.
	fmt.Println("User id:", os.Getuid())

	var u *user.User
	// Получить текущего пользователя.
	u, _ = user.Current()
	fmt.Print("Group ids: ")
	// Получить список идентификаторов групп, к которым принадлежит текущий пользователь.
	groupIDs, _ := u.GroupIds()
	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
