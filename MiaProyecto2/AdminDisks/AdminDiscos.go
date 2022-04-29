package AdminDisks

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
