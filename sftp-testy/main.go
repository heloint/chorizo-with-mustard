package main

import (
    "io"
	"fmt"
	"log"
	"net"
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getLala(c *gin.Context) {
    var file *sftp.File
    var err error

	port := 22
	hostname := "supersecret.ip.com"
	user := "user"
	addr := fmt.Sprintf("%s:%d", hostname, port)

	fmt.Println(addr)

	config := ssh.ClientConfig{
		User: user,
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Auth: []ssh.AuthMethod{
            // ssh.Password("password"), // Won't work always, if host hasn't got it enabled. The password auth.
			ssh.KeyboardInteractive(SshInteractive), // Work around by imitating the keyboard interaction.
		},
	}
	conn, err := ssh.Dial("tcp", addr, &config)

	if err != nil {
		log.Fatal("unable to connect: ", err)
	}

    fmt.Println("yippyy")
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)

	if err != nil {
		log.Fatal("unable to connect: ")
	}
    defer sftpClient.Close()

	file, err = sftpClient.Open("/path/to/target/file.lala")
	if err != nil {
		log.Fatal("unable to connect: ")
	}
    defer file.Close()

    c.Header("Content-Disposition", "attachment; filename="+file.Name())

    // Copying file content to client response
    if _, err := io.Copy(c.Writer, file); err != nil {
        log.Fatalf("Error copying file content to response: %v", err)
    }
}

func main() {

	// Start writing the SDTOUT & STDERR to log files.
	// apiLogs.InitAPILogs()


	r := gin.Default()

    // Apply defined headers.
	r.Use(CORSMiddleware())

    // Routes
	r.GET("/", getLala)

    // Dynamic routing example.
    /* r. GET("/user/:username", func (c *gin.Context) {
        username := c.Param("username")
        log.Println(username)
    }) */

	r.Run("0.0.0.0:8000")
}

// Set API headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header(
            "Access-Control-Allow-Headers", 
            `Content-Type,
             access-control-allow-origin,
             Content-Length,
             Accept-Encoding,
             X-CSRF-Token,
             Authorization,
             accept,
             origin,
             Cache-Control,
             X-Requested-With`,
         )
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

/*
Work around for that case if password authentication is not enabled on the host.
*/
func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	// The second parameter is unused
	for n, _ := range questions {
		answers[n] = "supersecretpassword"
	}

	return answers, nil
}
