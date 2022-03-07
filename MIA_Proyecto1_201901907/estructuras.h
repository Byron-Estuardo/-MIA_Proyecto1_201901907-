#ifndef ESTRUCTURAS_H
#define ESTRUCTURAS_H

#include <time.h>
#include <QString>

typedef struct {
    char part_status;
    char part_type;
    char part_fit;
    int part_start;
    int part_size;
    char part_name[16];
} Partition;

typedef struct{
    int mbr_size;
    time_t mbr_date_created;
    int mbr_disk_signature;
    char mbr_disk_fit;
    Partition mbr_partition[4];
}MBR;

typedef struct{
    char part_status;
    char part_fit;
    int part_start;
    int part_size;
    int part_next;
    char part_name[16];
}EBR;

typedef struct{
    int s_filesystem_type;
    int s_inodes_count;
    int s_blocks_count;
    int s_free_blocks_count;
    int s_free_inodes_count;
    time_t s_mtime;
    time_t s_umtime;
    int s_mnt_count;
    int s_magic;
    int s_inode_size;
    int s_block_size;
    int s_first_ino;
    int s_first_blo;
    int s_bm_inode_start;
    int s_bm_block_start;
    int s_inode_start;
    int s_block_start;
}SuperBloque;

typedef struct{
    int i_uid;
    int i_gid;
    int i_size;
    int i_block[15];
    char i_type;
    int i_perm;
    time_t i_atime;
    time_t i_ctime;
    time_t i_mtime;
}InodoTable;

typedef struct{
    char b_name[12];
    int b_inodo;
}Content;

typedef struct{
    Content b_content[4];
}BloqueCarpeta;

typedef struct{
    char b_content[64];
}BloqueArchivo;

typedef struct{
    int b_pointer[16];
}BloqueApuntadores;

typedef struct{
    char journal_operation_type[10];
    int journal_type;
    char journal_name[100];
    char journal_content[100];
    time_t journal_date;
    int journal_owner;
    int journal_permissions;
}Journal;

typedef struct{
    int id_user;
    int id_grp;
    int inicioSuper;
    int inicioJournal;
    int tipo_sistema;
    QString direccion;
    char fit;
}Sesion;

typedef struct{
    int id_usr;
    int id_grp;
    char username[12];
    char password[12];
    char group[12];
}Usuario;

#endif // ESTRUCTURAS_H
