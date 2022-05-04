package AdminDisks

import (
	"MiaProyecto2/structs"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type NodoMount struct {
	direccion string
	nombre    string
	id        string
	letra     string
	num       int
}

var ListaMount = []NodoMount{}
var BanderaLogin bool = false
var CurrentSession = structs.Sesion{}

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

	file, err := os.Create(direccion)
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
			currentPosition, _ := file.Seek(0, 1)
			for int(currentPosition) < sumas {
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

func CrearParticionPrimaria(direccion string, nombre string, size int, fit string, unit string) {
	var auxFit string = ""
	var auxUnit string = ""
	var auxPath string = direccion
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
				fmt.Println("Espacio Disponible: " + string(rune(convertido-EspacioUsado)) + "Bytes")
				fmt.Println("Espacio Necesario: " + string(rune(size_bytes)) + "Bytes")
				if (convertido - EspacioUsado) >= size_bytes {
					if !ExisteParticion(direccion, nombre) {
						if string(MBR.Mbr_disk_fit[:]) == "F" {
							copy(MBR.Mbr_partition[Numero_Particion].Part_type[:], "P")
							copy(MBR.Mbr_partition[Numero_Particion].Part_fit[:], auxFit)
							if Numero_Particion == 0 {
								copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(rune(Tama)))
							} else {
								var unmenos = Numero_Particion - 1
								PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
								PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
								total := PartStart + PartSize
								copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(rune(total)))
							}
							copy(MBR.Mbr_partition[Numero_Particion].Part_size[:], string(rune(size_bytes)))
							copy(MBR.Mbr_partition[Numero_Particion].Part_status[:], "0")
							copy(MBR.Mbr_partition[Numero_Particion].Part_name[:], string(nombre))
							file.Seek(0, 0)
							var bufferControl bytes.Buffer
							binary.Write(&bufferControl, binary.BigEndian, &MBR)
							file.Write(bufferControl.Bytes())
							s := &buffer1
							var binario bytes.Buffer
							binary.Write(&binario, binary.BigEndian, s)
							partstart, _ := strconv.Atoi(string(MBR.Mbr_partition[Numero_Particion].Part_start[:]))
							file.Seek(int64(partstart), 0)
							for i := 0; i < size_bytes; i++ {
								_, err := file.Write(binario.Bytes())

								if err != nil {
									fmt.Println(err)
								}
							}
							fmt.Println("-------------------- FDISK --------------------")
							fmt.Println("------------------ Disco Primario Creado ------------------")
							fmt.Println("Part Status: " + string(MBR.Mbr_partition[Numero_Particion].Part_status[:]))
							fmt.Println("Part Fit: " + string(MBR.Mbr_partition[Numero_Particion].Part_fit[:]))
							fmt.Println("Part Start: " + string(MBR.Mbr_partition[Numero_Particion].Part_start[:]))
							fmt.Println("Part Size: " + string(MBR.Mbr_partition[Numero_Particion].Part_size[:]))
							fmt.Println("Part name: " + string(MBR.Mbr_partition[Numero_Particion].Part_name[:]))
						} else if string(MBR.Mbr_disk_fit[:]) == "B" {
							var BestIndex int = Numero_Particion
							for i := 0; i < 4; i++ {
								convert1, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
								if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" && convert1 >= size_bytes {
									if i != Numero_Particion {
										convert2, _ := strconv.Atoi(string(MBR.Mbr_partition[BestIndex].Part_size[:]))
										convert3, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
										if convert2 > convert3 {
											BestIndex = i
											break
										}
									}
								}
							}
							copy(MBR.Mbr_partition[BestIndex].Part_type[:], "P")
							copy(MBR.Mbr_partition[BestIndex].Part_fit[:], auxFit)
							if BestIndex == 0 {
								copy(MBR.Mbr_partition[BestIndex].Part_start[:], string(rune(Tama)))
							} else {
								var unmenos = BestIndex - 1
								PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
								PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
								total := PartStart + PartSize
								copy(MBR.Mbr_partition[BestIndex].Part_start[:], string(rune(total)))
							}
							copy(MBR.Mbr_partition[BestIndex].Part_size[:], string(rune(size_bytes)))
							copy(MBR.Mbr_partition[BestIndex].Part_status[:], "0")
							copy(MBR.Mbr_partition[BestIndex].Part_name[:], string(nombre))
							file.Seek(0, 0)
							var bufferControl bytes.Buffer
							binary.Write(&bufferControl, binary.BigEndian, &MBR)
							file.Write(bufferControl.Bytes())
							s := &buffer1
							var binario bytes.Buffer
							binary.Write(&binario, binary.BigEndian, s)
							partstart, _ := strconv.Atoi(string(MBR.Mbr_partition[BestIndex].Part_start[:]))
							file.Seek(int64(partstart), 0)
							for i := 0; i < size_bytes; i++ {
								_, err := file.Write(binario.Bytes())
								if err != nil {
									fmt.Println(err)
								}
							}
							fmt.Println("-------------------- FDISK --------------------")
							fmt.Println("------------------ Disco Primario Creado ------------------")
							fmt.Println("Part Status: " + string(MBR.Mbr_partition[BestIndex].Part_status[:]))
							fmt.Println("Part Fit: " + string(MBR.Mbr_partition[BestIndex].Part_fit[:]))
							fmt.Println("Part Start: " + string(MBR.Mbr_partition[BestIndex].Part_start[:]))
							fmt.Println("Part Size: " + string(MBR.Mbr_partition[BestIndex].Part_size[:]))
							fmt.Println("Part name: " + string(MBR.Mbr_partition[BestIndex].Part_name[:]))
						} else if string(MBR.Mbr_disk_fit[:]) == "W" {
							var WorstIndex int = Numero_Particion
							for i := 0; i < 4; i++ {
								convert1, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
								if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" && convert1 >= size_bytes {
									if i != Numero_Particion {
										convert2, _ := strconv.Atoi(string(MBR.Mbr_partition[WorstIndex].Part_size[:]))
										convert3, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
										if convert2 < convert3 {
											WorstIndex = i
											break
										}
									}
								}
							}
							copy(MBR.Mbr_partition[WorstIndex].Part_type[:], "P")
							copy(MBR.Mbr_partition[WorstIndex].Part_fit[:], auxFit)
							if WorstIndex == 0 {
								copy(MBR.Mbr_partition[WorstIndex].Part_start[:], string(rune(Tama)))
							} else {
								var unmenos = WorstIndex - 1
								PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
								PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
								total := PartStart + PartSize
								copy(MBR.Mbr_partition[WorstIndex].Part_start[:], string(rune(total)))
							}
							copy(MBR.Mbr_partition[WorstIndex].Part_size[:], string(rune(size_bytes)))
							copy(MBR.Mbr_partition[WorstIndex].Part_status[:], "0")
							copy(MBR.Mbr_partition[WorstIndex].Part_name[:], string(nombre))
							file.Seek(0, 0)
							var bufferControl bytes.Buffer
							binary.Write(&bufferControl, binary.BigEndian, &MBR)
							file.Write(bufferControl.Bytes())
							s := &buffer1
							var binario bytes.Buffer
							binary.Write(&binario, binary.BigEndian, s)
							partstart, _ := strconv.Atoi(string(MBR.Mbr_partition[WorstIndex].Part_start[:]))
							file.Seek(int64(partstart), 0)
							for i := 0; i < size_bytes; i++ {
								_, err := file.Write(binario.Bytes())

								if err != nil {
									fmt.Println(err)
								}
							}
							fmt.Println("-------------------- FDISK --------------------")
							fmt.Println("------------------ Disco Primario Creado ------------------")
							fmt.Println("Part Status: " + string(MBR.Mbr_partition[WorstIndex].Part_status[:]))
							fmt.Println("Part Fit: " + string(MBR.Mbr_partition[WorstIndex].Part_fit[:]))
							fmt.Println("Part Start: " + string(MBR.Mbr_partition[WorstIndex].Part_start[:]))
							fmt.Println("Part Size: " + string(MBR.Mbr_partition[WorstIndex].Part_size[:]))
							fmt.Println("Part name: " + string(MBR.Mbr_partition[WorstIndex].Part_name[:]))
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

func CrearParticionExtendida(direccion string, nombre string, size int, fit string, unit string) {
	var auxFit string = ""
	var auxUnit string = ""
	var auxPath string = direccion
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
		var FlagExtendida bool = false
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
				if string(MBR.Mbr_partition[i].Part_type[:]) == "E" {
					FlagExtendida = true
					break
				}
			}
			if !FlagExtendida {

				for i := 0; i < 4; i++ {
					var sizess string = string(MBR.Mbr_partition[i].Part_size[:])
					convertidos, _ := strconv.Atoi(sizess)
					if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" && convertidos >= size_bytes {
						FlagParticion = true
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
					convertido, _ := strconv.Atoi(string(MBR.Mbr_size[:]))
					if (convertido - EspacioUsado) >= size_bytes {
						if !ExisteParticion(direccion, nombre) {
							if string(MBR.Mbr_disk_fit[:]) == "F" {
								copy(MBR.Mbr_partition[Numero_Particion].Part_type[:], "E")
								copy(MBR.Mbr_partition[Numero_Particion].Part_fit[:], auxFit)
								if Numero_Particion == 0 {
									copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(rune(Tama)))
								} else {
									var unmenos = Numero_Particion - 1
									PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
									PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
									total := PartStart + PartSize
									copy(MBR.Mbr_partition[Numero_Particion].Part_start[:], string(rune(total)))
								}
								copy(MBR.Mbr_partition[Numero_Particion].Part_size[:], string(rune(size_bytes)))
								copy(MBR.Mbr_partition[Numero_Particion].Part_status[:], "0")
								copy(MBR.Mbr_partition[Numero_Particion].Part_name[:], string(nombre))
								file.Seek(0, 0)
								var bufferControl bytes.Buffer
								binary.Write(&bufferControl, binary.BigEndian, &MBR)
								file.Write(bufferControl.Bytes())
								EBR := structs.EBR{}
								copy(EBR.Part_fit[:], auxFit)
								copy(EBR.Part_status[:], "0")
								copy(EBR.Part_start[:], string(MBR.Mbr_partition[Numero_Particion].Part_size[:]))
								copy(EBR.Part_size[:], "0")
								copy(EBR.Part_next[:], "-1")
								copy(EBR.Part_name[:], "")
								seeking, _ := strconv.Atoi(string(MBR.Mbr_partition[Numero_Particion].Part_start[:]))
								file.Seek(int64(seeking), 0)
								var bufferControlEBR bytes.Buffer
								binary.Write(&bufferControlEBR, binary.BigEndian, &EBR)
								file.Write(bufferControl.Bytes())

								s := &buffer1
								var binario bytes.Buffer
								binary.Write(&binario, binary.BigEndian, s)
								var EBRSIZE int = int(unsafe.Sizeof(structs.EBR{}))
								for i := 0; i < (size_bytes - EBRSIZE); i++ {
									_, err := file.Write(binario.Bytes())

									if err != nil {
										fmt.Println(err)
									}
								}
								fmt.Println("-------------------- FDISK --------------------")
								fmt.Println("------------------ Disco Extendido Creado ------------------")
								fmt.Println("Part Status: " + string(MBR.Mbr_partition[Numero_Particion].Part_status[:]))
								fmt.Println("Part Fit: " + string(MBR.Mbr_partition[Numero_Particion].Part_fit[:]))
								fmt.Println("Part Start: " + string(MBR.Mbr_partition[Numero_Particion].Part_start[:]))
								fmt.Println("Part Size: " + string(MBR.Mbr_partition[Numero_Particion].Part_size[:]))
								fmt.Println("Part name: " + string(MBR.Mbr_partition[Numero_Particion].Part_name[:]))
							} else if string(MBR.Mbr_disk_fit[:]) == "B" {

								var BestIndex int = Numero_Particion
								for i := 0; i < 4; i++ {
									convert1, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
									if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" && convert1 >= size_bytes {
										if i != Numero_Particion {
											convert2, _ := strconv.Atoi(string(MBR.Mbr_partition[BestIndex].Part_size[:]))
											convert3, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
											if convert2 > convert3 {
												BestIndex = i
												break
											}
										}
									}
								}
								copy(MBR.Mbr_partition[BestIndex].Part_type[:], "E")
								copy(MBR.Mbr_partition[BestIndex].Part_fit[:], auxFit)
								if BestIndex == 0 {
									copy(MBR.Mbr_partition[BestIndex].Part_start[:], string(rune(Tama)))
								} else {
									var unmenos = BestIndex - 1
									PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
									PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
									total := PartStart + PartSize
									copy(MBR.Mbr_partition[BestIndex].Part_start[:], string(rune(total)))
								}
								copy(MBR.Mbr_partition[BestIndex].Part_size[:], string(rune(size_bytes)))
								copy(MBR.Mbr_partition[BestIndex].Part_status[:], "0")
								copy(MBR.Mbr_partition[BestIndex].Part_name[:], string(nombre))
								file.Seek(0, 0)
								var bufferControl bytes.Buffer
								binary.Write(&bufferControl, binary.BigEndian, &MBR)
								file.Write(bufferControl.Bytes())

								EBR := structs.EBR{}
								copy(EBR.Part_fit[:], auxFit)
								copy(EBR.Part_status[:], "0")
								copy(EBR.Part_start[:], string(MBR.Mbr_partition[BestIndex].Part_size[:]))
								copy(EBR.Part_size[:], "0")
								copy(EBR.Part_next[:], "-1")
								copy(EBR.Part_name[:], "")

								seeking, _ := strconv.Atoi(string(MBR.Mbr_partition[BestIndex].Part_start[:]))
								file.Seek(int64(seeking), 0)
								var bufferControlEBR bytes.Buffer
								binary.Write(&bufferControlEBR, binary.BigEndian, &EBR)
								file.Write(bufferControl.Bytes())

								s := &buffer1
								var binario bytes.Buffer
								binary.Write(&binario, binary.BigEndian, s)
								var EBRSIZE int = int(unsafe.Sizeof(structs.EBR{}))
								for i := 0; i < (size_bytes - EBRSIZE); i++ {
									_, err := file.Write(binario.Bytes())

									if err != nil {
										fmt.Println(err)
									}
								}
								fmt.Println("-------------------- FDISK --------------------")
								fmt.Println("------------------ Disco Extendido Creado ------------------")
								fmt.Println("Part Status: " + string(MBR.Mbr_partition[BestIndex].Part_status[:]))
								fmt.Println("Part Fit: " + string(MBR.Mbr_partition[BestIndex].Part_fit[:]))
								fmt.Println("Part Start: " + string(MBR.Mbr_partition[BestIndex].Part_start[:]))
								fmt.Println("Part Size: " + string(MBR.Mbr_partition[BestIndex].Part_size[:]))
								fmt.Println("Part name: " + string(MBR.Mbr_partition[BestIndex].Part_name[:]))
							} else if string(MBR.Mbr_disk_fit[:]) == "W" {
								var WorstIndex int = Numero_Particion
								for i := 0; i < 4; i++ {
									convert1, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
									if string(MBR.Mbr_partition[i].Part_start[:]) == "-1" || string(MBR.Mbr_partition[i].Part_status[:]) == "1" && convert1 >= size_bytes {
										if i != Numero_Particion {
											convert2, _ := strconv.Atoi(string(MBR.Mbr_partition[WorstIndex].Part_size[:]))
											convert3, _ := strconv.Atoi(string(MBR.Mbr_partition[i].Part_size[:]))
											if convert2 < convert3 {
												WorstIndex = i
												break
											}
										}
									}
								}
								copy(MBR.Mbr_partition[WorstIndex].Part_type[:], "E")
								copy(MBR.Mbr_partition[WorstIndex].Part_fit[:], auxFit)
								if WorstIndex == 0 {
									copy(MBR.Mbr_partition[WorstIndex].Part_start[:], string(rune(Tama)))
								} else {
									var unmenos = WorstIndex - 1
									PartStart, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_start[:]))
									PartSize, _ := strconv.Atoi(string(MBR.Mbr_partition[unmenos].Part_size[:]))
									total := PartStart + PartSize
									copy(MBR.Mbr_partition[WorstIndex].Part_start[:], string(rune(total)))
								}
								copy(MBR.Mbr_partition[WorstIndex].Part_size[:], string(rune(size_bytes)))
								copy(MBR.Mbr_partition[WorstIndex].Part_status[:], "0")
								copy(MBR.Mbr_partition[WorstIndex].Part_name[:], string(nombre))
								file.Seek(0, 0)
								var bufferControl bytes.Buffer
								binary.Write(&bufferControl, binary.BigEndian, &MBR)
								file.Write(bufferControl.Bytes())

								EBR := structs.EBR{}
								copy(EBR.Part_fit[:], auxFit)
								copy(EBR.Part_status[:], "0")
								copy(EBR.Part_start[:], string(MBR.Mbr_partition[WorstIndex].Part_size[:]))
								copy(EBR.Part_size[:], "0")
								copy(EBR.Part_next[:], "-1")
								copy(EBR.Part_name[:], "")

								seeking, _ := strconv.Atoi(string(MBR.Mbr_partition[WorstIndex].Part_start[:]))
								file.Seek(int64(seeking), 0)
								var bufferControlEBR bytes.Buffer
								binary.Write(&bufferControlEBR, binary.BigEndian, &EBR)
								file.Write(bufferControl.Bytes())

								s := &buffer1
								var binario bytes.Buffer
								binary.Write(&binario, binary.BigEndian, s)
								var EBRSIZE int = int(unsafe.Sizeof(structs.EBR{}))
								for i := 0; i < (size_bytes - EBRSIZE); i++ {
									_, err := file.Write(binario.Bytes())

									if err != nil {
										fmt.Println(err)
									}
								}
								fmt.Println("-------------------- FDISK --------------------")
								fmt.Println("------------------ Disco Extendido Creado ------------------")
								fmt.Println("Part Status: " + string(MBR.Mbr_partition[WorstIndex].Part_status[:]))
								fmt.Println("Part Fit: " + string(MBR.Mbr_partition[WorstIndex].Part_fit[:]))
								fmt.Println("Part Start: " + string(MBR.Mbr_partition[WorstIndex].Part_start[:]))
								fmt.Println("Part Size: " + string(MBR.Mbr_partition[WorstIndex].Part_size[:]))
								fmt.Println("Part name: " + string(MBR.Mbr_partition[WorstIndex].Part_name[:]))
							}
						}
					}
				}
			}
		}
	}
}

func BuscarNumero(direccion string, nombre string) int {
	var retorno int = 1
	for i := len(ListaMount) - 1; i >= 0; i-- {
		if direccion == ListaMount[i].direccion && nombre == ListaMount[i].nombre {
			return -1
		} else {
			if direccion == ListaMount[i].direccion {
				return ListaMount[i].num
			} else if retorno <= ListaMount[i].num {
				retorno++
			}
		}
	}
	return retorno
}

func BuscarLetra(direccion string) string {
	listaLetras := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var retornos int = 1
	for i := len(ListaMount) - 1; i >= 0; i-- {
		if direccion == ListaMount[i].direccion && listaLetras[retornos] == ListaMount[i].letra {
			retornos++
		}
	}
	return string(listaLetras[retornos])
}

func EXT2(inicio int, tamano int, direccion string) {
	n := float64(tamano-int(unsafe.Sizeof(structs.SuperBloque{}))) / float64(4+int(unsafe.Sizeof(structs.InodoTable{}))+3*int(unsafe.Sizeof(structs.BloqueArchivo{})))
	var num_estructuras int = int(math.Floor(n))
	var num_bloques int = int(3 * num_estructuras)

	super := structs.SuperBloque{}

	super.S_filesystem_type = 2
	super.S_inodes_count = uint64(num_estructuras)
	super.S_blocks_count = uint64(num_bloques)
	super.S_free_blocks_count = uint64(num_bloques - 2)
	super.S_free_inodes_count = uint64(num_estructuras - 2)
	copy(super.S_mtime[:], ObtenerFecha())
	copy(super.S_umtime[:], "0")
	super.S_mnt_count = 0
	super.S_magic = uint64(0xEF53)
	super.S_inode_size = uint64(unsafe.Sizeof(structs.InodoTable{}))
	super.S_block_size = uint64(unsafe.Sizeof(structs.BloqueArchivo{}))
	super.S_first_ino = 2
	super.S_first_blo = 2
	super.S_bm_inode_start = uint64(inicio + int(unsafe.Sizeof(structs.SuperBloque{})))
	super.S_bm_inode_start = uint64(inicio + int(unsafe.Sizeof(structs.SuperBloque{})) + num_estructuras)
	super.S_inode_start = uint64(inicio + int(unsafe.Sizeof(structs.SuperBloque{})) + num_estructuras + num_bloques)
	super.S_block_start = uint64(inicio + int(unsafe.Sizeof(structs.SuperBloque{})) + num_estructuras + num_bloques + (int(unsafe.Sizeof(structs.SuperBloque{})) * num_estructuras))

	inodo := structs.InodoTable{}
	bloque := structs.BloqueCarpeta{}
	var buffer int8 = 0
	var buffer2 int8 = 1
	var buffer3 int8 = 2
	file, _ := os.OpenFile(direccion, os.O_RDWR, 0777)
	defer file.Close()
	/* -----------Super Bloque------------- */
	file.Seek(int64(inicio), 0)
	var bufferControl bytes.Buffer
	binary.Write(&bufferControl, binary.BigEndian, &super)
	_, errs := file.Write(bufferControl.Bytes())
	if errs != nil {
		fmt.Println("ERROR WE")
	}
	fmt.Println("-------------------- Super Bloque --------------------")
	fmt.Println("S_filesystem_type: " + string(rune(super.S_filesystem_type)))
	fmt.Println("S_inodes_count: " + string(rune(super.S_inodes_count)))
	fmt.Println("S_blocks_count: " + string(rune(super.S_blocks_count)))
	fmt.Println("S_free_blocks_count: " + string(rune(super.S_free_blocks_count)))
	fmt.Println("S_free_inodes_count: " + string(rune(super.S_free_inodes_count)))
	fmt.Println("S_mtime: " + string(super.S_mtime[:]))
	fmt.Println("S_umtime: " + string(super.S_umtime[:]))
	fmt.Println("S_mnt_count: " + string(rune(super.S_mnt_count)))
	fmt.Println("S_magic: " + string("0xEF53"))
	fmt.Println("S_inode_size: " + string(rune(super.S_inode_size)))
	fmt.Println("S_block_size: " + string(rune(super.S_block_size)))
	fmt.Println("S_first_ino: " + string(rune(super.S_first_ino)))
	fmt.Println("S_first_blo: " + string(rune(super.S_first_blo)))
	fmt.Println("S_bm_inode_start: " + string(rune(super.S_bm_inode_start)))
	fmt.Println("S_bm_block_start: " + string(rune(super.S_bm_block_start)))
	fmt.Println("S_inode_start: " + string(rune(super.S_inode_start)))
	fmt.Println("S_block_start: " + string(rune(super.S_block_start)))

	/* -----------BitMap de Inodos------------- */
	s := &buffer
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	for i := 0; i < num_estructuras; i++ {
		file.Seek(int64(int(super.S_bm_inode_start)+i), 0)
		_, err := file.Write(binario.Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}
	/* -----------bit para / y users.txt en BM------------- */
	file.Seek(int64(super.S_bm_inode_start), 0)
	s2 := &buffer2
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s2)
	_, err := file.Write(binario2.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	binary.Write(&binario2, binary.BigEndian, s2)
	_, err1 := file.Write(binario2.Bytes())
	if err1 != nil {
		fmt.Println(err1)
	}
	/* -----------BitMap de Bloques------------- */
	binary.Write(&binario, binary.BigEndian, s)
	for i := 0; i < num_bloques; i++ {
		file.Seek(int64(int(super.S_bm_block_start)+i), 0)
		_, err := file.Write(binario.Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}
	/* -----------bit para / y users.txt en BM------------- */
	file.Seek(int64(super.S_bm_block_start), 0)
	binary.Write(&binario2, binary.BigEndian, s2)
	_, err2 := file.Write(binario2.Bytes())
	if err2 != nil {
		fmt.Println(err2)
	}
	s3 := &buffer3
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, s3)
	_, err3 := file.Write(binario3.Bytes())
	if err3 != nil {
		fmt.Println(err3)
	}
	/* -----------inodo para carpeta root------------- */
	inodo.I_uid = 1
	inodo.I_gid = 1
	inodo.I_size = 0
	copy(inodo.I_atime[:], ObtenerFecha())
	copy(inodo.I_ctime[:], ObtenerFecha())
	copy(inodo.I_mtime[:], ObtenerFecha())
	inodo.I_block[0] = 0
	for i := 1; i < 15; i++ {
		inodo.I_block[i] = -1
	}
	copy(inodo.I_type[:], "0")
	inodo.I_perm = 664
	file.Seek(int64(super.S_inode_start), 0)
	var bufferControlSI bytes.Buffer
	binary.Write(&bufferControlSI, binary.BigEndian, &inodo)
	_, errsSI := file.Write(bufferControlSI.Bytes())
	if errsSI != nil {
		fmt.Println("ERROR WE")
	}
	fmt.Println("-------------------- Inodo Carpeta Root --------------------")
	fmt.Println("I_uid: " + string(rune(inodo.I_uid)))
	fmt.Println("I_gid: " + string(rune(inodo.I_gid)))
	fmt.Println("I_size: " + string(rune(inodo.I_size)))
	fmt.Println("I_type: " + string(inodo.I_type[:]))
	fmt.Println("I_perm: " + string(rune(inodo.I_perm)))
	fmt.Println("I_atime: " + string(inodo.I_atime[:]))
	fmt.Println("I_ctime: " + string(inodo.I_ctime[:]))
	fmt.Println("I_mtime: " + string(inodo.I_mtime[:]))
	/* -----------Bloque para la carpeta root------------- */
	copy(bloque.B_content[0].B_name[:], ".")
	bloque.B_content[0].B_inodo = 0

	copy(bloque.B_content[1].B_name[:], "..")
	bloque.B_content[1].B_inodo = 0

	copy(bloque.B_content[2].B_name[:], "users.txt")
	bloque.B_content[2].B_inodo = 1

	copy(bloque.B_content[0].B_name[:], ".")
	bloque.B_content[0].B_inodo = -1

	file.Seek(int64(super.S_block_start), 0)
	var bufferControlB bytes.Buffer
	binary.Write(&bufferControlB, binary.BigEndian, &bloque)
	_, errB := file.Write(bufferControlB.Bytes())
	if errB != nil {
		fmt.Println("ERROR WE")
	}
	/* -----------inodo para carpeta users.txt------------- */
	inodo.I_uid = 1
	inodo.I_gid = 1
	inodo.I_size = 27
	copy(inodo.I_atime[:], ObtenerFecha())
	copy(inodo.I_ctime[:], ObtenerFecha())
	copy(inodo.I_mtime[:], ObtenerFecha())
	inodo.I_block[0] = 1
	for i := 1; i < 15; i++ {
		inodo.I_block[i] = -1
	}
	copy(inodo.I_type[:], "1")
	inodo.I_perm = 755
	sizesInodeTable := uint64(unsafe.Sizeof(structs.InodoTable{}))
	file.Seek(int64(super.S_inode_start+sizesInodeTable), 0)
	var bufferControl2 bytes.Buffer
	binary.Write(&bufferControl2, binary.BigEndian, &inodo)
	_, errs2 := file.Write(bufferControl2.Bytes())
	if errs2 != nil {
		fmt.Println("ERROR WE")
	}
	fmt.Println("-------------------- Inodo Carpeta users.txt --------------------")
	fmt.Println("I_uid: " + string(rune(inodo.I_uid)))
	fmt.Println("I_gid: " + string(rune(inodo.I_gid)))
	fmt.Println("I_size: " + string(rune(inodo.I_size)))
	fmt.Println("I_type: " + string(inodo.I_type[:]))
	fmt.Println("I_perm: " + string(rune(inodo.I_perm)))
	fmt.Println("I_atime: " + string(inodo.I_atime[:]))
	fmt.Println("I_ctime: " + string(inodo.I_ctime[:]))
	fmt.Println("I_mtime: " + string(inodo.I_mtime[:]))
	/* -----------Bloque para users.txt------------- */
	archivo := structs.BloqueArchivo{}
	copy(archivo.B_content[:], "1,G,root\n1,U,root,root,123\n")
	sizesBloqueCarpeta := uint64(unsafe.Sizeof(structs.BloqueCarpeta{}))
	file.Seek(int64(super.S_inode_start+sizesBloqueCarpeta), 0)
	var bufferControl3 bytes.Buffer
	binary.Write(&bufferControl3, binary.BigEndian, &archivo)
	_, errs3 := file.Write(bufferControl3.Bytes())
	if errs3 != nil {
		fmt.Println("ERROR WE")
	}
	fmt.Println("-------------------- Bloque Users.txt --------------------")
	fmt.Println("B_content: 1,G,root\n1,U,root,root,123")

}

// mkdisk("5", "wf", "m", "/home/curious1924/Escritorio/Joder/disco3.dk")
func Mkdisk(size string, fit string, unit string, path string) {
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
				copy(MBR.Mbr_date_created[:], string(fechas))
				rand.Seed(int64(time.Now().UnixNano()))
				nuevoNumero := rand.Int()
				nuevs := strconv.Itoa(nuevoNumero)
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

				file.Seek(0, 0)
				var bufferControl bytes.Buffer
				binary.Write(&bufferControl, binary.BigEndian, &MBR)
				_, errs := file.Write(bufferControl.Bytes())
				if errs != nil {
					fmt.Println("ERROR WE")
				}
				file.Seek(0, 0)
				var temporal int8 = 0
				s := &temporal
				var binario bytes.Buffer
				binary.Write(&binario, binary.BigEndian, s)
				for i := 0; i < total_size; i++ {
					EscribirBytes(file, binario.Bytes())
				}
				fmt.Println("Disco creado con exito!")
				fmt.Println("-------------------- MKDISK --------------------")
				fmt.Println("------------------ Disco Creado ------------------")
				fmt.Println("MBR Size: " + string(MBR.Mbr_size[:]))
				fmt.Println("MBR Date Created: " + string(MBR.Mbr_date_created[:]))
				fmt.Println("MBR disk signature: " + string(MBR.Mbr_disk_signature[:]))
				fmt.Println("MBR disk fit: " + string(MBR.Mbr_disk_fit[:]))
			} else {
				fmt.Println("ERROR Parametro -size no definido ")
			}
		} else {
			fmt.Println("ERROR Parametro -path no definido ")
		}
	}
}

func Rmdisk(path string) {
	fmt.Println("Esta seguro que desea eliminar el disco? S/N :")
	reader := bufio.NewReader(os.Stdin)
	comando, _ := reader.ReadString('\n')
	if strings.Contains(comando, "S\n") || strings.Contains(comando, "s\n") {
		_ = os.Remove(path)
		fmt.Println("-------------------- RMDISK --------------------")
		fmt.Println("Disco eliminado con exito!")
	} else if strings.Contains(comando, "N\n") || strings.Contains(comando, "n\n") {
		fmt.Println("Cancelado con exito!")
	} else {
		fmt.Println("Opcion incorrecta")
	}
}

func Fdisk(size string, fit string, unit string, path string, types string, name string) {
	var BanderaSize bool = false
	var BanderaPath bool = false
	var BanderaType bool = false
	var BanderaName bool = false
	var Bandera bool = false
	var ValSize int = 0
	var ValFit string = fit
	var ValUnit string = unit
	var ValPath string = path
	var ValType string = types
	var ValName string = name
	if size != "" {
		Prueba, err := strconv.Atoi(size)
		ValSize = Prueba
		fmt.Println(err)
		BanderaSize = true
		if ValSize < 0 {
			fmt.Println("ERROR: parametro -size menor a cero")
			Bandera = true
		}
	}
	if path != "" {
		ValPath = path
		BanderaPath = true
		if ValPath == "" {
			fmt.Println("ERROR: Valor del parametro path vacio")
			Bandera = true
		}
	}
	if fit != "" {
		if fit == "b" {
			ValFit = "B"
		} else if fit == "f" {
			ValFit = "F"
		} else if fit == "w" {
			ValFit = "W"
		}
	}
	if unit != "" {
		if unit == "b" || unit == "B" {
			ValUnit = "b"
		} else if unit == "k" || unit == "K" {
			ValUnit = "k"
		} else if unit == "m" || unit == "M" {
			ValUnit = "m"
		} else {
			fmt.Println("ERROR: Valor del parametro unit no reconocido")
			Bandera = true
		}
	}
	if types != "" {
		BanderaType = true
		ValType = types
		if ValType == "P" || ValType == "p" {
			ValType = "P"
		} else if ValType == "E" || ValType == "e" {
			ValType = "E"
		} else if ValType == "L" || ValType == "l" {
			ValType = "L"
		} else {
			fmt.Println("ERROR: Valor del parametro type no reconocido")
			Bandera = true
		}
	}
	if name != "" {
		ValName = name
		BanderaName = true
		if ValName == "" {
			fmt.Println("ERROR: Valor del parametro name vacio")
			Bandera = true
		}
	}
	if !Bandera {
		if BanderaPath {
			if BanderaName {
				if BanderaSize {
					if BanderaType {
						if ValType == "P" {
							CrearParticionPrimaria(ValPath, ValName, ValSize, ValFit, ValUnit)
						} else if ValType == "E" {
							CrearParticionExtendida(ValPath, ValName, ValSize, ValFit, ValUnit)
						} else if ValType == "L" {
							fmt.Println("Aqui va la particion logica")
						}
					} else {
						CrearParticionPrimaria(ValPath, ValName, ValSize, ValFit, ValUnit)
					}
				}
			} else {
				fmt.Println("ERROR parametro -name no definido")
			}
		} else {
			fmt.Println("ERROR parametro -path no definido")
		}
	}
}

func Mount(path string, name string) {
	var BanderaPath bool = false
	var BanderaName bool = false
	var Bandera bool = false
	var ValPath string = ""
	var ValName string = ""
	if path != "" {
		BanderaPath = true
		ValPath = path
	}
	if name != "" {
		BanderaName = true
		ValName = name
	}
	if !Bandera {
		if BanderaPath {
			if BanderaName {
				IndexP := BuscarParticion_Primaria_Extendida(ValPath, ValName)
				if IndexP != -1 {
					file, err := os.OpenFile(ValPath, os.O_RDWR, 0777)
					defer file.Close()
					if err != nil {
						fmt.Println("ERROR no se encuentra el disco")
					} else {
						MBR := structs.MBR{}
						file.Seek(0, 0)
						var sizetemp int = int(unsafe.Sizeof(MBR))
						data := leerBytes(file, sizetemp)
						buffer := bytes.NewBuffer(data)
						Errores := binary.Read(buffer, binary.BigEndian, &MBR)
						if Errores != nil {
							fmt.Println("F perro")
						}
						copy(MBR.Mbr_partition[IndexP].Part_status[:], "2")
						file.Seek(0, 0)
						var bufferControl bytes.Buffer
						binary.Write(&bufferControl, binary.BigEndian, &MBR)
						_, errs := file.Write(bufferControl.Bytes())
						if errs != nil {
							fmt.Println("ERROR WE")
						}
						nums := BuscarNumero(ValPath, ValName)
						if nums == -1 {
							fmt.Println("ERROR la particion ya se encuentra  montada, no se puede montar de nuevo")
						} else {
							var letra string = BuscarLetra(ValPath)
							var ids string = "07"
							ids += string(rune(nums)) + string(letra)
							nodonuevo := NodoMount{direccion: ValPath, nombre: ValName, id: ids, letra: letra, num: nums}
							ListaMount = append(ListaMount, nodonuevo)
							fmt.Println("Particion Montada con exito!")
							fmt.Println()
							fmt.Println("*--------- Particiones Montadas ---------*")
							fmt.Println("*               Name | ID                *")
							fmt.Println("*----------------------------------------*")
							for i := len(ListaMount) - 1; i >= 0; i-- {
								fmt.Println("*     " + string(ListaMount[i].nombre) + "    |     " + string(ListaMount[i].id) + "     *")
							}
						}
					}
				} else {
					fmt.Println("Aqui van las particiones logicas")
				}
			}
		}
	}
}

func Mkfs(ids string, typess string) {
	var BanderaID bool = false
	var id string = ""
	var banderaEncontrado bool = false
	var types string = ""

	if ids != "" {
		BanderaID = true
		id = ids
	}
	if typess != "" {
		types = typess
	}
	fmt.Println(types)
	if BanderaID {
		var nodo = NodoMount{}
		for i := len(ListaMount) - 1; i >= 0; i-- {
			if ListaMount[i].id == id {
				nodo = ListaMount[i]
				banderaEncontrado = true
				break
			}
		}
		if banderaEncontrado {
			index := BuscarParticion_Primaria_Extendida(nodo.direccion, nodo.nombre)
			if index != -1 {
				MBR := structs.MBR{}
				file, err := os.OpenFile(nodo.direccion, os.O_RDWR, 0777)
				if err != nil {
					fmt.Println("Error al abrir el archivo")
					fmt.Println(err)
				}
				defer file.Close()
				file.Seek(0, 0)
				TempMBR := structs.MBR{}
				var sizetemp int = int(unsafe.Sizeof(TempMBR))
				data := leerBytes(file, sizetemp)
				buffer := bytes.NewBuffer(data)
				Errores := binary.Read(buffer, binary.BigEndian, &MBR)
				if Errores != nil {
					fmt.Println("F perro no se pudo leer el mbr")
				}
				inicio, _ := strconv.Atoi(string(MBR.Mbr_partition[index].Part_start[:]))
				tamano, _ := strconv.Atoi(string(MBR.Mbr_partition[index].Part_size[:]))
				EXT2(inicio, tamano, nodo.direccion)
			}
		}
	}
}

func LogOut() {
	if BanderaLogin {
		BanderaLogin = false
		CurrentSession.Id_user = -1
		copy(CurrentSession.Direccion[:], "")
		CurrentSession.InicioSuper = -1
		fmt.Println("El usuario se deslogueo, ya no se encuentra la session activa")
	} else {
		fmt.Println("ERROR no hay ninguna sesion activa")
	}
}

func getc(f *os.File) (byte, error) {
	b := make([]byte, 1)
	_, err := f.Read(b)
	return b[0], err
}

func BuscarExisteGrupo(name string) int {
	file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
	defer file.Close()
	var cadena string = ""
	super := structs.SuperBloque{}
	inodo := structs.InodoTable{}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempSB := structs.SuperBloque{}
	var sizetemp int = int(unsafe.Sizeof(TempSB))
	data := leerBytes(file, sizetemp)
	buffer := bytes.NewBuffer(data)
	Errores := binary.Read(buffer, binary.BigEndian, &super)
	if Errores != nil {
		fmt.Println("F perro")
	}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempTI := structs.SuperBloque{}
	var sizetemp1 int = int(unsafe.Sizeof(TempTI))
	data1 := leerBytes(file, sizetemp1)
	buffer1 := bytes.NewBuffer(data1)
	Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
	if Errores1 != nil {
		fmt.Println("F perro")
	}

	for i := 0; i < 15; i++ {
		if inodo.I_block[i] != -1 {
			archivo := structs.BloqueArchivo{}
			file.Seek(int64(super.S_block_start), 0)
			for j := 0; j <= int(inodo.I_block[i]); j++ {
				temparchivo := structs.BloqueArchivo{}
				var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
				data2 := leerBytes(file, sizetemp2)
				buffer2 := bytes.NewBuffer(data2)
				Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
				if Errores2 != nil {
					fmt.Println("F perro")
				}
			}
			cadena += string(archivo.B_content[:])
		}
	}
	token := strings.Split(cadena, "\n")
	for i := 0; i < len(token); i++ {
		id := ""
		tipo := ""
		group := ""
		token2 := strings.Split(token[i], ",")
		for i := 0; i < len(token2); i++ {
			id = token2[0]
			if id != "0" {
				tipo = token2[1]
				if tipo == "G" {
					group = token2[2]
					if strings.Compare(group, name) == 0 {
						idEnvio, _ := strconv.Atoi(id)
						return idEnvio
					}
				}
			}
		}
	}
	return -1
}

func BuscarExisteUsuario(name string) bool {
	file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
	defer file.Close()
	var cadena string = ""
	super := structs.SuperBloque{}
	inodo := structs.InodoTable{}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempSB := structs.SuperBloque{}
	var sizetemp int = int(unsafe.Sizeof(TempSB))
	data := leerBytes(file, sizetemp)
	buffer := bytes.NewBuffer(data)
	Errores := binary.Read(buffer, binary.BigEndian, &super)
	if Errores != nil {
		fmt.Println("F perro")
	}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempTI := structs.SuperBloque{}
	var sizetemp1 int = int(unsafe.Sizeof(TempTI))
	data1 := leerBytes(file, sizetemp1)
	buffer1 := bytes.NewBuffer(data1)
	Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
	if Errores1 != nil {
		fmt.Println("F perro")
	}

	for i := 0; i < 15; i++ {
		if inodo.I_block[i] != -1 {
			archivo := structs.BloqueArchivo{}
			file.Seek(int64(super.S_block_start), 0)
			for j := 0; j <= int(inodo.I_block[i]); j++ {
				temparchivo := structs.BloqueArchivo{}
				var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
				data2 := leerBytes(file, sizetemp2)
				buffer2 := bytes.NewBuffer(data2)
				Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
				if Errores2 != nil {
					fmt.Println("F perro")
				}
			}
			cadena += string(archivo.B_content[:])
		}
	}
	token := strings.Split(cadena, "\n")
	for i := 0; i < len(token); i++ {
		id := ""
		tipo := ""
		usuario := ""
		token2 := strings.Split(token[i], ",")
		for i := 0; i < len(token2); i++ {
			id = token2[0]
			if id != "0" {
				tipo = token2[1]
				if tipo == "U" {
					usuario = token2[3]
					if strings.Compare(usuario, name) == 0 {
						return true
					}
				}
			}
		}
	}
	return false
}

func GetId_Grp() int {
	file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
	defer file.Close()
	var cadena string = ""
	var aux_id int = -1
	super := structs.SuperBloque{}
	inodo := structs.InodoTable{}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempSB := structs.SuperBloque{}
	var sizetemp int = int(unsafe.Sizeof(TempSB))
	data := leerBytes(file, sizetemp)
	buffer := bytes.NewBuffer(data)
	Errores := binary.Read(buffer, binary.BigEndian, &super)
	if Errores != nil {
		fmt.Println("F perro")
	}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempTI := structs.SuperBloque{}
	var sizetemp1 int = int(unsafe.Sizeof(TempTI))
	data1 := leerBytes(file, sizetemp1)
	buffer1 := bytes.NewBuffer(data1)
	Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
	if Errores1 != nil {
		fmt.Println("F perro")
	}

	for i := 0; i < 15; i++ {
		if inodo.I_block[i] != -1 {
			archivo := structs.BloqueArchivo{}
			file.Seek(int64(super.S_block_start), 0)
			for j := 0; j <= int(inodo.I_block[i]); j++ {
				temparchivo := structs.BloqueArchivo{}
				var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
				data2 := leerBytes(file, sizetemp2)
				buffer2 := bytes.NewBuffer(data2)
				Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
				if Errores2 != nil {
					fmt.Println("F perro")
				}
			}
			cadena += string(archivo.B_content[:])
		}
	}
	token := strings.Split(cadena, "\n")
	for i := 0; i < len(token); i++ {
		id := ""
		tipo := ""
		token2 := strings.Split(token[i], ",")
		for i := 0; i < len(token2); i++ {
			id = token2[0]
			if id != "0" {
				tipo = token2[1]
				if tipo == "G" {
					enviowe, _ := strconv.Atoi(id)
					aux_id = enviowe

				}
			}
		}
	}
	aux_id += 1
	return aux_id
}

func GetId_User() int {
	file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
	defer file.Close()
	var cadena string = ""
	var aux_id int = 0
	super := structs.SuperBloque{}
	inodo := structs.InodoTable{}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempSB := structs.SuperBloque{}
	var sizetemp int = int(unsafe.Sizeof(TempSB))
	data := leerBytes(file, sizetemp)
	buffer := bytes.NewBuffer(data)
	Errores := binary.Read(buffer, binary.BigEndian, &super)
	if Errores != nil {
		fmt.Println("F perro")
	}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempTI := structs.SuperBloque{}
	var sizetemp1 int = int(unsafe.Sizeof(TempTI))
	data1 := leerBytes(file, sizetemp1)
	buffer1 := bytes.NewBuffer(data1)
	Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
	if Errores1 != nil {
		fmt.Println("F perro")
	}

	for i := 0; i < 15; i++ {
		if inodo.I_block[i] != -1 {
			archivo := structs.BloqueArchivo{}
			file.Seek(int64(super.S_block_start), 0)
			for j := 0; j <= int(inodo.I_block[i]); j++ {
				temparchivo := structs.BloqueArchivo{}
				var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
				data2 := leerBytes(file, sizetemp2)
				buffer2 := bytes.NewBuffer(data2)
				Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
				if Errores2 != nil {
					fmt.Println("F perro")
				}
			}
			cadena += string(archivo.B_content[:])
		}
	}
	token := strings.Split(cadena, "\n")
	for i := 0; i < len(token); i++ {
		id := ""
		tipo := ""
		token2 := strings.Split(token[i], ",")
		for i := 0; i < len(token2); i++ {
			id = token2[0]
			if id != "0" {
				tipo = token2[1]
				if tipo == "U" {
					enviowe, _ := strconv.Atoi(id)
					aux_id = enviowe + 1

				}
			}
		}
	}
	aux_id += 1
	return aux_id
}

func VerificarDatos(user string, password string, direccion string) int {
	file, _ := os.OpenFile(direccion, os.O_RDWR, 0777)
	defer file.Close()
	var cadena string = ""
	super := structs.SuperBloque{}
	inodo := structs.InodoTable{}

	file.Seek(CurrentSession.InicioSuper, 0)
	TempSB := structs.SuperBloque{}
	var sizetemp int = int(unsafe.Sizeof(TempSB))
	data := leerBytes(file, sizetemp)
	buffer := bytes.NewBuffer(data)
	Errores := binary.Read(buffer, binary.BigEndian, &super)
	if Errores != nil {
		fmt.Println("F perro")
	}

	file.Seek(int64(super.S_inode_start+uint64(unsafe.Sizeof(structs.InodoTable{}))), 0)
	TempTI := structs.InodoTable{}
	var sizetemp1 int = int(unsafe.Sizeof(TempTI))
	data1 := leerBytes(file, sizetemp1)
	buffer1 := bytes.NewBuffer(data1)
	Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
	if Errores1 != nil {
		fmt.Println("F perro")
	}

	for i := 0; i < 15; i++ {
		if inodo.I_block[i] != -1 {
			archivo := structs.BloqueArchivo{}
			file.Seek(int64(super.S_block_start), 0)
			for j := 0; j <= int(inodo.I_block[i]); j++ {
				temparchivo := structs.BloqueArchivo{}
				var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
				data2 := leerBytes(file, sizetemp2)
				buffer2 := bytes.NewBuffer(data2)
				Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
				if Errores2 != nil {
					fmt.Println("F perro")
				}
			}
			cadena += string(archivo.B_content[:])
		}
	}
	token := strings.Split(cadena, "\n")
	for i := 0; i < len(token); i++ {
		id := ""
		tipo := ""
		group := ""
		user_ := ""
		password_ := ""
		fmt.Println(token[i])
		token2 := strings.Split(token[i], ",")
		for i := 0; i < len(token2); i++ {
			id = token2[0]
			if id != "0" {
				tipo = token2[1]
				if tipo == "U" {
					group = token2[2]
					user_ = token2[3]
					password_ = token2[4]
					if strings.Compare(user_, user) == 0 {
						if strings.Compare(password_, password) == 0 {
							idEnvio, _ := strconv.Atoi(id)
							copy(CurrentSession.Direccion[:], []byte(direccion))
							CurrentSession.Id_user = int64(idEnvio)
							CurrentSession.Id_grp = int64(BuscarExisteGrupo(group))
							return 1
						} else {
							return 2
						}
					}
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	return 2
}

func PreLogin(direccion string, nombre string, user string, password string) int {
	var index int = BuscarParticion_Primaria_Extendida(direccion, nombre)
	if index != 0 {
		MBR := structs.MBR{}
		super := structs.SuperBloque{}
		inodo := structs.InodoTable{}
		file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
		defer file.Close()

		file.Seek(0, 0)
		TempMBR := structs.MBR{}
		var sizetempMBR int = int(unsafe.Sizeof(TempMBR))
		data := leerBytes(file, sizetempMBR)
		buffer := bytes.NewBuffer(data)
		Errores := binary.Read(buffer, binary.BigEndian, &MBR)
		if Errores != nil {
			fmt.Println("F perro")
		}

		PartStartMBR, _ := strconv.Atoi(string(MBR.Mbr_partition[index].Part_start[:]))
		file.Seek(int64(PartStartMBR), 0)
		Tempsuper := structs.SuperBloque{}
		var sizetempsuper int = int(unsafe.Sizeof(Tempsuper))
		datasuper := leerBytes(file, sizetempsuper)
		buffersuper := bytes.NewBuffer(datasuper)
		Erroressuper := binary.Read(buffersuper, binary.BigEndian, &super)
		if Erroressuper != nil {
			fmt.Println("F perro")
		}

		file.Seek(int64(super.S_bm_inode_start+uint64(unsafe.Sizeof(structs.InodoTable{}))), 0)
		TempInodeTable := structs.SuperBloque{}
		var sizetempInodeTable int = int(unsafe.Sizeof(TempInodeTable))
		dataInodeTable := leerBytes(file, sizetempInodeTable)
		bufferInodeTable := bytes.NewBuffer(dataInodeTable)
		ErroresInodeTable := binary.Read(bufferInodeTable, binary.BigEndian, &inodo)
		if ErroresInodeTable != nil {
			fmt.Println("F perro")
		}
		file.Seek(int64(super.S_bm_inode_start+uint64(unsafe.Sizeof(structs.InodoTable{}))), 0)
		copy(inodo.I_mtime[:], ObtenerFecha())
		var bufferControl2 bytes.Buffer
		binary.Write(&bufferControl2, binary.BigEndian, &inodo)
		_, errs2 := file.Write(bufferControl2.Bytes())
		if errs2 != nil {
			fmt.Println("ERROR WE")
		}
		currentSessionInicioSuper, _ := strconv.Atoi(string(MBR.Mbr_partition[index].Part_start[:]))
		CurrentSession.InicioSuper = int64(currentSessionInicioSuper)
		CurrentSession.Fit = MBR.Mbr_partition[index].Part_fit
		CurrentSession.Tipo_sistema = 2
		return VerificarDatos(user, password, direccion)
	}
	return -1
}

func Login(users string, passwords string, ids string) {
	var BanderaUser bool = false
	var BanderaPassword bool = false
	var BanderaID bool = false
	var Bandera bool = false
	var user string = ""
	var password string = ""
	var id string = ""

	if users != "" {
		BanderaUser = true
		user = users

	} else {
		Bandera = true
	}
	if passwords != "" {
		BanderaPassword = true
		password = passwords

	} else {
		Bandera = true
	}
	if ids != "" {
		BanderaID = true
		id = ids

	} else {
		Bandera = true
	}
	if !Bandera {
		if BanderaUser {
			if BanderaPassword {
				if BanderaID {
					if !BanderaLogin {
						for i := len(ListaMount) - 1; i >= 0; i-- {
							nuevoNodo := NodoMount{}
							if ListaMount[i].id == id {
								nuevoNodo = ListaMount[i]
							}
							if nuevoNodo.id != "" {
								res := PreLogin(nuevoNodo.direccion, nuevoNodo.nombre, user, password)
								if res == 1 {
									BanderaLogin = true
									fmt.Println("Sesion iniciada con exito")
								} else if res == 2 {
									fmt.Println("ERROR Contrasena incorrecta")
								} else if res == 0 {
									fmt.Println("ERROR Usuario no encontrado")
								} else {
									fmt.Println("ERROR Desconocido")
								}
							} else {
								fmt.Println("ERROR no se encuentra ninguna particion montada con ese id ")
							}
						}
					} else {
						fmt.Println("ERROR sesion activa, cierre sesion para poder volver a iniciar sesion")
					}
				} else {
					fmt.Println("ERROR parametro -id no definido")
				}
			} else {
				fmt.Println("ERROR parametro -password no definido")
			}
		} else {
			fmt.Println("ERROR parametro -usuario no definido")
		}
	}
}

func Mkusr(name string, passwords string, group string) {
	var BanderaUser bool = false
	var BanderaPassword bool = false
	var BanderaGroup bool = false
	var ReEscritura string = ""
	var users string = ""
	var password string = ""
	var groups string = ""
	if name != "" {
		BanderaUser = true
		users = name
	}
	if passwords != "" {
		BanderaPassword = true
		password = passwords
	}
	if group != "" {
		BanderaGroup = true
		groups = group
	}
	if BanderaUser {
		if BanderaPassword {
			if BanderaGroup {
				if len(users) <= 10 {
					if len(password) <= 10 {
						if len(groups) <= 10 {
							if BanderaLogin {
								if CurrentSession.Id_user == 1 && CurrentSession.Id_grp == 1 {
									grupo := BuscarExisteGrupo(name)
									if grupo != -1 {
										usuarios := BuscarExisteUsuario(name)
										if !usuarios {
											idUser := GetId_User()
											file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
											defer file.Close()
											var cadena string = ""
											super := structs.SuperBloque{}
											inodo := structs.InodoTable{}

											file.Seek(CurrentSession.InicioSuper, 0)
											TempSB := structs.SuperBloque{}
											var sizetemp int = int(unsafe.Sizeof(TempSB))
											data := leerBytes(file, sizetemp)
											buffer := bytes.NewBuffer(data)
											Errores := binary.Read(buffer, binary.BigEndian, &super)
											if Errores != nil {
												fmt.Println("F perro")
											}

											file.Seek(CurrentSession.InicioSuper, 0)
											TempTI := structs.SuperBloque{}
											var sizetemp1 int = int(unsafe.Sizeof(TempTI))
											data1 := leerBytes(file, sizetemp1)
											buffer1 := bytes.NewBuffer(data1)
											Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
											if Errores1 != nil {
												fmt.Println("F perro")
											}

											for i := 0; i < 15; i++ {
												if inodo.I_block[i] != -1 {
													archivo := structs.BloqueArchivo{}
													file.Seek(int64(super.S_block_start), 0)
													for j := 0; j <= int(inodo.I_block[i]); j++ {
														temparchivo := structs.BloqueArchivo{}
														var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
														data2 := leerBytes(file, sizetemp2)
														buffer2 := bytes.NewBuffer(data2)
														Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
														if Errores2 != nil {
															fmt.Println("F perro")
														}
													}
													cadena += string(archivo.B_content[:])
												}
											}
											nuevoId := strconv.Itoa(idUser)
											NuevVar := nuevoId + ",U," + groups + "," + users + "," + password + "\n"
											ReEscritura = cadena + NuevVar
											libre := 127 - len(cadena)
											if len(NuevVar) <= libre {
												archivo := structs.BloqueArchivo{}
												copy(archivo.B_content[:], ReEscritura)
												sizesBloqueCarpeta := uint64(unsafe.Sizeof(structs.BloqueCarpeta{}))
												file.Seek(int64(super.S_inode_start+sizesBloqueCarpeta), 0)
												var bufferControl3 bytes.Buffer
												binary.Write(&bufferControl3, binary.BigEndian, &archivo)
												_, errs3 := file.Write(bufferControl3.Bytes())
												if errs3 != nil {
													fmt.Println("ERROR WE")
												}
												fmt.Println("Usuario Creado con Exito")
												fmt.Println("------------------users.txt------------------")
												fmt.Println(ReEscritura)
											} else {
												fmt.Println("ERROR no se puede escribir ya que sobrepasa el tamano pre establecido")
											}
										} else {
											fmt.Println("ERROR ya existe un usuario con ese nombre")
										}
									} else {
										fmt.Println("ERROR ya existe un grupo con ese nombre")
									}
								} else {
									fmt.Println("ERROR solo el usuario root puede ejecutar este comando")
								}
							} else {
								fmt.Println("ERROR necesita iniciar sesion para poder ejecutar este comando")
							}
						} else {
							fmt.Println("ERROR grupo del usuario excede de los 10 caracteres permitidos")
						}
					} else {
						fmt.Println("ERROR contrasena de usuario excede de los 10 caracteres permitidos")
					}
				} else {
					fmt.Println("ERROR nombre de usuario excede de los 10 caracteres permitidos")
				}
			} else {
				fmt.Println("ERROR nombre de usuario excede de los 10 caracteres permitidos")
			}
		} else {
			fmt.Println("ERROR parametro -pwd no definido")
		}
	} else {
		fmt.Println("ERROR parametro -usuario no definido ")
	}
}

func RMkusr(name string) {
	var ReEscritura string = ""
	if BanderaLogin {
		if CurrentSession.Id_user == 1 && CurrentSession.Id_grp == 1 {
			grupo := BuscarExisteGrupo(name)
			if grupo != -1 {
				usuarios := BuscarExisteUsuario(name)
				if !usuarios {
					file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
					defer file.Close()
					var cadena string = ""
					super := structs.SuperBloque{}
					inodo := structs.InodoTable{}

					file.Seek(CurrentSession.InicioSuper, 0)
					TempSB := structs.SuperBloque{}
					var sizetemp int = int(unsafe.Sizeof(TempSB))
					data := leerBytes(file, sizetemp)
					buffer := bytes.NewBuffer(data)
					Errores := binary.Read(buffer, binary.BigEndian, &super)
					if Errores != nil {
						fmt.Println("F perro")
					}

					file.Seek(CurrentSession.InicioSuper, 0)
					TempTI := structs.SuperBloque{}
					var sizetemp1 int = int(unsafe.Sizeof(TempTI))
					data1 := leerBytes(file, sizetemp1)
					buffer1 := bytes.NewBuffer(data1)
					Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
					if Errores1 != nil {
						fmt.Println("F perro")
					}

					for i := 0; i < 15; i++ {
						if inodo.I_block[i] != -1 {
							archivo := structs.BloqueArchivo{}
							file.Seek(int64(super.S_block_start), 0)
							for j := 0; j <= int(inodo.I_block[i]); j++ {
								temparchivo := structs.BloqueArchivo{}
								var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
								data2 := leerBytes(file, sizetemp2)
								buffer2 := bytes.NewBuffer(data2)
								Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
								if Errores2 != nil {
									fmt.Println("F perro")
								}
							}
							cadena += string(archivo.B_content[:])
						}
					}
					token := strings.Split(cadena, "\n")
					for i := 0; i < len(token); i++ {
						usuario := ""
						token2 := strings.Split(token[i], ",")
						for i := 0; i < len(token2); i++ {
							if token2[1] == "U" {
								usuario = token2[3]
								if strings.Compare(usuario, name) == 0 {
									NuevVar := "0" + ",U," + token2[2] + "," + token2[3] + "," + token2[4] + "\n"
									ReEscritura += NuevVar

								} else {
									str2 := strings.Join(token2, " ")
									ReEscritura += str2 + "\n"
								}
							} else {
								str2 := strings.Join(token2, " ")
								ReEscritura += str2 + "\n"
							}
						}
					}
					archivo := structs.BloqueArchivo{}
					copy(archivo.B_content[:], ReEscritura)
					sizesBloqueCarpeta := uint64(unsafe.Sizeof(structs.BloqueCarpeta{}))
					file.Seek(int64(super.S_inode_start+sizesBloqueCarpeta), 0)
					var bufferControl3 bytes.Buffer
					binary.Write(&bufferControl3, binary.BigEndian, &archivo)
					_, errs3 := file.Write(bufferControl3.Bytes())
					if errs3 != nil {
						fmt.Println("ERROR WE")
					}
					fmt.Println("Usuario Creado con Exito")
					fmt.Println("------------------users.txt------------------")
					fmt.Println(ReEscritura)
				} else {
					fmt.Println("ERROR ya existe un usuario con ese nombre")
				}
			} else {
				fmt.Println("ERROR ya existe un grupo con ese nombre")
			}
		} else {
			fmt.Println("ERROR solo el usuario root puede ejecutar este comando")
		}
	} else {
		fmt.Println("ERROR necesita iniciar sesion para poder ejecutar este comando")
	}
}

func Mkgrp(name string) {
	var ReEscritura string = ""
	if BanderaLogin {
		if CurrentSession.Id_user == 1 && CurrentSession.Id_grp == 1 {
			grupo := BuscarExisteGrupo(name)
			if grupo == -1 {
				idGro := GetId_Grp()
				file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
				defer file.Close()
				var cadena string = ""
				super := structs.SuperBloque{}
				inodo := structs.InodoTable{}

				file.Seek(CurrentSession.InicioSuper, 0)
				TempSB := structs.SuperBloque{}
				var sizetemp int = int(unsafe.Sizeof(TempSB))
				data := leerBytes(file, sizetemp)
				buffer := bytes.NewBuffer(data)
				Errores := binary.Read(buffer, binary.BigEndian, &super)
				if Errores != nil {
					fmt.Println("F perro")
				}

				file.Seek(CurrentSession.InicioSuper, 0)
				TempTI := structs.SuperBloque{}
				var sizetemp1 int = int(unsafe.Sizeof(TempTI))
				data1 := leerBytes(file, sizetemp1)
				buffer1 := bytes.NewBuffer(data1)
				Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
				if Errores1 != nil {
					fmt.Println("F perro")
				}

				for i := 0; i < 15; i++ {
					if inodo.I_block[i] != -1 {
						archivo := structs.BloqueArchivo{}
						file.Seek(int64(super.S_block_start), 0)
						for j := 0; j <= int(inodo.I_block[i]); j++ {
							temparchivo := structs.BloqueArchivo{}
							var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
							data2 := leerBytes(file, sizetemp2)
							buffer2 := bytes.NewBuffer(data2)
							Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
							if Errores2 != nil {
								fmt.Println("F perro")
							}
						}
						cadena += string(archivo.B_content[:])
					}
				}
				nuevoId := strconv.Itoa(idGro)
				NuevVar := nuevoId + ",G," + name + "\n"
				ReEscritura = cadena + NuevVar
				libre := 127 - len(cadena)
				if len(NuevVar) <= libre {
					archivo := structs.BloqueArchivo{}
					copy(archivo.B_content[:], ReEscritura)
					sizesBloqueCarpeta := uint64(unsafe.Sizeof(structs.BloqueCarpeta{}))
					file.Seek(int64(super.S_inode_start+sizesBloqueCarpeta), 0)
					var bufferControl3 bytes.Buffer
					binary.Write(&bufferControl3, binary.BigEndian, &archivo)
					_, errs3 := file.Write(bufferControl3.Bytes())
					if errs3 != nil {
						fmt.Println("ERROR WE")
					}
					fmt.Println("Grupo Creado con Exito")
					fmt.Println("------------------users.txt------------------")
					fmt.Println(ReEscritura)
				} else {
					fmt.Println("ERROR no se puede escribir ya que sobrepasa el tamano pre establecido")
				}

			} else {
				fmt.Println("ERROR ya existe un grupo con ese nombre")
			}
		} else {
			fmt.Println("ERROR solo el usuario root puede ejecutar este comando")
		}
	} else {
		fmt.Println("ERROR necesita iniciar sesion para poder ejecutar este comando")
	}
}

func RMkgrp(name string) {
	var ReEscritura string = ""
	if BanderaLogin {
		if CurrentSession.Id_user == 1 && CurrentSession.Id_grp == 1 {
			grupo := BuscarExisteGrupo(name)
			if grupo == -1 {
				file, _ := os.OpenFile(string(CurrentSession.Direccion[:]), os.O_RDWR, 0777)
				defer file.Close()
				var cadena string = ""
				super := structs.SuperBloque{}
				inodo := structs.InodoTable{}

				file.Seek(CurrentSession.InicioSuper, 0)
				TempSB := structs.SuperBloque{}
				var sizetemp int = int(unsafe.Sizeof(TempSB))
				data := leerBytes(file, sizetemp)
				buffer := bytes.NewBuffer(data)
				Errores := binary.Read(buffer, binary.BigEndian, &super)
				if Errores != nil {
					fmt.Println("F perro")
				}

				file.Seek(CurrentSession.InicioSuper, 0)
				TempTI := structs.SuperBloque{}
				var sizetemp1 int = int(unsafe.Sizeof(TempTI))
				data1 := leerBytes(file, sizetemp1)
				buffer1 := bytes.NewBuffer(data1)
				Errores1 := binary.Read(buffer1, binary.BigEndian, &inodo)
				if Errores1 != nil {
					fmt.Println("F perro")
				}

				for i := 0; i < 15; i++ {
					if inodo.I_block[i] != -1 {
						archivo := structs.BloqueArchivo{}
						file.Seek(int64(super.S_block_start), 0)
						for j := 0; j <= int(inodo.I_block[i]); j++ {
							temparchivo := structs.BloqueArchivo{}
							var sizetemp2 int = int(unsafe.Sizeof(temparchivo))
							data2 := leerBytes(file, sizetemp2)
							buffer2 := bytes.NewBuffer(data2)
							Errores2 := binary.Read(buffer2, binary.BigEndian, &archivo)
							if Errores2 != nil {
								fmt.Println("F perro")
							}
						}
						cadena += string(archivo.B_content[:])
					}
				}
				NuevVar := "0" + ",G," + name + "\n"
				ReEscritura = cadena + NuevVar
				archivo := structs.BloqueArchivo{}
				copy(archivo.B_content[:], ReEscritura)
				sizesBloqueCarpeta := uint64(unsafe.Sizeof(structs.BloqueCarpeta{}))
				file.Seek(int64(super.S_inode_start+sizesBloqueCarpeta), 0)
				var bufferControl3 bytes.Buffer
				binary.Write(&bufferControl3, binary.BigEndian, &archivo)
				_, errs3 := file.Write(bufferControl3.Bytes())
				if errs3 != nil {
					fmt.Println("ERROR WE")
				}
				fmt.Println("Grupo Eliminado con Exito")
				fmt.Println("------------------users.txt------------------")
				fmt.Println(ReEscritura)

			} else {
				fmt.Println("ERROR ya existe un grupo con ese nombre")
			}
		} else {
			fmt.Println("ERROR solo el usuario root puede ejecutar este comando")
		}
	} else {
		fmt.Println("ERROR necesita iniciar sesion para poder ejecutar este comando")
	}
}

func MkFile(path string, r bool, sized string, cont string) {

}

//"1,G,root\n1,U,root,root,123\n"
//para las particiones logicas probar utilizar el seek current o en vez del 0 utilizar el 1 que es lo mismo, aun no he echo las pruebas pero hay que intentarlo
//Revisar los obligatorios y los opcionales en los comandos principales, que creo que solo el login tiene las validaciones bien echas lgp
