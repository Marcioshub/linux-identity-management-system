package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func UpdatePasswd(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("echo '%s:123456' | sudo chpasswd", username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "password was updated!")
}