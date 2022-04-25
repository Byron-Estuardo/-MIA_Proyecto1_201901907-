package AdminDisks

import (
	"MiaProyecto2/structs"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

//Funciones Extras/Necesarias

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
		fmt.Println(err)
	}

}

func ObtenerFecha() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d/%02d/%02d %02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute())
	var Envio string = fecha
	//fmt.Println(Envio)
	return Envio
}

func Pausa() {
	fmt.Println("sistema pausado presione 'Enter' para continuar")
	fmt.Scanln()
}

func leerBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

func EscribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}

func BuscarParticion_Primaria_Extendida(direccion string, nombre string) int {
	file, err := os.OpenFile(direccion, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return -1
	} else {
		file.Seek(0, 0)
		TempMBR := structs.MBR{}
		var sizetemp int = int(unsafe.Sizeof(TempMBR))
		data := leerBytes(file, sizetemp)
		buffer := bytes.NewBuffer(data)
		Errores := binary.Read(buffer, binary.BigEndian, &TempMBR)
		if Errores != nil {
			fmt.Println("F perro")
		}
		for i := 0; i < 4; i++ {
			nameMbr := string(TempMBR.Mbr_partition[i].Part_name[:])
			if string(TempMBR.Mbr_partition[i].Part_status[:]) != "1" {
				if strings.Compare(nameMbr, nombre) == 0 {
					return i
				}
			}
		}
		return -1
	}
}

func BuscarParticion_Logica(direccion string, nombre string) int {
	var extendida int = -1
	file, err := os.OpenFile(direccion, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return -1
	} else {
		//var size1 int = int(unsafe.Sizeof(structs.MBR{}))
		file.Seek(0, 0)
		TempMBR := structs.MBR{}
		var sizetemp int = int(unsafe.Sizeof(TempMBR))
		data := leerBytes(file, sizetemp)
		buffer := bytes.NewBuffer(data)
		Errores := binary.Read(buffer, binary.BigEndian, &TempMBR)
		if Errores != nil {
			fmt.Println("F perro")
			return -1
		}
		for i := 0; i < 4; i++ {
			if string(TempMBR.Mbr_partition[i].Part_type[:]) == "E" {
				extendida = i
				break
			}
		}
		if extendida != -1 {
			var stringStart string = string(TempMBR.Mbr_partition[extendida].Part_start[:])
			var stringSize string = string(TempMBR.Mbr_partition[extendida].Part_size[:])
			InicioParticion, _ := strconv.Atoi(stringStart)
			InicioSize, _ := strconv.Atoi(stringSize)
			file.Seek(int64(InicioParticion), 0)
			var sizeEBR int64 = int64(unsafe.Sizeof(structs.EBR{}))
			var InicioEnviado int64 = int64(InicioParticion)
			var sumas = InicioParticion + InicioSize
			for int(InicioEnviado) < sumas {
				var Prueba string = leerEstructura(file, InicioEnviado, nombre)
				if Prueba == "error" {
					fmt.Println("Error al encontrar y leer las estructuras")
					return -1
				} else if Prueba == "1" {
					return int(InicioEnviado - sizeEBR)
				} else if Prueba == "-1" {
					return -1
				} else {
					InicioEnviado += sizeEBR
				}
			}
		}
	}
	return -1
}

/*
func leerEstructuraEBR(file *os.File, posicionEs int64, nombre string) structs.EBR{
	file.Seek(posicionEs, 0)
	estructuraTemporal := structs.EBR{}
	var size int = int(unsafe.Sizeof(estructuraTemporal))
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &estructuraTemporal)

	if err != nil {
		fmt.Println("binary.Read failed", err)
		return "error"

	} else {
		nombres := string(estructuraTemporal.Part_name[:]) + "\n"
		if strings.Compare(nombres,nombre) == 0{
			return "1"
		}else if string(estructuraTemporal.Part_next[:]) == "-1"{
			return "-1"
		}
	}
	return "error"
}
*/
func leerEstructura(file *os.File, posicionEs int64, nombre string) string {
	file.Seek(posicionEs, 0)
	estructuraTemporal := structs.EBR{}
	var size int = int(unsafe.Sizeof(estructuraTemporal))
	data := leerBytes(file, size)
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &estructuraTemporal)

	if err != nil {
		fmt.Println("binary.Read failed", err)
		return "error"

	} else {
		nombres := string(estructuraTemporal.Part_name[:]) + "\n"
		if strings.Compare(nombres, nombre) == 0 {
			return "1"
		} else if string(estructuraTemporal.Part_next[:]) == "-1" {
			return "-1"
		}
	}
	return "error"
}

func ExisteParticion(direccion string, nombre string) bool {
	var extendida int = -1
	var retorno bool = false
	file, err := os.OpenFile(direccion, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return retorno
	} else {
		//var size1 int = int(unsafe.Sizeof(structs.MBR{}))
		file.Seek(0, 0)
		TempMBR := structs.MBR{}
		var sizetemp int = int(unsafe.Sizeof(TempMBR))
		data := leerBytes(file, sizetemp)
		buffer := bytes.NewBuffer(data)
		Errores := binary.Read(buffer, binary.BigEndian, &TempMBR)
		if Errores != nil {
			fmt.Println("F perro")
		}
		for i := 0; i < 4; i++ {
			nameMbr := string(TempMBR.Mbr_partition[i].Part_name[:])
			if strings.Compare(nameMbr, nombre) == 0 {
				retorno = true
				return retorno
			} else if string(TempMBR.Mbr_partition[i].Part_type[:]) == "E" {
				extendida = i
			}
		}
		if extendida != -1 {
			var stringStart string = string(TempMBR.Mbr_partition[extendida].Part_start[:])
			var stringSize string = string(TempMBR.Mbr_partition[extendida].Part_size[:])
			InicioParticion, _ := strconv.Atoi(stringStart)
			InicioSize, _ := strconv.Atoi(stringSize)
			file.Seek(int64(InicioParticion), 0)
			var sizeEBR int64 = int64(unsafe.Sizeof(structs.EBR{}))
			var InicioEnviado int64 = int64(InicioParticion)
			var sumas = InicioParticion + InicioSize
			for int(InicioEnviado) < sumas {
				var Prueba string = leerEstructura(file, InicioEnviado, nombre)
				if Prueba == "error" {
					fmt.Println("Error al encontrar y leer las estructuras")
					return false
				} else if Prueba == "1" {
					return true
				} else if Prueba == "-1" {
					return false
				} else {
					InicioEnviado += sizeEBR
				}
			}
		}
	}
	return false
}

func CrearParticionPrimaria(direccion string, nombre string, size int, fit string, unit string, archivo string) {
	var auxFit string = ""
	var auxUnit string = ""
	var auxPath string = ""
	var size_bytes int = 1024
	var buffer1 int8 = 1
	TemporalMBR := structs.MBR{}
	var Tama int = int(unsafe.Sizeof(TemporalMBR))
	if fit != "" {
		auxFit = fit
	} else {
		auxFit = "W"
	}
	if unit != "" {
		auxUnit = unit
		if auxUnit == "b" {
			size_bytes = size
		} else if auxUnit == "k" {
			size_bytes = size * 1024
		} else {
			size_bytes = size * 1024 * 1024
		}
	} else {
		size_bytes = size * 1024
	}
	MBR := structs.MBR{}
	file, err := os.OpenFile(auxPath, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("Error El disco no existe")
	} else {
		var FlagParticion bool = false
		var Numero_Particion int = 0
		file.Seek(0, 0)
		var size int = int(unsafe.Sizeof(MBR))
		data := leerBytes(file, size)
		buffer := bytes.NewBuffer(data)
		err := binary.Read(buffer, binary.BigEndian, &MBR)
		if err != nil {
			fmt.Println("Error")
		} else {
			for i := 0; i < 4; i++ {
				var sizes string = string(MBR.Mbr_partition[i].Part_size[:])
				convertido, _ := strconv.Atoi(sizes)
				if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" || convertido >= size_bytes {
					FlagParticion = true
					Numero_Particion = i
					break
				}
			}
			if FlagParticion {
				var EspacioUsado int = 0
				for i := 0; i < 4; i++ {
					if string(MBR.Mbr_partition[i].Part_status[:]) == "1" {
						var sizes string = string(MBR.Mbr_partition[i].Part_size[:])
						convertido, _ := strconv.Atoi(sizes)
						EspacioUsado += convertido
					}
				}
				var sizes string = string(MBR.Mbr_size[:])
				convertido, _ := strconv.Atoi(sizes)
				fmt.Println("Espacio Disponible: " + string(convertido-EspacioUsado) + "Bytes")
				fmt.Println("Espacio Necesario: " + string(size_bytes) + "Bytes")
				if (convertido - EspacioUsado) >= size_bytes {
					if !ExisteParticion(direccion, nombre) {
						if string(MBR.Mbr_disk_fit[:]) == "F" {
							copy(MBR.Mbr_partition[Numero_Particion].Part_type[:], "P")
							copy(MBR.Mbr_partition[Numero_Particion].Part_fit[:], auxFit)
							if Numero_Particion == 0 {
								copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(Tama))
							} else {
								var unmenos = Numero_Particion - 1
								PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
								PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
								total := PartStart + PartSize
								copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(total))
							}
							copy(MBR.Mbr_partition[Numero_Particion].Part_size[:], string(size_bytes))
							copy(MBR.Mbr_partition[Numero_Particion].Part_status[:], "0")
							copy(MBR.Mbr_partition[Numero_Particion].Part_name[:], string(nombre))
							file.Seek(0, 0)
							var bufferControl bytes.Buffer
							binary.Write(&bufferControl, binary.BigEndian, &MBR)
							file.Write(bufferControl.Bytes())
							s := &buffer1
							var binario bytes.Buffer
							binary.Write(&binario, binary.BigEndian, s)
							for i := 0; i < size_bytes; i++ {
								EscribirBytes(file, binario.Bytes())
							}

						} else if string(MBR.Mbr_disk_fit[:]) == "B" {

						} else if string(MBR.Mbr_disk_fit[:]) == "W" {

						}
					} else {
						fmt.Println("Error Ya existe una particion con ese nombre")
					}
				} else {
					fmt.Println("Error la particion a crear sobrepasa el espacio libre disponible")
				}
			} else {
				fmt.Println("Error Ya existen 4 particiones, no se puede  crear otra")
			}
		}
	}
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
		}
	}
	if !BanderaRepetidos {
		if BanderaPath {
			if BanderaSize {
				MBR := structs.MBR{}
				total_size := 1024
				CrearDisco(ValPath)
				fechas := ObtenerFecha()
				//fmt.Println("Fecha")
				copy(MBR.Mbr_date_created[:], string(fechas))
				//fmt.Println(string(fechas))
				//fmt.Println(string(MBR.Mbr_date_created[:]))
				rand.Seed(int64(time.Now().UnixNano()))
				nuevoNumero := rand.Int()
				nuevs := strconv.Itoa(nuevoNumero)
				//fmt.Println(nuevs)
				copy(MBR.Mbr_disk_signature[:], nuevs)
				if BanderaUnit {
					if ValUnit == "m" {
						total_size = ValSize * 1024 * 1024
						ValSizeC := strconv.Itoa(ValSize * (1024 * 1024))
						copy(MBR.Mbr_size[:], ValSizeC)
					} else {
						total_size = ValSize * 1024
						ValSizeC := strconv.Itoa(ValSize * 1024)
						copy(MBR.Mbr_size[:], ValSizeC)
					}
				} else {
					total_size = ValSize * 1024 * 1024
					ValSizeC := strconv.Itoa(ValSize * (1024 * 1024))
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
					copy(MBR.Mbr_partition[i].Part_name[:], "0")
				}

				file, err := os.OpenFile(ValPath, os.O_RDWR, 0777)
				if err != nil {
					fmt.Println("Error al crear el archivo")
					fmt.Println(err)
				}
				defer file.Close()
				//fmt.Println("total_size")
				//fmt.Println(total_size)
				var temporal int8 = 0
				s := &temporal
				var binario bytes.Buffer
				binary.Write(&binario, binary.BigEndian, s)
				for i := 0; i < total_size; i++ {
					EscribirBytes(file, binario.Bytes())
				}

				file.Seek(0, 0)
				var bufferControl bytes.Buffer
				binary.Write(&bufferControl, binary.BigEndian, &MBR)
				_, errs := file.Write(bufferControl.Bytes())
				if errs != nil {
					fmt.Println("ERROR WE")
				}

			}
		}
	}
}

func Rmdisk(path string) {
	fmt.Println("Esta seguro que desea eliminar el disco? S/N :")
	reader := bufio.NewReader(os.Stdin)
	comando, _ := reader.ReadString('\n')
	if strings.Contains(comando, "S\n") || strings.Contains(comando, "s\n") {
		_ = os.Remove(path)
		fmt.Println("Disco eliminado con exito!")
	} else if strings.Contains(comando, "N\n") || strings.Contains(comando, "n\n") {
		fmt.Println("Cancelado con exito!")
	} else {
		fmt.Println("Opcion incorrecta")
	}
}
