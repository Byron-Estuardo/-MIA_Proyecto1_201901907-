package structs

type Partition struct {
	Part_status [1]byte
	Part_type   [1]byte
	Part_fit    [1]byte
	Part_start  [20]byte
	Part_size   [20]byte
	Part_name   [16]byte
}

type MBR struct {
	Mbr_size           [20]byte
	Mbr_date_created   [20]byte
	Mbr_disk_signature [2]byte
	Mbr_disk_fit       [1]byte
	Mbr_partition      [4]Partition
}

type EBR struct {
	Part_status [1]byte
	Part_fit    [1]byte
	Part_start  [20]byte
	Part_size   [20]byte
	Part_next   [10]byte
	Part_name   [16]byte
}

type SuperBloque struct {
	S_filesystem_type   uint64
	S_inodes_count      uint64
	S_blocks_count      uint64
	S_free_blocks_count uint64
	S_free_inodes_count uint64
	S_mtime             [20]byte
	S_umtime            [20]byte
	S_mnt_count         uint64
	S_magic             uint64
	S_inode_size        uint64
	S_block_size        uint64
	S_first_ino         uint64
	S_first_blo         uint64
	S_bm_inode_start    uint64
	S_bm_block_start    uint64
	S_inode_start       uint64
	S_block_start       uint64
}

type InodoTable struct {
	I_uid   uint64
	I_gid   uint64
	I_size  uint64
	I_block [15]uint64
	I_type  []byte
	I_perm  uint64
	I_atime [20]byte
	I_ctime [20]byte
	I_mtime [20]byte
}

type Content struct {
	B_name  [12]byte
	B_inodo uint64
}

type BloqueCarpeta struct {
	B_content [4]Content
}

type BloqueArchivo struct {
	B_content [64]byte
}

type BloqueApuntadores struct {
	B_pointer uint64
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
	Id_usr   uint64
	Id_grp   uint64
	Username [12]byte
	Password [12]byte
	Group    [12]byte
}
