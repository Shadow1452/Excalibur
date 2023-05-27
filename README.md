Clone the directory:

git clone https://github.com/Shadow1452/Excalibur/

Cd into directory

Run the installer.sh:

sudo bash installer.sh

Run the program: Once the install is successful, you can run the program using the following command:

./ssh-cracker -host <target_host> -port <target_port> -username <target_username> -password-file <path/to/passwords.txt> -V
Replace <target_host>, <target_port>, <target_username>, and <path/to/passwords.txt> with the appropriate values. The -V flag enables the display of password attempts.

Example:

./ssh-cracker -host 192.168.0.100 -port 22 -username admin -password-file passwords.txt -V

The program will start attempting SSH logins using the provided passwords.

Thank you for yousing Excalibur.

