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

type NodoMount struct {
	direccion string
	nombre    string
	id        string
	letra     string
	num       int
}

var ListaMount = []NodoMount{}

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
							}
						}
					}
				}
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

/*
for{
	if{

	}
}
*/

/*
func CrearParticionLogicas(direccion string, nombre string, size int, fit string, unit string) {
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
		var Numero_Extendida = -1
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
					Numero_Extendida = i
					break
				}
			}
			if !ExisteParticion(direccion, nombre) {
				if Numero_Extendida != -1 {
					EBR := structs.EBR{}
					cont, _ := strconv.Atoi(string(MBR.Mbr_partition[Numero_Extendida].Part_start[:]))
					file.Seek(int64(cont), 0)
					var size int = int(unsafe.Sizeof(EBR))
					data := leerBytes(file, size)
					buffer := bytes.NewBuffer(data)
					err := binary.Read(buffer, binary.BigEndian, &EBR)
					if err != nil {
						fmt.Println("Error")
					} else {
						if string(EBR.Part_size[:]) == "0" {
							partsizes, _ := strconv.Atoi(string(MBR.Mbr_partition[Numero_Extendida].Part_size[:]))
							if partsizes < size_bytes {
								fmt.Println("ERROR la particion logica a crear excede el espacio disponible de la particion extendida ")
							} else {
								var stringStart string = string(TempMBR.Mbr_partition[extendida].Part_start[:])
								InicioParticion, _ := strconv.Atoi(stringStart)
								var sizeEBR int64 = int64(unsafe.Sizeof(structs.EBR{}))
								var InicioEnviado int64 = int64(InicioParticion)
								copy(EBR.Part_status[:], "0")
								copy(EBR.Part_fit[:], auxFit)
								copy(EBR.Part_start[:], string(int(InicioEnviado - sizeEBR)))
								copy(EBR.Part_status[:], "0")
							}
						}
					}

				}
			}
		}
	}
}
*/
