package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/crypto/ssh"
)

const COMPOSE_FILE_PATH = "./docker-compose.yml"

func main() {
	var (
		server   = flag.String("server", "localhost", "Deploy to localhost or remote server")
		username = flag.String("username", "", "Remote server ssh username")
		password = flag.String("password", "", "Remote server ssh password")
	)
	flag.Parse()

	fmt.Printf("Deploying to %s ...\n", *server)
	if *server == "localhost" {
		deployToLocalhost()
	} else {
		deployToRemoteServer(*server, *username, *password)
	}
}

func deployToLocalhost() {
	cmd := exec.Command("docker", "compose", "-f", COMPOSE_FILE_PATH, "up", "-d")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\nOutput: %s\n", err, string(output))
	}
	fmt.Println(string(output))
}

func deployToRemoteServer(server string, username string, password string) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", server+":22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	cmd := exec.Command("scp", COMPOSE_FILE_PATH, username+"@"+server+":~")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\nOutput: %s\n", err, string(output))
	}
	fmt.Println(string(output))

	cmd = exec.Command("ssh", username+"@"+server, "docker compose", "-f", "~/docker-compose.yml", "up", "-d")
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\nOutput: %s\n", err, string(output))
	}
	fmt.Println(string(output))
}
