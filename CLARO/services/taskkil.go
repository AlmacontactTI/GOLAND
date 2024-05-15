package services

import (
	"log"
	"os/exec"
)

func KillProcess(processName string) {
	cmd := exec.Command("taskkill", "/IM", processName, "/F", "/T")
	err := cmd.Run()
	if err != nil {
		log.Printf("Error cerrando el proceso %s: %v", processName, err)
	} else {
		log.Printf("Proceso %s cerrado exitosamente", processName)
	}
}
