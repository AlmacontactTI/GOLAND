package main

import (
	"log"

	"github.com/EdwinPirajan/RenewalCore.git/services"
)

func main() {

	client, err := services.NewSFTPClient()
	if err != nil {
		log.Fatalf("Error al crear el cliente SFTP: %s", err)
	}

	defer client.Close()

	services.PrintFiles(client)
	// 	procesosACerrar := []string{"AC Administrador de Clientes.exe"}

	// 	// Cierra los procesos especificados
	// 	for _, proceso := range procesosACerrar {
	// 		cmd := exec.Command("taskkill", "/IM", proceso, "/F")
	// 		err := cmd.Run()
	// 		if err != nil {
	// 			fmt.Println("Error cerrando el proceso:", err)
	// 		}
	// 	}

	// 	// Copia los archivos desde la ubicación remota a la local
	// 	origen := `\\10.98.16.27\publica_claro$\Tecnologia\ac\AC_ACTUALIZADORES 22.01.2019\AC_Personas`
	// 	destino := `C:\progra~1\acadmi~1`
	// 	err := CopyDir(origen, destino)
	// 	if err != nil {
	// 		fmt.Println("Error copiando los archivos:", err)
	// 	}

	// }

	// // CopyDir copia un directorio completo de origen a destino
	// func CopyDir(src string, dest string) error {
	// 	var err error
	// 	var fds []os.FileInfo
	// 	var srcinfo os.FileInfo

	// 	if srcinfo, err = os.Stat(src); err != nil {
	// 		return eservices
	// 	}

	// 	if err = os.MkdirAll(dest, srcinfo.Mode()); err != nil {
	// 		return err
	// 	}

	// 	if fds, err = ioutil.ReadDir(src); err != nil {
	// 		return err
	// 	}
	// 	for _, fd := range fds {
	// 		srcfp := filepath.Join(src, fd.Name())
	// 		destfp := filepath.Join(dest, fd.Name())

	// 		if fd.IsDir() {
	// 			if err = CopyDir(srcfp, destfp); err != nil {
	// 				fmt.Println(err)
	// 			}
	// 		} else {
	// 			if err = CopyFile(srcfp, destfp); err != nil {
	// 				fmt.Println(err)
	// 			}
	// 		}
	// 	}
	// 	return nil
	// }

	// func CopyFile(src, dest string) error {
	// 	var err error
	// 	var srcfd *os.File
	// 	var destfd *os.File
	// 	var srcinfo os.FileInfo

	// 	if srcfd, err = os.Open(src); err != nil {
	// 		return err
	// 	}
	// 	defer srcfd.Close()

	// 	if destfd, err = os.Create(dest); err != nil {
	// 		return err
	// 	}
	// 	defer destfd.Close()

	// 	if _, err = io.Copy(destfd, srcfd); err != nil {
	// 		return err
	// 	}
	// 	if srcinfo, err = os.Stat(src); err != nil {
	// 		return err
	// 	}
	// 	return os.Chmod(dest, srcinfo.Mode())
}
