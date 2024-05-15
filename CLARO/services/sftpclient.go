package services

import (
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Realizando la conexión SFTP
func NewSFTPClient(server string, config *ssh.ClientConfig) (*sftp.Client, error) {
	conn, err := ssh.Dial("tcp", server, config)
	if err != nil {
		return nil, err
	}

	log.Println("Conexión SFTP exitosa.")
	return sftp.NewClient(conn)
}
