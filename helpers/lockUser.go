package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func LockUser(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo passwd --lock %s", username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "account has been locked")
}