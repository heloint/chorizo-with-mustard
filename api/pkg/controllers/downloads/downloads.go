package downloads

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var fileStorageUser string
var fileStorageHost string
var fileStoragePort string
var fileStorageAddr string
var fileStoragePassword string
var sshConfig ssh.ClientConfig

func init() {

	fileStorageUser = os.Getenv("FILETESTORAGE_USER")
	fileStorageHost = os.Getenv("FILESTORAGE_HOST")
	fileStoragePort = os.Getenv("FILESTORAGE_PORT")
	fileStoragePassword = os.Getenv("FILETESTORAGE_PASS")
	fileStorageAddr = fmt.Sprintf("%s:%s", fileStorageHost, fileStoragePort)

	sshConfig = ssh.ClientConfig{
		User: fileStorageUser,
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Auth: []ssh.AuthMethod{
			// ssh.Password(fileStoragePassword), // Won't work always, if host hasn't got it enabled. The password auth.
			ssh.KeyboardInteractive(sshInteractive), // Work around by imitating the keyboard interaction.
		},
	}

}

// Work around for that case if password authentication is not enabled on the host.
func sshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))

	for n, _ := range questions {
		answers[n] = fileStoragePassword
	}

	return answers, nil
}

// Let's try to imitate Python's context manager?
// Should the callback function use something like args?
func withSFTPConnection(fn func(sftpClient *sftp.Client)) {
	var err error

	conn, err := ssh.Dial("tcp", fileStorageAddr, &sshConfig)

	if err != nil {
		log.Fatal("unable to connect: ", err)
	}

	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)

	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer sftpClient.Close()

	fn(sftpClient)
}

// Download a single file from the remote server
func DownloadFile(c *gin.Context) {
	fileToDownload := c.Param("file")
	withSFTPConnection(func(sftpClient *sftp.Client) {
		file, err := sftpClient.Open(fileToDownload)
		if err != nil {
			log.Fatal("unable to connect: ")
		}
		defer file.Close()

		c.Header("Content-Disposition", "attachment; filename="+file.Name())

		// Copying file content to client response
		if _, err := io.Copy(c.Writer, file); err != nil {
			log.Fatalf("Error copying file content to response: %v", err)
		}
	})
}

func ListAllFiles(c *gin.Context) {
	// TODO
	log.Println("Lala")
}
