package main

import (
	"MiaProyecto2/structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetDirectorio(direccion string) string {
	aux := direccion
	//aux := "/home/curious1924/Escritorio/disco23.dk"
	res := strings.Split(aux, "/")
	//fmt.Println(res)
	cantidad := strings.Count(aux, "/") // [home curious1924 Escritorio disco23.dk]
	//fmt.Println(cantidad)
	var envio string = ""
	for i := 0; i < cantidad; i++ { //ciclo for ignorando el ultimo item
		envio += res[i] + "/" //unimos la ruta ignorando el ultimo valor
	}
	//fmt.Println(envio)
	return envio
}

func Shellout(command string) error {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err
}

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

func CrearDisco(direccion string) {
	var aux string = GetDirectorio(direccion)
	err := os.Mkdir(aux, 0777)
	if err != nil {
		fmt.Println(err)
	}
	err1 := os.Chmod(aux, 0777)
	if err1 != nil {
		fmt.Println(err)
	}

	var arch string = direccion
	file, err := os.Create(arch)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	if err != nil {
		//fmt.Println("Error al crear el archivo")
		log.Fatal(err)
	}

}

func ObtenerFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Println("La fecha actual es =>", fecha)
	return fecha
}

func Pausa() {
	fmt.Println("sistema pausado presione 'Enter' para continuar")
	fmt.Scanln()
}

// mkdisk("5", "wf", "m", "/home/curious1924/Escritorio/Joder/disco3.dk")
func Mkdisk(size string, fit string, unit string, path string) {
	fmt.Println("Que pdo!!!")
	var BanderaSize bool = false
	var BanderaFit bool = false
	var BanderaUnit bool = false
	var BanderaPath bool = false
	var BanderaRepetidos bool = false
	var ValSize int = 0
	var ValFit string = ""
	var ValUnit string = ""
	var ValPath string = ""
	if size != "" {
		Prueba, err := strconv.Atoi(size)
		ValSize = Prueba
		fmt.Println(err)
		BanderaSize = true
		if ValSize < 0 {
			BanderaRepetidos = true
			BanderaSize = false
		}
	}
	if path != "" {
		ValPath = path
		BanderaPath = true
		if ValPath == "" {
			BanderaRepetidos = true
			BanderaPath = false
		}
	}
	if fit != "" {
		BanderaFit = true
		if fit == "b" {
			ValFit = "B"
		} else if fit == "f" {
			ValFit = "F"
		} else if fit == "w" {
			ValFit = "W"
		}
	}
	if unit != "" {
		BanderaUnit = true
		if unit == "k" || unit == "K" {
			ValUnit = "k"
		} else if unit == "m" || unit == "M" {
			ValUnit = "m"
		} else if fit == "w" {
			ValFit = "W"
		}
	}
	if !BanderaRepetidos {
		if BanderaPath {
			if BanderaSize {
				MBR := structs.MBR{}
				total_size := 1024
				CrearDisco(ValPath)
				fechas := ObtenerFecha()
				copy(MBR.Mbr_date_created[:], fechas)
				rand.Seed(int64(time.Now().UnixNano()))
				nuevoNumero := rand.Int()
				nuevs := strconv.Itoa(nuevoNumero)
				copy(MBR.Mbr_date_created[:], nuevs)
				if BanderaUnit {
					if ValUnit == "m" {
						total_size = ValSize * 1024
						ValSizeC := strconv.Itoa(ValSize * (1024 * 1024))
						copy(MBR.Mbr_size[:], ValSizeC)
					} else {
						total_size = ValSize
						ValSizeC := strconv.Itoa(ValSize * 1024)
						copy(MBR.Mbr_size[:], ValSizeC)
					}
				} else {
					total_size = ValSize * 1024
					ValSizeC := strconv.Itoa(ValSize * 1048576)
					copy(MBR.Mbr_size[:], ValSizeC)
				}

				if BanderaFit {
					copy(MBR.Mbr_disk_fit[:], ValFit)
				} else {
					copy(MBR.Mbr_disk_fit[:], "F")
				}
				for i := 0; i < 4; i++ {
					copy(MBR.Mbr_partition[i].Part_status[:], "0")
					copy(MBR.Mbr_partition[i].Part_type[:], "0")
					copy(MBR.Mbr_partition[i].Part_fit[:], "0")
					copy(MBR.Mbr_partition[i].Part_size[:], "0")
					copy(MBR.Mbr_partition[i].Part_start[:], "-1")
					copy(MBR.Mbr_partition[i].Part_name[:], "a")
				}

				var comando string = "dd if=/dev/zero of=" + ValPath + " bs=1024 count=" + strconv.Itoa(total_size)
				err := Shellout(comando)
				if err != nil {
					log.Printf("error: %v\n", err)
				}
				file, err := os.OpenFile(ValPath, os.O_RDWR, 0777)
				if err != nil {
					fmt.Println("Error al crear el archivo")
					log.Fatal(err)
				}
				defer file.Close()
				var binario bytes.Buffer
				for i := 0; i < total_size; i++ {
					escribirBytes(file, binario.Bytes())
				}

				file.Seek(0, 0)
				var bufferControl bytes.Buffer
				binary.Write(&bufferControl, binary.BigEndian, &MBR)
				_, errs := file.Write(bufferControl.Bytes())
				if errs != nil {
					fmt.Println("ERROR WE")
					log.Fatal(err)
				}
				//MBR.Mb

			}
		}
	}
}

/*
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Partition struct {
	Part_status []byte
	Part_type   []byte
	Part_fit    []byte
	Part_start  []byte
	Part_size   []byte
	Part_name   []byte
}

type MBR struct {
	Mbr_size           []byte
	Mbr_date_created   []byte
	Mbr_disk_signature []byte
	Mbr_disk_fit       []byte
	Mbr_partition      [4]Partition
}

type EBR struct {
	Part_status []byte
	Part_fit    []byte
	Part_start  []byte
	Part_size   []byte
	Part_next   []byte
	Part_name   []byte
}

type SuperBloque struct {
	S_filesystem_type   []byte
	S_inodes_count      []byte
	S_blocks_count      []byte
	S_free_blocks_count []byte
	S_free_inodes_count []byte
	S_mtime             []byte
	S_umtime            []byte
	S_mnt_count         []byte
	S_magic             []byte
	S_inode_size        []byte
	S_block_size        []byte
	S_first_ino         []byte
	S_first_blo         []byte
	S_bm_inode_start    []byte
	S_bm_block_start    []byte
	S_inode_start       []byte
	S_block_start       []byte
}

type InodoTable struct {
	I_uid   []byte
	I_gid   []byte
	I_size  []byte
	I_block []byte
	I_type  []byte
	I_perm  []byte
	I_atime []byte
	I_ctime []byte
	I_mtime []byte
}

type Content struct {
	B_name  []byte
	B_inodo []byte
}

type BloqueCarpeta struct {
	B_content [4]Content
}

type BloqueArchivo struct {
	B_content []byte
}

type BloqueApuntadores struct {
	B_pointer []byte
}

type Sesion struct {
	Id_user       []byte
	Id_grp        []byte
	InicioSuper   []byte
	InicioJournal []byte
	Tipo_sistema  []byte
	Direccion     []byte
	Fit           []byte
}

type Usuario struct {
	Id_usr   []byte
	Id_grp   []byte
	Username []byte
	Password []byte
	Group    []byte
}

func getDirectorio(direccion string) string {
	aux := direccion
	//aux := "/home/curious1924/Escritorio/disco23.dk"
	res := strings.Split(aux, "/")
	//fmt.Println(res)
	cantidad := strings.Count(aux, "/") // [home curious1924 Escritorio disco23.dk]
	//fmt.Println(cantidad)
	var envio string = ""
	for i := 0; i < cantidad; i++ { //ciclo for ignorando el ultimo item
		envio += res[i] + "/" //unimos la ruta ignorando el ultimo valor
	}
	//fmt.Println(envio)
	return envio
}

func Shellout(command string) error {
	const ShellToUse = "bash"
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err
}

func crearDisco(direccion string) {
	var aux string = getDirectorio(direccion)
	var comando string = "sudo mkdir -p '" + aux + "'"
	err := Shellout(comando)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	var comando2 string = "sudo chmod -R 777 '" + aux + "'"
	err1 := Shellout(comando2)
	if err1 != nil {
		log.Printf("error: %v\n", err)
	}
	var arch string = direccion

	file, err := os.OpenFile(arch, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error al crear el archivo")
		log.Fatal(err)
	}
	defer file.Close()
}

func ObtenerFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Println("La fecha actual es =>", fecha)
	return fecha
}

func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}

// mkdisk("5", "wf", "m", "/home/curious1924/disco3.dk")
func mkdisk(size string, fit string, unit string, path string) {
	fmt.Println("Que pdo!!!")
	var BanderaSize bool = false
	var BanderaFit bool = false
	var BanderaUnit bool = false
	var BanderaPath bool = false
	var BanderaRepetidos bool = false
	var ValSize int = 0
	var ValFit string = ""
	var ValUnit string = ""
	var ValPath string = ""
	fmt.Println(BanderaFit)
	fmt.Println(BanderaUnit)
	fmt.Println(ValFit)
	fmt.Println(ValUnit)
	if size != "" {
		Prueba, err := strconv.Atoi(size)
		ValSize = Prueba
		fmt.Println(err)
		BanderaSize = true
		if ValSize < 0 {
			BanderaRepetidos = true
			BanderaSize = false
		}
	}
	if path != "" {
		ValPath = size
		BanderaPath = true
		if ValPath == "" {
			BanderaRepetidos = true
			BanderaPath = false
		}
	}
	if fit != "" {
		BanderaFit = true
		if fit == "b" {
			ValFit = "B"
		} else if fit == "f" {
			ValFit = "F"
		} else if fit == "w" {
			ValFit = "W"
		}
	}
	if unit != "" {
		BanderaUnit = true
		if unit == "k" || unit == "K" {
			ValUnit = "k"
		} else if unit == "m" || unit == "M" {
			ValUnit = "m"
		} else if fit == "w" {
			ValFit = "W"
		}
	}
	if !BanderaRepetidos {
		if BanderaPath {
			if BanderaSize {
				MBR1 := MBR{}
				total_size := 1024
				fmt.Println(total_size)
				crearDisco(ValPath)
				fechas := ObtenerFecha()
				MBR1.Mbr_date_created = []byte(fechas)
				//fmt.Println(MBR.Mbr_date_created)
				rand.Seed(int64(time.Now().UnixNano()))
				nuevoNumero := rand.Int()
				nuevs := strconv.Itoa(nuevoNumero)
				MBR1.Mbr_date_created = []byte(nuevs)
				fmt.Println(nuevoNumero)
				if BanderaUnit {
					if ValUnit == "m" {
						total_size = ValSize * 1024
						ValSizeC := strconv.Itoa(ValSize * (1024 * 1024))
						MBR1.Mbr_size = []byte(ValSizeC)
					} else {
						total_size = ValSize
						ValSizeC := strconv.Itoa(ValSize * 1024)
						MBR1.Mbr_size = []byte(ValSizeC)
					}
				} else {
					total_size = ValSize * 1024
					ValSizeC := strconv.Itoa(ValSize * 1048576)
					MBR1.Mbr_size = []byte(ValSizeC)
				}

				if BanderaFit {
					MBR1.Mbr_disk_fit = []byte(ValFit)
				} else {
					MBR1.Mbr_disk_fit = []byte("F")
				}
				for i := 0; i < 4; i++ {
					MBR1.Mbr_partition[i].Part_status = []byte("0")
					MBR1.Mbr_partition[i].Part_type = []byte("0")
					MBR1.Mbr_partition[i].Part_fit = []byte("0")
					MBR1.Mbr_partition[i].Part_size = []byte("0")
					MBR1.Mbr_partition[i].Part_start = []byte("-1")
					MBR1.Mbr_partition[i].Part_name = []byte("")
				}

				var comando string = "dd if=/dev/zero of=" + ValPath + " bs=1024 count=" + strconv.Itoa(total_size)
				err := Shellout(comando)
				if err != nil {
					log.Printf("error: %v\n", err)
				}
				file, err := os.OpenFile(ValPath, os.O_RDWR, 0777)
				if err != nil {
					fmt.Println("Error al crear el archivo")
					log.Fatal(err)
				}
				defer file.Close()

				file.Seek(0, 0)
				var bufferControl bytes.Buffer
				binary.Write(&bufferControl, binary.BigEndian, &MBR1)
				_, errs := file.Write(bufferControl.Bytes())
				if errs != nil {
					log.Fatal(err)
				}
				//MBR.Mb

			}
		}
	}
}
*/
