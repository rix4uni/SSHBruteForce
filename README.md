# SSHBruteForce
 
# Usage
ssh password wordlists generator
```
usage: ssh-password-generator.py [-h] [-l LENGTH] [-n NUMOFPASS] [-t THREADS] [-o OUTPUT]

Generate random passwords

options:
  -h, --help            show this help message and exit
  -l LENGTH, --length LENGTH
                        Number of passwords length to generate (default 8)
  -n NUMOFPASS, --numofpass NUMOFPASS
                        Number of passwords to generate (default 1000)
  -t THREADS, --threads THREADS
                        Number of threads to use (default 100)
  -o OUTPUT, --output OUTPUT
                        Output file for passwords (default ssh-password.txt)
examples:
  python3 ssh-password-generator.py -l 8 -n 1000 -t 100 -o ssh-password.txt
```

# Usage
ssh login bruteforce
```
options:
  -ip string
        IP and port for SSH login (default "127.0.0.1:22")
  -p string
        file containing passwords to try (default "ssh-passwords.txt")
  -t int
        number of threads to use (default 20)
  -u string
        username for SSH login (default "root")
examples:
  go run ssh-brute-force.go -u root -ip 127.0.0.1:22 -p ssh-password.txt -t 100
```
