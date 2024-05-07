package services

import (
	"log"
	"os"

	"github.com/EdwinPirajan/RenewalCore.git/models"
	"github.com/joho/godotenv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func LoadCredentials() *models.Credentials {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar las variables de entorno: %s", err)
	}

	return &models.Credentials{
		Server:   os.Getenv("SFTP_SERVER"),
		User:     os.Getenv("SFTP_USER"),
		Password: os.Getenv("SFTP_PASSWORD"),
	}
}

func NewSFTPClient() (*sftp.Client, error) {
	credentials := LoadCredentials()

	config := &ssh.ClientConfig{
		User: credentials.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(credentials.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", credentials.Server, config)
	if err != nil {
		return nil, err
	}

	log.Println("Conexión SFTP exitosa.")

	return sftp.NewClient(conn)
}
