package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func BackupUserData(session *ssh.Session, username string) {
	cmd := fmt.Sprintf("sudo tar jcvf /home/pi/user-backups/%s.tar.bz2 /home/%s", username, username)

	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username, "was saved succefully!")
}