package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func CreateNewDirs(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo mdkir /home/%s/documents /home/%s/pictures /home/%s/books", username, username, username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println("default directories were created for", username)
}