package services

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
)

// ListFiles obtiene la lista de archivos en el directorio /hdvsftp/ftpti
func ListFiles(client *sftp.Client) ([]string, error) {
	files, err := client.ReadDir("/hdvsftp/ftpti/AC")
	if err != nil {
		return nil, err
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

// PrintFiles imprime la lista de archivos en el directorio /hdvsftp/ftpti
func PrintFiles(client *sftp.Client) {
	files, err := ListFiles(client)
	if err != nil {
		log.Fatalf("Error al listar archivos: %s", err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
