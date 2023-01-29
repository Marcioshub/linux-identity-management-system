package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func AddToCompanyGroup(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo usermod -a -G nysif %s", username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "was added to the nysif group!")
}