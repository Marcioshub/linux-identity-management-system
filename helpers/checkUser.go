package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func CheckUser(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("grep %s /etc/passwd | awk -F ':' '{print $1'\r'}'", username)

	if cmd == "" {
		log.Fatal("User does not exists")
	}

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "exists in system")
}
