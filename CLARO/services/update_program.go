package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
)

func CopyFile(client *sftp.Client, remotePath, localPath string) error {
	remoteFile, err := client.Open(remotePath)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	localFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		return err
	}

	localFileInfo, err := localFile.Stat()
	if err != nil {
		return err
	}

	remoteFileInfo, err := remoteFile.Stat()
	if err != nil {
		return err
	}

	if localFileInfo.Size() != remoteFileInfo.Size() {
		return fmt.Errorf("el tamaño del archivo local no coincide con el tamaño del archivo remoto")
	}

	log.Printf("Archivo copiado de %s a %s", remotePath, localPath)
	return nil
}
func CopyAndMoveFile(client *sftp.Client, remoteFilePath, localDirectory string) error {

	localFilePath := filepath.Join(localDirectory)
	err := CopyFile(client, remoteFilePath, localFilePath)
	if err != nil {
		return fmt.Errorf("error al copiar el archivo: %s", err)
	}
	log.Println("Archivo copiado exitosamente.")

	newLocation := filepath.Join(localDirectory, "new") // Ejemplo de nuevo directorio, ajusta según necesites
	if err := os.Rename(localFilePath, newLocation); err != nil {
		return fmt.Errorf("error al mover el archivo: %s", err)
	}
	log.Println("Archivo movido exitosamente.")

	return nil
}
