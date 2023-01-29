package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func DeleteUser(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo userdel -r %s", username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "was deleted!")
}