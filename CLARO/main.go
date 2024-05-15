package main

import (
	"log"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/AlmacontactTI/GOLAND.git/services"
	"golang.org/x/crypto/ssh"
)

func main() {
	a := app.New()
	w := a.NewWindow("RenewalCore")

	// Credenciales SFTP
	sftpServer := "10.96.16.28:22"
	sftpUser := "ftpti"
	sftpPassword := "adc#2024"
	config := &ssh.ClientConfig{
		User: sftpUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sftpPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := services.NewSFTPClient(sftpServer, config)
	if err != nil {
		log.Fatalf("Error al crear el cliente SFTP: %s", err)
		return
	}
	defer client.Close()

	remoteFilePath := "/hdvsftp/ftpti/AC/AC Administrador de Clientes.exe"
	// localDirectory := "C:\\progra~1\\acadmi~1\\"
	localDirectory := "C:\\Program Files\\AC Administración de Clientes\\"
	// localDirectory := "C:\\Users\\jpulido\\Documents\\Goland_copy"
	localFileName := "AC Administrador de Clientes.exe"

	services.KillProcess("AC Administrador de Clientes.exe")

	button := widget.NewButton("Actualizar archivo", func() {
		localPath := filepath.Join(localDirectory, localFileName)
		if err := services.CopyFile(client, remoteFilePath, localPath); err != nil {
			log.Fatalf("Error al copiar el archivo: %s", err)
			return
		}

		// Reemplazar el contenido de la ventana principal
		confirmLabel := widget.NewLabel("La aplicación se ha actualizado correctamente.")
		confirmButton := widget.NewButton("Aceptar", func() {
			a.Quit()
		})
		confirmContainer := container.NewVBox(
			confirmLabel,
			confirmButton,
		)
		w.SetContent(confirmContainer)
		w.Resize(fyne.NewSize(300, 90))
	})

	container := container.NewVBox(button)
	w.SetContent(container)
	w.Resize(fyne.NewSize(300, 50))
	w.ShowAndRun()
}
