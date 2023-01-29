package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := os.ReadFile(file)

	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}


func main() {
	
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

    // ssh config
    hostKeyCallback, err := knownhosts.New(os.Getenv("KNOWN_HOSTS")) // "/../.ssh/known_hosts"
    
	if err != nil {
        log.Fatal(err)
    }
    
	config := &ssh.ClientConfig{
        User: "admin",
        Auth: []ssh.AuthMethod{
            PublicKeyFile(os.Getenv("PUBLIC_KEY")), // ssh.Password("password"),
        },
        HostKeyCallback: hostKeyCallback,
    }
    
	// connect ot ssh server
    conn, err := ssh.Dial("tcp", os.Getenv("IP_ADDRESS"), config)

    if err != nil {
        log.Fatal(err)
    }

	defer conn.Close()

	session, err := conn.NewSession()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// stdin pipe for commands
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// start remote shell
	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// var buff bytes.Buffer
	// session.Stdout = &buff
	// session.Stderr = &buff

	state := "new-user"
	username := "tommy"
	cmds := []string{}

	switch state {
		case "new-user":
			fmt.Println("create account for", username)
			cmds = append(cmds, fmt.Sprintf("sudo useradd -m %s", username)) // create user and home directory
			cmds = append(cmds, fmt.Sprintf("echo '%s:123456' | sudo chpasswd", username)) // update password
			cmds = append(cmds, fmt.Sprintf("sudo -u %s mkdir /home/%s/documents /home/%s/pictures /home/%s/books", username, username, username, username)) // make dirs
			cmds = append(cmds, fmt.Sprintf("sudo usermod -a -G nysif %s", username)) // add to company group

		case "update-password":
			fmt.Println("update password for", username)
			cmds = append(cmds, fmt.Sprintf("echo '%s:123456' | sudo chpasswd", username)) // update password

		case "lock-user": 
			fmt.Println("lockout", username)
			cmds = append(cmds, fmt.Sprintf("sudo passwd -l %s", username))
			cmds = append(cmds, fmt.Sprintf("sudo chage -E0 %s", username))	

		case "unlock-user":
			fmt.Println("unlock", username)
			cmds = append(cmds, fmt.Sprintf("sudo passwd -u %s", username))
			cmds = append(cmds, fmt.Sprintf("sudo chage -E -1 %s", username))	

		case "delete-user":
			fmt.Println("delete", username)
			cmds = append(cmds, fmt.Sprintf("sudo tar jcvf /home/admin/user-backups/%s.tar.bz2 /home/%s", username, username))
			cmds = append(cmds, fmt.Sprintf("sudo userdel -rf %s", username))

		default:
			fmt.Println("The state was incorrect...")
	}

	for _, cmd := range cmds {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cmd)
	}

}

