package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func CreateUser(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo useradd -m %s", username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "was created succefully!")
}