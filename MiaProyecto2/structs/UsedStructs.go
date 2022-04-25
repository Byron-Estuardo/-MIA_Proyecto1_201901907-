package structs

type Partition struct {
	Part_status [6]byte
	Part_type   [6]byte
	Part_fit    [6]byte
	Part_start  [60]byte
	Part_size   [60]byte
	Part_name   [30]byte
}

type MBR struct {
	Mbr_size           [60]byte
	Mbr_date_created   [16]byte
	Mbr_disk_signature [20]byte
	Mbr_disk_fit       [6]byte
	Mbr_partition      [4]Partition
}

type EBR struct {
	Part_status [6]byte
	Part_fit    [6]byte
	Part_start  [60]byte
	Part_size   [60]byte
	Part_next   [60]byte
	Part_name   [30]byte
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
