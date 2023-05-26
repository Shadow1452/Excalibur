Install Go: Make sure you have Go installed on your system. You can download and install it from the official Go website (https://golang.org) based on your operating system.

Create a file with password list: Create a text file and add the passwords you want to try, with each password on a new line.

Download the code: Save the code provided in a file with a ".go" extension, such as "ssh-cracker.go".

Build the program: Open a terminal or command prompt and navigate to the directory where you saved the code. Run the following command to build the executable:

shell
Copy code
go build ssh-cracker.go
This command will compile the Go code and create an executable file in the same directory.

Run the program: Once the build is successful, you can run the program using the following command:

./ssh-cracker -host <target_host> -port <target_port> -username <target_username> -password-file <path/to/passwords.txt> -V
Replace <target_host>, <target_port>, <target_username>, and <path/to/passwords.txt> with the appropriate values. The -V flag enables the display of password attempts.

Example:

./ssh-cracker -host 192.168.0.100 -port 22 -username admin -password-file passwords.txt -V
The program will start attempting SSH logins using the provided passwords.
