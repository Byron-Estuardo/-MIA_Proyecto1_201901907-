#include <iostream>
#include <stdio.h>
#include <sys/stat.h>
#include <fstream>
#include <qstring.h>
#include <QString>
#include <math.h>
#include <libgen.h>
#include <QFileInfo>
#include <thread>
#include <chrono>
#include "nodo.h"
#include "estructuras.h"
#include "mount.h"
#include "nodo.h"
#include "scanner.h"
#include "parser.h"

extern int yyparse();
extern Nodo *raiz;
ListaMount *lista = new ListaMount();
Sesion currentSession;
bool Bandera_Global = true;
bool Bandera_login = false;
void leerComando(char*);
using namespace std;
enum Choice
{
    MKDISK = 1,
    RMDISK = 2,
    FDISK = 3,
    MOUNT = 4,
    UNMOUNT = 5,
    REP = 6,
    PARAMETRO = 7,
    EXEC = 8,
    SIZE = 9,
    FIT = 10,
    UNIT = 11,
    PATH = 12,
    TYPE = 13,
    DELETE = 14,
    NAME = 15,
    ADD = 16,
    ID = 17,
    MKFS = 18,
    LOGIN = 19,
    LOGOUT = 20,
    MKGRP = 21,
    RMGRP = 22,
    MKUSR = 23,
    RMUSR = 24,
    CHMOD = 25,
    MKFILE = 26,
    CAT = 27,
    REM = 28,
    EDIT = 29,
    REN = 30,
    MKDIR = 31,
    CP = 32,
    MV = 33,
    CHOWN = 34,
    CHGRP = 35,
    PAUSE = 36,
    RECOVERY = 37,
    LOSS = 38,
    FS = 39,
    USER = 40,
    PASSWORD = 41,
    GROUP = 42,
    UGO = 43,
    CONT = 44,
    FILEE = 45,
    P = 46,
    DEST = 47,
    R = 48,
    INODE = 49,
    JOURNALING = 50,
    BLOCK = 51,
    BM_INODE  = 52,
    BM_BLOCK = 53,
    TREE = 54,
    SB = 55,
    RUTA = 56,
};

QString getDirectorio(QString direccion){
    string aux = direccion.toStdString();
    string delimiter = "/";
    size_t pos = 0;
    string res = "";
    while((pos = aux.find(delimiter))!=string::npos){
        res += aux.substr(0,pos)+"/";
        aux.erase(0,pos + delimiter.length());
    }
    return QString::fromStdString(res);
}

QString getExtension(QString direccion){
    string aux = direccion.toStdString();
    string delimiter = ".";
    size_t pos = 0;
    while((pos = aux.find(delimiter))!=string::npos){
        aux.erase(0,pos+delimiter.length());
    }
    return QString::fromStdString(aux);
}

QString getFileName(QString direccion){
    string aux = direccion.toStdString();
    string delimiter = "/";
    size_t pos = 0;
    string res = "";
    while((pos = aux.find(delimiter))!=string::npos){
        aux.erase(0,pos + delimiter.length());
    }
    delimiter = ".";
    pos = aux.find(delimiter);
    res = aux.substr(0,pos);
    return QString::fromStdString(res);
}

//Administracion de discos y particiones

void crearDisco(QString direccion){
    QString aux = getDirectorio(direccion);
    string comando = "sudo mkdir -p \'"+aux.toStdString()+"\'";//pa crear el archivo
    //cout<<"direccion";
    //cout<<aux.toStdString();
    system(comando.c_str());
    string comando2 = "sudo chmod -R 777 \'"+aux.toStdString()+"\'";//y pa darle los permisos al archivo
    //cout<<"permisos direccion";
    //cout<<aux.toStdString();
    system(comando2.c_str());
    string arch = direccion.toStdString();
    FILE *fp = fopen(arch.c_str(),"wb");//aqui abrimos el archivo para ver si se creo bien
    if((fp = fopen(arch.c_str(),"wb")))
        fclose(fp);
    else
        cout << "Error al crear el archivo" << endl;
}

void FEXT2(int inicio, int tamano, QString direccion){
    double n = (tamano - static_cast<int>(sizeof(SuperBloque)))/(4 + static_cast<int>(sizeof(InodoTable)) +3*static_cast<int>(sizeof(BloqueArchivo)));
    int num_estructuras = static_cast<int>(floor(n));
    int num_bloques = 3*num_estructuras;

    SuperBloque super;
    super.s_filesystem_type = 2;
    super.s_inodes_count = num_estructuras;
    super.s_blocks_count = num_bloques;
    super.s_free_blocks_count = num_bloques -2;
    super.s_free_inodes_count = num_estructuras -2;
    super.s_mtime = time(nullptr);
    super.s_umtime = 0;
    super.s_mnt_count = 0;
    super.s_magic = 0xEF53;
    super.s_inode_size = sizeof(InodoTable);
    super.s_block_size = sizeof(BloqueArchivo);
    super.s_first_ino = 2;
    super.s_first_blo = 2;
    super.s_bm_inode_start = inicio + static_cast<int>(sizeof(SuperBloque));
    super.s_bm_block_start = inicio + static_cast<int>(sizeof(SuperBloque)) + num_estructuras;
    super.s_inode_start = inicio + static_cast<int>(sizeof (SuperBloque)) + num_estructuras + num_bloques;
    super.s_block_start = inicio + static_cast<int>(sizeof(SuperBloque)) + num_estructuras + num_bloques + (static_cast<int>(sizeof(InodoTable))*num_estructuras);

    InodoTable inodo;
    BloqueCarpeta bloque;

    char buffer = '0';
    char buffer2 = '1';
    char buffer3 = '2';

    FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");

    fseek(fp,inicio,SEEK_SET);
    fwrite(&super,sizeof(SuperBloque),1,fp);
    for(int i = 0; i < num_estructuras; i++){
        fseek(fp,super.s_bm_inode_start + i,SEEK_SET);
        fwrite(&buffer,sizeof(char),1,fp);
    }
    fseek(fp,super.s_bm_inode_start,SEEK_SET);
    fwrite(&buffer2,sizeof(char),1,fp);
    fwrite(&buffer2,sizeof(char),1,fp);
    for(int i = 0; i < num_bloques; i++){
        fseek(fp,super.s_bm_block_start + i,SEEK_SET);
        fwrite(&buffer,sizeof(char),1,fp);
    }
    fseek(fp,super.s_bm_block_start,SEEK_SET);
    fwrite(&buffer2,sizeof(char),1,fp);
    fwrite(&buffer3,sizeof(char),1,fp);
    inodo.i_uid = 1;
    inodo.i_gid = 1;
    inodo.i_size = 0;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    inodo.i_block[0] = 0;
    for(int i = 1; i < 15;i++)
        inodo.i_block[i] = -1;
    inodo.i_type = '0';
    inodo.i_perm = 664;
    fseek(fp,super.s_inode_start,SEEK_SET);
    fwrite(&inodo,sizeof(InodoTable),1,fp);
    strcpy(bloque.b_content[0].b_name,".");
    bloque.b_content[0].b_inodo=0;

    strcpy(bloque.b_content[1].b_name,"..");
    bloque.b_content[1].b_inodo=0;

    strcpy(bloque.b_content[2].b_name,"users.txt");
    bloque.b_content[2].b_inodo=1;

    strcpy(bloque.b_content[3].b_name,".");
    bloque.b_content[3].b_inodo=-1;
    fseek(fp,super.s_block_start,SEEK_SET);
    fwrite(&bloque,sizeof(BloqueCarpeta),1,fp);
    inodo.i_uid = 1;
    inodo.i_gid = 1;
    inodo.i_size = 27;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    inodo.i_block[0] = 1;
    for(int i = 1; i < 15;i++){
        inodo.i_block[i] = -1;
    }
    inodo.i_type = '1';
    inodo.i_perm = 755;
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
    fwrite(&inodo,sizeof(InodoTable),1,fp);
    BloqueArchivo archivo;
    memset(archivo.b_content,0,sizeof(archivo.b_content));
    strcpy(archivo.b_content,"1,G,root\n1,U,root,root,123\n");
    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)),SEEK_SET);
    fwrite(&archivo,sizeof(BloqueArchivo),1,fp);

    cout << "EXT2" << endl;
    cout << "..." << endl;
    cout << "Disco formateado con exito" << endl;

    fclose(fp);
}

void FEXT3( int inicio, int tamano, QString direccion){
    double n = (tamano - static_cast<int>(sizeof(SuperBloque)))/(4 + static_cast<int>(sizeof(InodoTable)) +3*static_cast<int>(sizeof(BloqueArchivo)));
    int num_estructuras = static_cast<int>(floor(n));
    int num_bloques = 3*num_estructuras;
    int super_size = static_cast<int>(sizeof(SuperBloque));
    int journal_size = static_cast<int>(sizeof(Journal))*num_estructuras;

    SuperBloque super;
    super.s_filesystem_type = 3;
    super.s_inodes_count = num_estructuras;
    super.s_blocks_count = num_bloques;
    super.s_free_blocks_count = num_bloques - 2;
    super.s_free_inodes_count = num_estructuras - 2;
    super.s_mtime = time(nullptr);
    super.s_umtime = 0;
    super.s_mnt_count = 0;
    super.s_magic = 0xEF53;
    super.s_inode_size = sizeof(InodoTable);
    super.s_block_size = sizeof(BloqueArchivo);
    super.s_first_ino = 2;
    super.s_first_blo = 2;
    super.s_bm_inode_start = inicio + super_size + journal_size;
    super.s_bm_block_start = inicio + super_size + journal_size + num_estructuras;
    super.s_inode_start = inicio + super_size + journal_size + num_estructuras + num_bloques;
    super.s_block_start = inicio + super_size + journal_size + num_estructuras + num_bloques + static_cast<int>(sizeof(InodoTable))*num_estructuras;

    InodoTable inodo;
    BloqueCarpeta bloque;

    char buffer = '0';
    char buffer2 = '1';
    char buffer3 = '2';

    FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");

    fseek(fp,inicio,SEEK_SET);
    fwrite(&super,sizeof(SuperBloque),1,fp);
    for(int i = 0; i < num_estructuras; i++){
        fseek(fp,super.s_bm_inode_start + i,SEEK_SET);
        fwrite(&buffer,sizeof(char),1,fp);
    }
    fseek(fp,super.s_bm_inode_start,SEEK_SET);
    fwrite(&buffer2,sizeof(char),1,fp);
    fwrite(&buffer2,sizeof(char),1,fp);
    for(int i = 0; i < num_bloques; i++){
        fseek(fp,super.s_bm_block_start + i,SEEK_SET);
        fwrite(&buffer,sizeof(char),1,fp);
    }
    fseek(fp,super.s_bm_block_start,SEEK_SET);
    fwrite(&buffer2,sizeof(char),1,fp);
    fwrite(&buffer3,sizeof(char),1,fp);
    inodo.i_uid = 1;
    inodo.i_gid = 1;
    inodo.i_size = 0;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    inodo.i_block[0] = 0;
    for(int i = 1; i < 15;i++){
        inodo.i_block[i] = -1;
    }
    inodo.i_type = '0';
    inodo.i_perm = 664;
    fseek(fp,super.s_inode_start,SEEK_SET);
    fwrite(&inodo,sizeof(InodoTable),1,fp);
    strcpy(bloque.b_content[0].b_name,".");
    bloque.b_content[0].b_inodo=0;

    strcpy(bloque.b_content[1].b_name,"..");
    bloque.b_content[1].b_inodo=0;

    strcpy(bloque.b_content[2].b_name,"users.txt");
    bloque.b_content[2].b_inodo=1;

    strcpy(bloque.b_content[3].b_name,".");
    bloque.b_content[3].b_inodo=-1;
    fseek(fp,super.s_block_start,SEEK_SET);
    fwrite(&bloque,sizeof(BloqueCarpeta),1,fp);
    inodo.i_uid = 1;
    inodo.i_gid = 1;
    inodo.i_size = 27;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    inodo.i_block[0] = 1;
    for(int i = 1; i < 15;i++){
        inodo.i_block[i] = -1;
    }
    inodo.i_type = '1';
    inodo.i_perm = 755;
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
    fwrite(&inodo,sizeof(InodoTable),1,fp);
    BloqueArchivo archivo;
    memset(archivo.b_content,0,sizeof(archivo.b_content));
    strcpy(archivo.b_content,"1,G,root\n1,U,root,root,123\n");
    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta)),SEEK_SET);
    fwrite(&archivo,sizeof(BloqueArchivo),1,fp);

    cout << "EXT3" << endl;
    cout << "..." << endl;
    cout << "Disco formateado con exito" << endl;

    fclose(fp);
}

bool existeParticion(QString direccion, QString nombre){
    int extendida = -1;
    FILE *fp;
    if((fp = fopen(direccion.toStdString().c_str(),"rb+"))){
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot, sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(strcmp(masterboot.mbr_partition[i].part_name,nombre.toStdString().c_str()) == 0){
                fclose(fp);
                return true;
            }else if(masterboot.mbr_partition[i].part_type == 'E'){
                extendida = i;
            }
        }
        if(extendida != -1){
            fseek(fp, masterboot.mbr_partition[extendida].part_start,SEEK_SET);
            EBR extendedBoot;
            while((fread(&extendedBoot,sizeof(EBR),1,fp))!=0 && (ftell(fp) < (masterboot.mbr_partition[extendida].part_size + masterboot.mbr_partition[extendida].part_start))){
                if(strcmp(extendedBoot.part_name,nombre.toStdString().c_str()) == 0){
                    fclose(fp);
                    return true;
                }
                if(extendedBoot.part_next == -1){
                    fclose(fp);
                    return false;
                }
            }
        }
    }
    fclose(fp);
    return false;
}

int buscarParticion_Primaria_Extendida(QString direccion, QString nombre){
    string auxPath = direccion.toStdString();
    string auxName = nombre.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(),"rb+"))){
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_status != '1'){
                if(strcmp(masterboot.mbr_partition[i].part_name,auxName.c_str()) == 0){
                    return i;
                }
            }
        }

    }
    return -1;
}

int buscarParticion_Logica(QString direccion, QString nombre){
    string auxPath = direccion.toStdString();
    string auxName = nombre.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(),"rb+"))){
        int extendida = -1;
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_type == 'E'){
                extendida = i;
                break;
            }
        }
        if(extendida != -1){
            EBR extendedBoot;
            fseek(fp, masterboot.mbr_partition[extendida].part_start,SEEK_SET);
            while(fread(&extendedBoot,sizeof(EBR),1,fp)!=0 && (ftell(fp) < masterboot.mbr_partition[extendida].part_start + masterboot.mbr_partition[extendida].part_size)){
                if(strcmp(extendedBoot.part_name, auxName.c_str()) == 0){
                    return (ftell(fp) - sizeof(EBR));
                }
            }
        }
        fclose(fp);
    }
    return -1;
}

void crearParticionPrimaria(QString direccion, QString nombre, int size, char fit, char unit, QString archivo){
    char auxFit = 0;
    char auxUnit = 0;
    string auxPath = direccion.toStdString();
    int size_bytes = 1024;
    char buffer = '1';

    if(fit != 0)
        auxFit = fit;
    else
        auxFit = 'W';

    if(unit != 0){
        auxUnit = unit;
        if(auxUnit == 'b'){
            size_bytes = size;
        }else if(auxUnit == 'k'){
            size_bytes = size * 1024;
        }else{
            size_bytes = size*1024*1024;
        }
    }else{
        size_bytes = size * 1024;
    }

    FILE *fp;
    MBR masterboot;

    if((fp = fopen(auxPath.c_str(), "rb+"))){
        bool flagParticion = false;
        int numParticion = 0;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                flagParticion = true;
                numParticion = i;
                break;
            }
        }

        if(flagParticion){
            int espacioUsado = 0;
            for(int i = 0; i < 4; i++){
                if(masterboot.mbr_partition[i].part_status!='1'){
                    espacioUsado += masterboot.mbr_partition[i].part_size;
                }
            }
            if(archivo == "principal"){
                cout << "Espacio disponible: " << (masterboot.mbr_size - espacioUsado) << " Bytes" << endl;
                cout << "Espacio necesario:  " << size_bytes << " Bytes" << endl;
            }
            if((masterboot.mbr_size - espacioUsado) >= size_bytes){
                if(!existeParticion(direccion,nombre)){
                    if(masterboot.mbr_disk_fit == 'F'){
                        masterboot.mbr_partition[numParticion].part_type = 'P';
                        masterboot.mbr_partition[numParticion].part_fit = auxFit;
                        if(numParticion == 0){
                            masterboot.mbr_partition[numParticion].part_start = sizeof(masterboot);
                        }else{
                            masterboot.mbr_partition[numParticion].part_start = masterboot.mbr_partition[numParticion-1].part_start + masterboot.mbr_partition[numParticion-1].part_size;
                        }
                        masterboot.mbr_partition[numParticion].part_size = size_bytes;
                        masterboot.mbr_partition[numParticion].part_status = '0';
                        strcpy(masterboot.mbr_partition[numParticion].part_name,nombre.toStdString().c_str());
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[numParticion].part_start,SEEK_SET);
                        for(int i = 0; i < size_bytes; i++){
                            fwrite(&buffer,1,1,fp);
                        }
                        if(archivo == "principal") cout << "...\n" << "Particion primaria creada con exito!" <<  endl;
                    }else if(masterboot.mbr_disk_fit == 'B'){
                        int bestIndex = numParticion;
                        for(int i = 0; i < 4; i++){
                            if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                                if(i != numParticion){
                                    if(masterboot.mbr_partition[bestIndex].part_size > masterboot.mbr_partition[i].part_size){
                                        bestIndex = i;
                                        break;
                                    }
                                }
                            }
                        }
                        masterboot.mbr_partition[bestIndex].part_type = 'P';
                        masterboot.mbr_partition[bestIndex].part_fit = auxFit;
                        if(bestIndex == 0){
                            masterboot.mbr_partition[bestIndex].part_start = sizeof(masterboot);
                        }else{
                            masterboot.mbr_partition[bestIndex].part_start = masterboot.mbr_partition[bestIndex-1].part_start + masterboot.mbr_partition[bestIndex-1].part_size;
                        }
                        masterboot.mbr_partition[bestIndex].part_size = size_bytes;
                        masterboot.mbr_partition[bestIndex].part_status = '0';
                        strcpy(masterboot.mbr_partition[bestIndex].part_name,nombre.toStdString().c_str());
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[bestIndex].part_start,SEEK_SET);
                        for(int i = 0; i < size_bytes; i++){
                            fwrite(&buffer,1,1,fp);
                        }
                        if(archivo == "principal") cout << "...\n" << "Particion primaria creada con exito!" <<  endl;
                    }else if(masterboot.mbr_disk_fit == 'W'){
                        int  worstIndex= numParticion;
                        for(int i = 0; i < 4; i++){
                            if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                                if(i != numParticion){
                                    if(masterboot.mbr_partition[worstIndex].part_size < masterboot.mbr_partition[i].part_size){
                                        worstIndex = i;
                                        break;
                                    }
                                }
                            }
                        }
                        masterboot.mbr_partition[worstIndex].part_type = 'P';
                        masterboot.mbr_partition[worstIndex].part_fit = auxFit;
                        if(worstIndex == 0){
                            masterboot.mbr_partition[worstIndex].part_start = sizeof(masterboot);
                        }else{
                            masterboot.mbr_partition[worstIndex].part_start = masterboot.mbr_partition[worstIndex-1].part_start + masterboot.mbr_partition[worstIndex-1].part_size;
                        }
                        masterboot.mbr_partition[worstIndex].part_size = size_bytes;
                        masterboot.mbr_partition[worstIndex].part_status = '0';
                        strcpy(masterboot.mbr_partition[worstIndex].part_name,nombre.toStdString().c_str());
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[worstIndex].part_start,SEEK_SET);
                        for(int i = 0; i < size_bytes; i++){
                            fwrite(&buffer,1,1,fp);
                        }
                        if(archivo == "principal") cout << "...\n" << "Particion primaria creada con exito!" <<  endl;
                    }
                }else{
                    cout << "ERROR ya existe una particion con ese nombre" << endl;
                }

            }else{
                cout << "ERROR la particion a crear sobrepasa el espacio libre disponible" << endl;
            }
        }else{
            cout << "ERROR: Ya existen 4 particiones, no se puede crear otra" << endl;
            cout << "Elimine una de las particiones existentes para poder crear otra" << endl;
        }
    fclose(fp);
    }else{
        cout << "ERROR el disco no existe" << endl;
    }
}

void crearParticionExtendida(QString direccion, QString nombre, int size, char fit, char unit, QString archivo){
    char auxFit = 0;
    char auxUnit = 0;
    string auxPath = direccion.toStdString();
    int size_bytes = 1024;
    char buffer = '1';

    if(fit != 0)
        auxFit = fit;
    else
        auxFit = 'W';

    if(unit != 0){
        auxUnit = unit;
        if(auxUnit == 'b'){
            size_bytes = size;
        }else if(auxUnit == 'k'){
            size_bytes = size * 1024;
        }else{
            size_bytes = size*1024*1024;
        }
    }else{
        size_bytes = size * 1024;
    }

    FILE *fp;
    MBR masterboot;
    if((fp = fopen(auxPath.c_str(), "rb+"))){
        bool flagParticion = false;
        bool flagExtendida = false;
        int numParticion = 0;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if (masterboot.mbr_partition[i].part_type == 'E'){
                flagExtendida = true;
                break;
            }
        }
        if(!flagExtendida){
            for(int i = 0; i < 4; i++){
                if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                    flagParticion = true;
                    numParticion = i;
                    break;
                }
            }
            if(flagParticion){
                int espacioUsado = 0;
                for(int i = 0; i < 4; i++){
                    if(masterboot.mbr_partition[i].part_status!='1'){
                       espacioUsado += masterboot.mbr_partition[i].part_size;
                    }
                }
                if(archivo == "principal"){
                    cout << "Espacio disponible: " << (masterboot.mbr_size - espacioUsado) <<" Bytes"<< endl;
                    cout << "Espacio necesario:  " << size_bytes << " Bytes" << endl;
                }
                if((masterboot.mbr_size - espacioUsado) >= size_bytes){
                    if(!(existeParticion(direccion,nombre))){
                        if(masterboot.mbr_disk_fit == 'F'){
                            masterboot.mbr_partition[numParticion].part_type = 'E';
                            masterboot.mbr_partition[numParticion].part_fit = auxFit;

                            if(numParticion == 0){
                                masterboot.mbr_partition[numParticion].part_start = sizeof(masterboot);
                            }else{
                                masterboot.mbr_partition[numParticion].part_start =  masterboot.mbr_partition[numParticion-1].part_start + masterboot.mbr_partition[numParticion-1].part_size;
                            }
                            masterboot.mbr_partition[numParticion].part_size = size_bytes;
                            masterboot.mbr_partition[numParticion].part_status = '0';
                            strcpy(masterboot.mbr_partition[numParticion].part_name,nombre.toStdString().c_str());
                            fseek(fp,0,SEEK_SET);
                            fwrite(&masterboot,sizeof(MBR),1,fp);
                            fseek(fp, masterboot.mbr_partition[numParticion].part_start,SEEK_SET);
                            EBR extendedBoot;
                            extendedBoot.part_fit = auxFit;
                            extendedBoot.part_status = '0';
                            extendedBoot.part_start = masterboot.mbr_partition[numParticion].part_start;
                            extendedBoot.part_size = 0;
                            extendedBoot.part_next = -1;
                            strcpy(extendedBoot.part_name, "");
                            fwrite(&extendedBoot,sizeof (EBR),1,fp);
                            for(int i = 0; i < (size_bytes - (int)sizeof(EBR)); i++){
                                fwrite(&buffer,1,1,fp);
                            }
                            if(archivo == "principal") cout << "...\n" << "Particion extendida creada con exito!"<< endl;
                        }else if(masterboot.mbr_disk_fit == 'B'){
                            int bestIndex = numParticion;
                            for(int i = 0; i < 4; i++){
                                if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                                    if(i != numParticion){
                                        if(masterboot.mbr_partition[bestIndex].part_size > masterboot.mbr_partition[i].part_size){
                                            bestIndex = i;
                                            break;
                                        }
                                    }
                                }
                            }
                            masterboot.mbr_partition[bestIndex].part_type = 'E';
                            masterboot.mbr_partition[bestIndex].part_fit = auxFit;
                            if(bestIndex == 0){
                                masterboot.mbr_partition[bestIndex].part_start = sizeof(masterboot);
                            }else{
                                masterboot.mbr_partition[bestIndex].part_start =  masterboot.mbr_partition[bestIndex-1].part_start + masterboot.mbr_partition[bestIndex-1].part_size;
                            }
                            masterboot.mbr_partition[bestIndex].part_size = size_bytes;
                            masterboot.mbr_partition[bestIndex].part_status = '0';
                            strcpy(masterboot.mbr_partition[bestIndex].part_name,nombre.toStdString().c_str());
                            fseek(fp,0,SEEK_SET);
                            fwrite(&masterboot,sizeof(MBR),1,fp);
                            fseek(fp, masterboot.mbr_partition[bestIndex].part_start,SEEK_SET);
                            EBR extendedBoot;
                            extendedBoot.part_fit = auxFit;
                            extendedBoot.part_status = '0';
                            extendedBoot.part_start = masterboot.mbr_partition[bestIndex].part_start;
                            extendedBoot.part_size = 0;
                            extendedBoot.part_next = -1;
                            strcpy(extendedBoot.part_name, "");
                            fwrite(&extendedBoot,sizeof (EBR),1,fp);
                            for(int i = 0; i < (size_bytes - (int)sizeof(EBR)); i++){
                                fwrite(&buffer,1,1,fp);
                            }
                            if(archivo == "principal") cout << "...\n" << "Particion extendida creada con exito!"<< endl;
                        }else if(masterboot.mbr_disk_fit == 'W'){
                            int  worstIndex= numParticion;
                            for(int i = 0; i < 4; i++){
                                if(masterboot.mbr_partition[i].part_start == -1 || (masterboot.mbr_partition[i].part_status == '1' && masterboot.mbr_partition[i].part_size>=size_bytes)){
                                    if(i != numParticion){
                                        if(masterboot.mbr_partition[worstIndex].part_size < masterboot.mbr_partition[i].part_size){
                                            worstIndex = i;
                                            break;
                                        }
                                    }
                                }
                            }
                            masterboot.mbr_partition[worstIndex].part_type = 'E';
                            masterboot.mbr_partition[worstIndex].part_fit = auxFit;
                            if(worstIndex == 0){
                                masterboot.mbr_partition[worstIndex].part_start = sizeof(masterboot);
                            }else{
                                masterboot.mbr_partition[worstIndex].part_start =  masterboot.mbr_partition[worstIndex-1].part_start + masterboot.mbr_partition[worstIndex-1].part_size;
                            }
                            masterboot.mbr_partition[worstIndex].part_size = size_bytes;
                            masterboot.mbr_partition[worstIndex].part_status = '0';
                            strcpy(masterboot.mbr_partition[worstIndex].part_name,nombre.toStdString().c_str());
                            fseek(fp,0,SEEK_SET);
                            fwrite(&masterboot,sizeof(MBR),1,fp);
                            fseek(fp, masterboot.mbr_partition[worstIndex].part_start,SEEK_SET);
                            EBR extendedBoot;
                            extendedBoot.part_fit = auxFit;
                            extendedBoot.part_status = '0';
                            extendedBoot.part_start = masterboot.mbr_partition[worstIndex].part_start;
                            extendedBoot.part_size = 0;
                            extendedBoot.part_next = -1;
                            strcpy(extendedBoot.part_name, "");
                            fwrite(&extendedBoot,sizeof (EBR),1,fp);
                            for(int i = 0; i < (size_bytes - (int)sizeof(EBR)); i++){
                                fwrite(&buffer,1,1,fp);
                            }
                            if(archivo == "principal") cout << "...\n" << "Particion extendida creada con exito!"<< endl;
                        }
                    }else{
                        cout << "ERROR ya existe una particion con ese nombre" << endl;
                    }
                }else{
                    cout << "ERROR la particion a crear sobrepasa el espacio libre disponible" << endl;
                }
            }else{
                cout << "ERROR: Ya existen 4 particiones, no se puede crear otra" << endl;
                cout << "Elimine una de las particiones existentes para poder crear otra" << endl;
            }
        }else{
            cout << "ERROR ya existe una particion extendida en este disco" << endl;
        }
    fclose(fp);
    }else{
        cout << "ERROR el disco no existe" << endl;
    }
}

void crearParticionLogica(QString direccion, QString nombre, int size, char fit, char unit, QString archivo){
    char auxFit = 0;
    char auxUnit = 0;
    string auxPath = direccion.toStdString();
    int size_bytes = 1024;
    char buffer = '1';

    if(fit != 0)
        auxFit = fit;
    else
        auxFit = 'W';
    if(unit != 0){
        auxUnit = unit;
        if(auxUnit == 'b'){
            size_bytes = size;
        }else if(auxUnit == 'k'){
            size_bytes = size * 1024;
        }else{
            size_bytes = size*1024*1024;
        }
    }else{
        size_bytes = size * 1024;
    }

    FILE *fp;
    MBR masterboot;
    if((fp = fopen(auxPath.c_str(), "rb+"))){
        int numExtendida = -1;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_type == 'E'){
                numExtendida = i;
                break;
            }
        }
        if(!existeParticion(direccion,nombre)){
            if(numExtendida != -1){
                EBR extendedBoot;
                int cont = masterboot.mbr_partition[numExtendida].part_start;
                fseek(fp,cont,SEEK_SET);
                fread(&extendedBoot, sizeof(EBR),1,fp);
                if(extendedBoot.part_size == 0){
                    if(masterboot.mbr_partition[numExtendida].part_size < size_bytes){
                        if(archivo == "principal") cout << "ERROR la particion logica a crear excede el espacio disponible de la particion extendida " << endl;
                    }else{
                        extendedBoot.part_status = '0';
                        extendedBoot.part_fit = auxFit;
                        extendedBoot.part_start = ftell(fp) - sizeof(EBR);
                        extendedBoot.part_size = size_bytes;
                        extendedBoot.part_next = -1;
                        strcpy(extendedBoot.part_name, nombre.toStdString().c_str());
                        fseek(fp, masterboot.mbr_partition[numExtendida].part_start ,SEEK_SET);
                        fwrite(&extendedBoot,sizeof(EBR),1,fp);
                        if(archivo == "principal") cout << "...\nParticion logica creada con exito! "<< endl;
                    }
                }else{
                    while((extendedBoot.part_next != -1) && (ftell(fp) < (masterboot.mbr_partition[numExtendida].part_size + masterboot.mbr_partition[numExtendida].part_start))){
                        fseek(fp,extendedBoot.part_next,SEEK_SET);
                        fread(&extendedBoot,sizeof(EBR),1,fp);
                    }
                    int espacioNecesario = extendedBoot.part_start + extendedBoot.part_size + size_bytes;
                    if(espacioNecesario <= (masterboot.mbr_partition[numExtendida].part_size + masterboot.mbr_partition[numExtendida].part_start)){
                        extendedBoot.part_next = extendedBoot.part_start + extendedBoot.part_size;
                        fseek(fp,ftell(fp) - sizeof (EBR),SEEK_SET);
                        fwrite(&extendedBoot, sizeof(EBR),1 ,fp);
                        fseek(fp,extendedBoot.part_start + extendedBoot.part_size, SEEK_SET);
                        extendedBoot.part_status = 0;
                        extendedBoot.part_fit = auxFit;
                        extendedBoot.part_start = ftell(fp);
                        extendedBoot.part_size = size_bytes;
                        extendedBoot.part_next = -1;
                        strcpy(extendedBoot.part_name,nombre.toStdString().c_str());
                        fwrite(&extendedBoot,sizeof(EBR),1,fp);
                        if(archivo == "principal") cout << "Particion logica creada con exito! "<< endl;
                    }else{
                        cout << "ERROR la particion logica a crear excede el" << endl;
                        cout << "espacio disponible de la particion extendida" << endl;
                    }
                }
            }else{
                cout << "ERROR se necesita una particion extendida para guardar la particion logica " << endl;
            }
        }else{
            cout << "ERROR ya existe una particion con ese nombre" << endl;
        }

    fclose(fp);
    }else{
        cout << "ERROR el disco no existe" << endl;
    }

}

void eliminarParticion(QString direccion, QString nombre, QString typeDelete, QString archivo){
    string auxPath = direccion.toStdString();
    string auxNombre = nombre.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(), "rb+"))){
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof (MBR),1,fp);
        int index = -1;
        int index_Extendida = -1;
        bool flagExtendida = false;
        string opcion = "";
        char buffer = '\0';
        for(int i = 0; i < 4; i++){
            if((strcmp(masterboot.mbr_partition[i].part_name, auxNombre.c_str())) == 0){
                index = i;
                if(masterboot.mbr_partition[i].part_type == 'E')
                    flagExtendida = true;
                break;
            }else if(masterboot.mbr_partition[i].part_type == 'E'){
                index_Extendida = i;
            }
        }
        if(archivo == "principal"){
            cout << "¿Seguro que desea eliminar la particion? S/N : " ;
            getline(cin, opcion);
        }else
            opcion = "S";
        if(opcion.compare("S") == 0 || opcion.compare("s") == 0){
            if(index != -1){
                if(!flagExtendida){
                    if(typeDelete == "fast"){
                        masterboot.mbr_partition[index].part_status = '1';
                        strcpy(masterboot.mbr_partition[index].part_name,"");
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        if(archivo == "principal") cout << "Particion primaria eliminada con exito!" << endl;

                    }else{
                        masterboot.mbr_partition[index].part_status = '1';
                        strcpy(masterboot.mbr_partition[index].part_name,"");
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
                        fwrite(&buffer,1,masterboot.mbr_partition[index].part_size,fp);
                        if(archivo == "principal") cout << "Particion primaria eliminada con exito!" << endl;
                    }
                }else{
                    if(typeDelete == "fast"){
                        masterboot.mbr_partition[index].part_status = '1';
                        strcpy(masterboot.mbr_partition[index].part_name,"");
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        if(archivo == "principal") cout << "Particion extendida eliminada con exito!" << endl;
                    }else{
                        masterboot.mbr_partition[index].part_status = '1';
                        strcpy(masterboot.mbr_partition[index].part_name,"");
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
                        fwrite(&buffer,1,masterboot.mbr_partition[index].part_size,fp);
                        if(archivo == "principal") cout << "Particion extendida eliminada con exito!" << endl;
                    }
                }
            }else{
                if(index_Extendida != -1){
                    bool flag = false;
                    EBR extendedBoot;
                    fseek(fp,masterboot.mbr_partition[index_Extendida].part_start, SEEK_SET);
                    fread(&extendedBoot,sizeof(EBR),1,fp);
                    if(extendedBoot.part_size!=0){
                        fseek(fp, masterboot.mbr_partition[index_Extendida].part_start,SEEK_SET);
                        while((fread(&extendedBoot,sizeof(EBR),1,fp))!=0 && (ftell(fp) < (masterboot.mbr_partition[index_Extendida].part_start + masterboot.mbr_partition[index_Extendida].part_size))) {
                            if(strcmp(extendedBoot.part_name,nombre.toStdString().c_str()) == 0 && extendedBoot.part_status != '1'){
                                flag = true;
                                break;
                            }else if(extendedBoot.part_next == -1){
                                break;
                            }
                        }
                    }
                    if(flag){
                        if(typeDelete == "fast"){
                            extendedBoot.part_status = '1';
                            strcpy(extendedBoot.part_name, "");
                            fseek(fp, ftell(fp)-sizeof(EBR),SEEK_SET);
                            fwrite(&extendedBoot,sizeof(EBR),1,fp);
                            if(archivo == "principal")
                            cout << "Particion logica eliminada con exito!" << endl;
                        }else{
                            extendedBoot.part_status = '1';
                            strcpy(extendedBoot.part_name, "");
                            fseek(fp, ftell(fp)-sizeof(EBR),SEEK_SET);
                            fwrite(&extendedBoot,sizeof(EBR),1,fp);
                            fwrite(&buffer,1,extendedBoot.part_size,fp);
                            if(archivo == "principal")
                            cout << "Particion logica eliminada con exito!" << endl;
                        }
                    }else{
                        cout << "ERROR no se encuentra la particion a eliminar" << endl;
                    }
                }else{
                    cout << "ERROR no se encuentra la particion a eliminar" << endl;
                }
            }
        }else if(opcion.compare("N") || opcion.compare("n") == 0){
            cout << "Cancelado con exito!" << endl;;
        }else{
            cout << "Opcion incorrecta" << endl;
        }

        fclose(fp);
    }else{
        cout << "ERROR el disco donde se va a eliminar no existe" << endl;
    }
}

void agregarParticion(QString direccion, QString nombre, int add, char unit, QString archivo){
    string auxPath = direccion.toStdString();
    string auxNombre = nombre.toStdString();
    int size_Bytes = 0;
    QString tipo = "";

    if(add > 0)
        tipo = "add";

    if(tipo != "add")
        add = add*(-1);

    if(unit == 'm')
        size_Bytes = add * 1024 * 1024;
    else if(unit == 'k')
        size_Bytes = add * 1024;
    else
        size_Bytes = add;

    FILE *fp;
    if((fp = fopen(auxPath.c_str(), "rb+"))){
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        int index = -1;
        int index_Extendida = -1;
        bool flagExtendida = false;
        for(int i = 0; i < 4; i++){
            if((strcmp(masterboot.mbr_partition[i].part_name, auxNombre.c_str())) == 0){
                index = i;
                if(masterboot.mbr_partition[i].part_type == 'E')
                    flagExtendida = true;
                break;
            }else if(masterboot.mbr_partition[i].part_type == 'E'){
                index_Extendida = i;
            }
        }
        if(index != -1){
            if(!flagExtendida){
                if(tipo == "add"){
                    if(index!=3){
                        int p1 = masterboot.mbr_partition[index].part_start + masterboot.mbr_partition[index].part_size;
                        int p2 = masterboot.mbr_partition[index+1].part_start;
                        if((p2 - p1) != 0){
                            int fragmentacion = p2-p1;
                            if(fragmentacion >= size_Bytes){
                                masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                fseek(fp,0,SEEK_SET);
                                fwrite(&masterboot,sizeof(MBR),1,fp);
                                if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                            }else{
                                cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                            }
                        }else{
                            if(masterboot.mbr_partition[index + 1].part_status == '1'){
                                if(masterboot.mbr_partition[index + 1].part_size >= size_Bytes){
                                    masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                    masterboot.mbr_partition[index + 1].part_size = (masterboot.mbr_partition[index + 1].part_size - size_Bytes);
                                    masterboot.mbr_partition[index + 1].part_start = masterboot.mbr_partition[index + 1].part_start + size_Bytes;
                                    fseek(fp,0,SEEK_SET);
                                    fwrite(&masterboot,sizeof(MBR),1,fp);
                                    if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                                }else{
                                    cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                                }
                            }
                        }
                    }else{
                        int p = masterboot.mbr_partition[index].part_start + masterboot.mbr_partition[index].part_size;
                        int total = masterboot.mbr_size + (int)sizeof(MBR);
                        if((total-p) != 0){
                            int fragmentacion = total - p;
                            if(fragmentacion >= size_Bytes){
                                masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                fseek(fp,0,SEEK_SET);
                                fwrite(&masterboot,sizeof(MBR),1,fp);
                                if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                            }else{
                                cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                            }
                        }else{
                            cout << "ERROR no es posible agregar espacio a la particion porque no hay espacio disponible a su derecha" << endl;
                        }
                    }
                }else{
                    if(size_Bytes >= masterboot.mbr_partition[index].part_size){
                        cout << "ERROR no es posible quitarle esta cantidad de espacio a la particion porque la borraria" << endl;
                    }else{
                        masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size - size_Bytes;
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        if(archivo == "principal") cout << "Se quito espacio a la particion de manera exitosa" << endl;
                    }
                }
            }else{
                if(tipo == "add"){
                    if(index!=3){
                        int p1 = masterboot.mbr_partition[index].part_start + masterboot.mbr_partition[index].part_size;
                        int p2 = masterboot.mbr_partition[index+1].part_start;
                        if((p2 - p1) != 0){
                            int fragmentacion = p2-p1;
                            if(fragmentacion >= size_Bytes){
                                masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                fseek(fp,0,SEEK_SET);
                                fwrite(&masterboot,sizeof(MBR),1,fp);
                                if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                            }else{
                                cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                            }
                        }else{
                            if(masterboot.mbr_partition[index + 1].part_status == '1'){
                                if(masterboot.mbr_partition[index + 1].part_size >= size_Bytes){
                                    masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                    masterboot.mbr_partition[index + 1].part_size = (masterboot.mbr_partition[index + 1].part_size - size_Bytes);
                                    masterboot.mbr_partition[index + 1].part_start = masterboot.mbr_partition[index + 1].part_start + size_Bytes;
                                    fseek(fp,0,SEEK_SET);
                                    fwrite(&masterboot,sizeof(MBR),1,fp);
                                    if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                                }else{
                                    cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                                }
                            }
                        }
                    }else{
                        int p = masterboot.mbr_partition[index].part_start + masterboot.mbr_partition[index].part_size;
                        int total = masterboot.mbr_size + (int)sizeof(MBR);
                        if((total-p) != 0){
                            int fragmentacion = total - p;
                            if(fragmentacion >= size_Bytes){
                                masterboot.mbr_partition[index].part_size = masterboot.mbr_partition[index].part_size + size_Bytes;
                                fseek(fp,0,SEEK_SET);
                                fwrite(&masterboot,sizeof(MBR),1,fp);
                                if(archivo == "principal") cout << "Se agrego espacio a la particion de manera exitosa" << endl;
                            }else{
                                cout << "ERROR no es posible agregar espacio a la particion porque no hay suficiente espacio disponible a su derecha" << endl;
                            }
                        }else{
                            cout << "ERROR no es posible agregar espacio a la particion porque no hay espacio disponible a su derecha" << endl;
                        }
                    }
                }else{
                    if(size_Bytes >= masterboot.mbr_partition[index_Extendida].part_size){
                        cout << "ERROR no es posible quitarle esta cantidad de espacio a la particion porque la borraria" << endl;
                    }else{
                        EBR extendedBoot;
                        fseek(fp, masterboot.mbr_partition[index_Extendida].part_start,SEEK_SET);
                        fread(&extendedBoot,sizeof(EBR),1,fp);
                        while((extendedBoot.part_next != -1) && (ftell(fp) < (masterboot.mbr_partition[index_Extendida].part_size + masterboot.mbr_partition[index_Extendida].part_start))){
                            fseek(fp,extendedBoot.part_next,SEEK_SET);
                            fread(&extendedBoot,sizeof(EBR),1,fp);
                        }
                        int ultimaLogica = extendedBoot.part_start+extendedBoot.part_size;
                        int aux = (masterboot.mbr_partition[index_Extendida].part_start + masterboot.mbr_partition[index_Extendida].part_size) - size_Bytes;
                        if(aux > ultimaLogica){
                            masterboot.mbr_partition[index_Extendida].part_size = masterboot.mbr_partition[index_Extendida].part_size - size_Bytes;
                            fseek(fp,0,SEEK_SET);
                            fwrite(&masterboot,sizeof(MBR),1,fp);
                            if(archivo == "principal") cout << "Se quito espacio a la particion de manera exitosa" << endl;
                        }else{
                            cout << "ERROR si quita ese espacio danaria una logica" << endl;
                        }
                    }
                }
            }
        }else{
            if(index_Extendida != -1){
                int logica = buscarParticion_Logica(direccion, nombre);
                if(logica != -1){
                    if(tipo == "add"){
                        EBR extendedBoot;
                        fseek(fp,logica,SEEK_SET);
                        fread(&extendedBoot,sizeof(EBR),1,fp);

                    }else{
                        EBR extendedBoot;
                        fseek(fp,logica,SEEK_SET);
                        fread(&extendedBoot,sizeof(EBR),1,fp);
                        if(size_Bytes >= extendedBoot.part_size){
                            cout << "ERROR si quita ese espacio eliminaria la logica" << endl;
                        }else{
                            extendedBoot.part_size = extendedBoot.part_size - size_Bytes;
                            fseek(fp,logica,SEEK_SET);
                            fwrite(&extendedBoot,sizeof(EBR),1,fp);
                            if(archivo == "principal") cout << "Se quito espacio a la particion de manera exitosa" << endl;
                        }
                    }
                }else{
                    cout << "ERROR no se encuentra la particion" << endl;
                }
            }else{
                cout << "ERROR no se encuentra la particion a redimensionar" << endl;
            }
        }
        fclose(fp);
    }else{
        cout << "ERROR el disco donde se desea agregar/quitar unidades no existe" << endl;
    }
}

//este es de prueba

int tamano_Particion_Logica(QString direccion, QString nombre){
    string auxPath = direccion.toStdString();
    string auxName = nombre.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(),"rb+"))){
        int extendida = -1;
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_type == 'E'){
                extendida = i;
                break;
            }
        }
        if(extendida != -1){
            EBR extendedBoot;
            fseek(fp, masterboot.mbr_partition[extendida].part_start,SEEK_SET);
            while(fread(&extendedBoot,sizeof(EBR),1,fp)!=0 && (ftell(fp) < masterboot.mbr_partition[extendida].part_start + masterboot.mbr_partition[extendida].part_size)){
                if(strcmp(extendedBoot.part_name, auxName.c_str()) == 0){
                    cout<<"Pruebas"<<endl;
                    cout << extendedBoot.part_start <<endl;
                    cout << (ftell(fp) - sizeof(EBR)) <<endl;
                    cout<<"Pruebas nombres"<<endl;
                    cout << extendedBoot.part_name <<endl;
                    cout << auxName.c_str() <<endl;
                    cout << "Pruebas tamano" <<endl;
                    cout << extendedBoot.part_size <<endl;
                    return (ftell(fp) - sizeof(EBR));
                }
            }
        }
        fclose(fp);
    }
    return -1;
}

//Reportes

//Administracion de discos

void RMKFS(Nodo *raiz){
    bool BanderaID = false;
    bool BanderaType = false;
    bool BanderaFS = false;
    bool Bandera = false;
    QString id = "";
    QString type = "";
    int fs = 2;

    for(int i = 0; i < raiz->hijos.count(); i++)
    {
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case ID:
        {
            if(BanderaID){
                cout << "ERROR parametro -id ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaID = true;
            id = n.valor;
        }
            break;
        case TYPE:
        {
            if(BanderaType){
                cout << "ERROR parametro -type ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaType = true;
            type = n.valor;
        }
            break;
        case FS:
        {
            if(BanderaFS){
                cout << "ERROR parametro -fs ya definido " << endl;
                Bandera = true;
                break;
            }
            BanderaFS = true;
            if(n.valor == "3fs")
                fs = 3;
            else
                fs = 2;
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaID){
            NodoListaMount *aux = lista->getNodo(id);
            if(aux!=nullptr){
                int index = buscarParticion_Primaria_Extendida(aux->direccion,aux->nombre);
                if(index != -1){
                    MBR masterboot;
                    FILE *fp = fopen(aux->direccion.toStdString().c_str(),"rb+");
                    fread(&masterboot,sizeof(MBR),1,fp);
                    int inicio = masterboot.mbr_partition[index].part_start;
                    int tamano = masterboot.mbr_partition[index].part_size;
                    if(fs == 3){
                        cout<< "a prro si jala 3" <<endl;
                        FEXT3(inicio,tamano,aux->direccion);
                    }else{
                        cout<< "a prro si jala 2" <<endl;
                        FEXT2(inicio,tamano,aux->direccion);
                    }
                    fclose(fp);
                }else{
                    index = buscarParticion_Logica(aux->direccion,aux->nombre);
                    tamano_Particion_Logica(aux->direccion,aux->nombre);
                    if(index != -1){
                        EBR extendedBoot;
                        FILE *fp = fopen(aux->direccion.toStdString().c_str(),"rb+");
                        fseek(fp, index, SEEK_SET);
                        fread(&extendedBoot, sizeof (EBR),1,fp);
                        int inicio = extendedBoot.part_start;
                        int tamano = extendedBoot.part_size;
                        if(fs == 3){
                            cout<< "a prro si jala 3" <<endl;
                            FEXT3(inicio,tamano,aux->direccion);
                        }else{
                            cout<< "a prro si jala 2" <<endl;
                            FEXT2(inicio,tamano,aux->direccion);
                        }
                        fclose(fp);
                    }


                }
            }else
                cout << "ERROR no se encuentra ninguna particion montada con ese id" << endl;
        }else
            cout << "ERROR parametro -id no definido" << endl;
    }
}

void RUNMOUNT(Nodo *raiz){
    QString valID = raiz->hijos.at(0).valor;
    int eliminado = lista->eliminarNodo(valID);
    if(eliminado == 1)
        cout << "Unidad desmontada con exito!" << endl;
    else
        cout << "ERROR no se encuentra esa unidad montada, porfavor revise el -id de la particion montada" << endl;
}

void RMOUNT(Nodo *raiz){
    bool BanderaPath = false;
    bool BanderaName = false;
    bool Bandera = false;
    QString valPath = "";
    QString valName = "";
    for(int i = 0; i < raiz->hijos.count(); i++)
    {
        Nodo n = raiz->hijos.at(i);
        switch(n.tipo_)
        {
        case PATH:
        {
            if(BanderaPath){
                Bandera = true;
                cout << "ERROR: Parametro -path ya definido" << endl;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case NAME:
        {
            if(BanderaName){
                Bandera = true;
                cout << "ERROR: Parametro -name ya definido" << endl;
                break;
            }
            BanderaName = true;
            valName = n.valor;
            valName = valName.replace("\"","");
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            if(BanderaName){
                int indexP = buscarParticion_Primaria_Extendida(valPath,valName);
                if(indexP != -1){
                    FILE *fp;
                    if((fp = fopen(valPath.toStdString().c_str(),"rb+"))){
                        MBR masterboot;
                        fseek(fp, 0, SEEK_SET);
                        fread(&masterboot, sizeof(MBR),1,fp);
                        masterboot.mbr_partition[indexP].part_status = '2';
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fclose(fp);
                        int num = lista->buscarNumero(valPath,valName);
                        if(num == -1){
                            cout << "ERROR la particion ya se encuentra  montada, no se puede montar de nuevo" << endl;
                        }else{
                            int letra = lista->buscarLetra(valPath, valName);
                            char auxLetra = static_cast<char>(letra);
                            string id = "07";
                            id += to_string(num) + auxLetra;
                            NodoListaMount *n = new NodoListaMount(valPath,valName,auxLetra,num);
                            lista->insertarNodo(n);
                            cout << "Particion montada con exito!" << endl;
                            cout << endl;
                            cout << "*--------- Particiones Montadas ---------*" <<endl;
                            cout << "*               Name | ID                *" <<endl;
                            cout << "*----------------------------------------*" <<endl;
                            lista->mostrarLista();
                        }
                    }else{
                        cout << "ERROR no se encuentra el disco" << endl;
                    }
                }else{
                    int indexP = buscarParticion_Logica(valPath,valName);
                    if(indexP != -1){
                        FILE *fp;
                        if((fp = fopen(valPath.toStdString().c_str(), "rb+"))){
                            EBR extendedBoot;
                            fseek(fp, indexP, SEEK_SET);
                            fread(&extendedBoot, sizeof(EBR),1,fp);
                            extendedBoot.part_status = '2';
                            fseek(fp,indexP,SEEK_SET);
                            fwrite(&extendedBoot,sizeof(EBR),1, fp);
                            fclose(fp);

                            int num = lista->buscarNumero(valPath,valName);
                            if(num == -1){
                                cout << "ERROR la particion ya se encuentra  montada, no se puede montar de nuevo" << endl;
                            }else{
                                int letra = lista->buscarLetra(valPath, valName);
                                char auxLetra = static_cast<char>(letra);
                                string id = "07";
                                id += to_string(num) + auxLetra;
                                NodoListaMount *n = new NodoListaMount(valPath, valName, auxLetra, num);
                                lista->insertarNodo(n);
                                cout << "Particion montada con exito!" << endl;
                                cout << endl;
                                cout << "*--------- Particiones Montadas ---------*" <<endl;
                                cout << "*               Name | ID                *" <<endl;
                                cout << "*----------------------------------------*" <<endl;
                                lista->mostrarLista();
                            }
                        }else{
                            cout << "ERROR no se encuentra el disco" << endl;
                        }
                    }else{
                        cout << "ERROR no se encuentra la particion que se desea montar" << endl;
                    }
                }
            }else{
                cout << "ERROR parametro -name no definido o encontrado" << endl;
            }
        }else{
            cout << "ERROR parametro -path no definido o encontrado" << endl;
        }
    }

}

void RFDISK(Nodo *raiz)
{
    bool BanderaSize = false;
    bool BanderaUnit = false;
    bool BanderaPath = false;
    bool BanderaType = false;
    bool BanderaFit = false;
    bool BanderaDelete = false;
    bool BanderaName = false;
    bool BanderaAdd = false;
    bool Bandera = false;
    int valSize = 0;
    int valAdd = 0;
    char valUnit = 0;
    char valType = 0;
    char valFit = 0;
    QString valPath = "";
    QString valName = "";
    QString valDelete = "";

    for(int i = 0; i < raiz->hijos.count(); i++)
    {
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_)
        {
        case SIZE:
        {
            if(BanderaSize){
                cout << "<< ERROR: Parametro -size ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaSize = true;
            valSize = n.valor.toInt();
            if(!(valSize > 0)){
                cout << "<< ERROR: parametro -size menor a cero" << endl;
                Bandera = true;
                break;
            }
        }
            break;
        case UNIT:
        {
            if(BanderaUnit){
                cout << "<< ERROR: Parametro -unit ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaUnit = true;
            string temp = n.valor.toStdString();
            valUnit = temp[0];
            if(valUnit == 'B' || valUnit == 'b'){
                valUnit = 'b';
            }else if(valUnit == 'K' || valUnit == 'k'){
                valUnit = 'k';
            }else if(valUnit == 'M' || valUnit == 'm'){
                valUnit = 'm';
            }else{
                cout << "<< ERROR: Valor del parametro unit no reconocido" << endl;
                Bandera = true;
                break;
            }

        }
            break;
        case PATH:
        {
            if(BanderaPath){
                cout << "<< ERROR: Parametro -path ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case TYPE:
        {
            if(BanderaType){
                cout << "<< ERROR: Parametro -type ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaType = true;
            string temp = n.valor.toStdString();
            valType = temp[0];
            if(valType == 'P' || valType == 'p'){
                valType = 'P';
            }else if(valType == 'E' || valType == 'e'){
                valType = 'E';
            }else if(valType == 'L' || valType == 'l'){
                valType = 'L';
            }else{
                cout << "<< ERROR: Valor del parametro -type no reconocido" << endl;
                Bandera = true;
                break;
            }
        }
            break;
        case FIT:
        {
            if(BanderaFit)
            {
                cout << "<< ERROR: Parametro -fit ya definido" << endl;
                Bandera = true;
            }
            BanderaFit = true;
            valFit = n.hijos.at(0).valor.toStdString()[0];
            if(valFit == 'b'){
                valFit = 'B';
            }else if(valFit == 'f'){
                valFit = 'F';
            }else if(valFit == 'w'){
                valFit = 'W';
            }
        }
            break;
        case DELETE:
        {
            if(BanderaDelete){
                cout << "<< ERROR: Parametro -delete ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaDelete = true;
            valDelete = n.valor;
        }
            break;
        case NAME:
        {
            if(BanderaName){
                cout << "<< ERROR: Parametro -name ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaName = true;
            valName = n.valor;
            valName = valName.replace("\"", "");
        }
            break;
        case ADD:
        {
            if(BanderaAdd){
                cout << "<< ERROR: Parametro -add ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaAdd = true;
            valAdd = n.valor.toInt();
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            if(BanderaName){
                if(BanderaSize){
                    if(BanderaDelete || BanderaAdd){
                        cout << "<< ERROR: Parametro -delete|-add demas" << endl;
                    }else{
                        if(BanderaType){
                            if(valType == 'P'){
                                crearParticionPrimaria(valPath, valName, valSize, valFit, valUnit,"principal");
                            }else if(valType == 'E'){
                                crearParticionExtendida(valPath, valName, valSize, valFit, valUnit,"principal");
                            }else if(valType == 'L'){
                                crearParticionLogica(valPath, valName, valSize, valFit, valUnit,"principal");
                            }
                        }else{
                            crearParticionPrimaria(valPath, valName, valSize, valFit, valUnit,"principal");
                        }
                    }
                }else if(BanderaAdd){
                    if(BanderaSize || BanderaDelete){
                        cout << "<< ERROR: Parametros -size|-delete demas" << endl;
                    }else{
                        if(BanderaUnit){
                            agregarParticion(valPath, valName, valAdd, valUnit,"principal");
                        }else{
                            cout << "<< ERROR parametro -unit no definido "<< endl;
                        }
                    }
                }else if(BanderaDelete){
                    if(BanderaSize || BanderaAdd || BanderaFit || BanderaType){
                        cout << "<< ERROR: Parametros demas" << endl;
                    }else{
                        eliminarParticion(valPath, valName, valDelete,"principal");
                    }
                }
            }else {
                cout << "<< ERROR parametro -name no definido" << endl;
            }
        }else{
            cout << "<< ERROR parametro -path no definido" << endl;
        }
    }
}

void RRMDISK(Nodo *raiz)
{
    QString valPath = raiz->hijos.at(0).valor;
    valPath = valPath.replace("\"","");
    FILE *fp;
    if((fp=fopen(valPath.toStdString().c_str(),"r"))){
        string opcion = "";
        cout << ">> Esta seguro que desea eliminar el disco? S/N : ";
        getline(cin,opcion);
        if(opcion.compare("S") == 0 || opcion.compare("s") == 0){
            string comando = "rm \""+valPath.toStdString()+"\"";
            system(comando.c_str());
            cout <<"Disco eliminado con exito!" << endl;
        }else if(opcion.compare("N") || opcion.compare("n") == 0){
            cout << "Cancelado con exito!" << endl;;
        }else{
            cout << "Opcion incorrecta" << endl;
        }
        fclose(fp);
    }else{
        cout << "No existe el disco a eliminar" << endl;
    }
}

void RMKDISK(Nodo *raiz){
    bool banderaSize = false;
    bool banderaFit = false;
    bool banderaUnit = false;
    bool banderaPath = false;
    bool banderaRepetidos = false;
    int valSize = 0;
    char valFit = 0;
    char valUnit = 0;
    QString valPath = "";

    for(int i = 0; i < raiz->hijos.count(); i++)
    {
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case SIZE:
        {
            if(banderaSize){
                cout << "ERROR: Parametro -size ya definido" << endl;
                banderaRepetidos = true;
                break;
            }
            banderaSize = true;
            valSize = n.valor.toInt();
            if(!(valSize > 0)){
                banderaRepetidos = true;
                cout << "ERROR: Parametro -size menor a cero" << endl;
                break;
            }
        }
            break;
        case FIT:
        {
            if(banderaFit){
                cout << "ERROR: Parametro -fit ya definido" << endl;
                banderaRepetidos = true;
                break;
            }
            banderaFit = true;
            valFit = n.hijos.at(0).valor.toStdString()[0];
            if(valFit == 'b'){
                valFit = 'B';
            }else if(valFit == 'f'){
                valFit = 'F';
            }else if(valFit == 'w'){
                valFit = 'W';
            }
        }
            break;
        case UNIT:
        {
            if(banderaUnit){
                cout << "ERROR: Parametro -unit ya definido" << endl;
                banderaRepetidos = true;
                break;
            }
            banderaUnit = true;
            string temp = n.valor.toStdString();
            valUnit = temp[0];
            if(valUnit == 'k'|| valUnit == 'K'){
                valUnit = 'k';
            }else if(valUnit == 'm' || valUnit == 'M'){
                valUnit = 'm';
            }else{
                cout << "ERROR: Valor del parametro -unit no reconocido" << endl;
                banderaRepetidos = true;
                break;
            }
        }
            break;
        case PATH:
        {
            if(banderaPath){
                cout << "ERROR: Parametro -path ya definido" << endl;
                banderaRepetidos = true;
                break;
            }
            banderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        }

    }

    if(!banderaRepetidos){
        if(banderaPath){
            if(banderaSize){
                MBR masterboot;
                int total_size = 1024;
                crearDisco(valPath);

                masterboot.mbr_date_created = time(nullptr);
                masterboot.mbr_disk_signature = static_cast<int>(time(nullptr));

                if(banderaUnit){
                    if(valUnit == 'm'){
                        masterboot.mbr_size = valSize*(1024*1024);
                        total_size = valSize * 1024;
                    }else{
                        masterboot.mbr_size = valSize * 1024;
                        total_size = valSize;
                    }
                }else{
                    masterboot.mbr_size = valSize*1048576;
                    total_size = valSize * 1024;
                }

                if(banderaFit)
                    masterboot.mbr_disk_fit = valFit;
                else
                    masterboot.mbr_disk_fit = 'F';

                for(int i = 0; i < 4; i++){
                    masterboot.mbr_partition[i].part_status = '0';
                    masterboot.mbr_partition[i].part_type = '0';
                    masterboot.mbr_partition[i].part_fit = '0';
                    masterboot.mbr_partition[i].part_size = 0;
                    masterboot.mbr_partition[i].part_start = -1;
                    strcpy(masterboot.mbr_partition[i].part_name,"");
                }
                string comando = "dd if=/dev/zero of=\""+valPath.toStdString()+"\" bs=1024 count="+to_string(total_size);//esto es para darle el tamano al archivo
                system(comando.c_str());
                FILE *fp = fopen(valPath.toStdString().c_str(),"rb+");
                fseek(fp,0,SEEK_SET);
                fwrite(&masterboot,sizeof(MBR),1,fp);
                fclose(fp);

                cout << "Disco creado con exito!" << endl;
            }else{
                cout << "ERROR Parametro -size no definido " << endl;
            }
        }else{
            cout << "<< ERROR Parametro -path  no definido" << endl;
        }
    }
}

//Administracion de Usuarios y Grupos

void LOG_OUT(){
    if(Bandera_login){
        Bandera_login = false;
        currentSession.id_user = -1;
        currentSession.direccion = "";
        currentSession.inicioSuper = -1;
        cout << "...\nEl usuario se deslogueo, ya no se encuentra la session activa " << endl;
    }else
        cout << "ERROR no hay ninguna sesion activa" << endl;
}

InodoTable crearInodo(int size,char type,int perm){
    InodoTable inodo ;
    inodo.i_uid = currentSession.id_user;
    inodo.i_gid = currentSession.id_grp;
    inodo.i_size = size;
    inodo.i_atime = time(nullptr);
    inodo.i_ctime = time(nullptr);
    inodo.i_mtime = time(nullptr);
    for(int i = 0; i < 15; i++)
        inodo.i_block[i] = -1;
    inodo.i_type = type;
    inodo.i_perm = perm;
    return inodo;
}

BloqueCarpeta crearBloqueCarpeta(){
    BloqueCarpeta carpeta;
    for(int i = 0; i < 4; i++){
        strcpy(carpeta.b_content[i].b_name,"");
        carpeta.b_content[i].b_inodo = -1;
    }
    return carpeta;
}

bool permisosDeEscritura(int permisos, bool BanderaUser, bool BanderaGroup){
    string aux = to_string(permisos);
    char propietario = aux[0];
    char grupo = aux[1];
    char otros = aux[2];
    if((propietario == '2' || propietario == '3' || propietario == '6' || propietario || '7') && BanderaUser)
        return true;
    else if((grupo == '2' || grupo == '3' || grupo == '6' || grupo == '7') && BanderaGroup)
        return true;
    else if(otros == '2' || otros == '3' || otros == '6' || otros == '7')
        return true;

    return false;
}

bool permisosDeLectura(int permisos, bool BanderaUser, bool BanderaGroup){
    string aux = to_string(permisos);
    int propietario = aux[0] - '0';
    int grupo = aux[1] - '0';
    int otros = aux[2] - '0';
    if((propietario >= 3) && BanderaUser)
        return true;
    else if((grupo >= 3) && BanderaGroup)
        return true;
    else if(otros >= 3)
        return true;
    return false;
}

bool permisosLecturaRecursivo(FILE* stream, int n){
    SuperBloque super;
    InodoTable inodo;
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*n,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);
    bool permisos = permisosDeLectura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
    if(permisos){
        bool response = true;
        if(inodo.i_type == '0'){
            for (int i = 0; i < 12; i++) {
                if(inodo.i_block[i] != -1){
                    char byte = '0';
                    fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
                    byte = static_cast<char>(fgetc(stream));
                    if(byte == '1'){
                        BloqueCarpeta carpeta;
                        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                        fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                        for (int j = 0; j < 4; j++) {
                            if(carpeta.b_content[j].b_inodo != -1){
                                if(strcmp(carpeta.b_content[j].b_name,".")!=0 && strcmp(carpeta.b_content[j].b_name,"..")!=0)
                                    response = permisosLecturaRecursivo(stream,carpeta.b_content[j].b_inodo);
                            }
                        }
                    }else{
                        InodoTable inodoAux;
                        fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*inodo.i_block[i],SEEK_SET);
                        fread(&inodoAux,sizeof(InodoTable),1,stream);
                        response = permisosDeLectura(inodoAux.i_perm,(inodoAux.i_uid == currentSession.id_user),(inodoAux.i_gid == currentSession.id_grp));
                    }
                }
            }
            return response;
        }else
            return true;
    }else
        return false;
}

bool permisosEscrituraRecursivo(FILE* stream, int n){
    SuperBloque super;
    InodoTable inodo;
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*n,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);

    bool permisos = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
    if(permisos){
        bool response = true;
        if(inodo.i_type == '0'){
            for (int i = 0; i < 12; i++) {
                if(inodo.i_block[i] != -1){
                    char byte = '0';
                    fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
                    byte = static_cast<char>(fgetc(stream));
                    if(byte == '1'){
                        BloqueCarpeta carpeta;
                        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                        fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                        for (int j = 0; j < 4; j++) {
                            if(carpeta.b_content[j].b_inodo != -1){
                                if(strcmp(carpeta.b_content[j].b_name,".")!=0 && strcmp(carpeta.b_content[j].b_name,"..")!=0)
                                    response = permisosEscrituraRecursivo(stream,carpeta.b_content[j].b_inodo);
                            }
                        }
                    }else{
                        InodoTable inodoAux;
                        fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*inodo.i_block[i],SEEK_SET);
                        fread(&inodoAux,sizeof(InodoTable),1,stream);
                        response = permisosDeEscritura(inodoAux.i_perm,(inodoAux.i_uid == currentSession.id_user),(inodoAux.i_gid == currentSession.id_grp));
                    }
                }
            }
            return response;
        }else //archivo
            return true;
    }else
        return false;
}

int byteInodoBloque(FILE *stream,int pos, char tipo){
    SuperBloque super;
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    if(tipo == '1'){
        return (super.s_inode_start + static_cast<int>(sizeof(InodoTable))*pos);
    }else if(tipo == '2')
        return (super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*pos);
    return 0;
}

void cambiarPermisosRecursivo(FILE* stream, int n, int permisos){
    SuperBloque super;
    InodoTable inodo;
    BloqueCarpeta carpeta;
    char byte ='0';
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*n,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);
    inodo.i_perm = permisos;
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*n,SEEK_SET);
    fwrite(&inodo,sizeof(InodoTable),1,stream);

    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
            byte = static_cast<char>(fgetc(stream));
            if(byte == '1'){
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                for(int j = 0; j < 4; j++){
                    if(carpeta.b_content[j].b_inodo != -1){
                        if(strcmp(carpeta.b_content[j].b_name,".")!=0 &&  strcmp(carpeta.b_content[j].b_name,"..")!=0)
                            cambiarPermisosRecursivo(stream,carpeta.b_content[j].b_inodo,permisos);
                    }
                }
            }
        }
    }
}

int buscarContentenido(FILE* stream,int numInodo,InodoTable &inodo,BloqueCarpeta &carpeta, BloqueApuntadores &apuntadores,int &content,int &bloque,int &pointer){
    int libre = 0;
    SuperBloque super;
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);
    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            if(i == 12){
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[i],SEEK_SET);
                fread(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                for(int j = 0; j < 16; j++){
                    if(apuntadores.b_pointer[j] != -1){
                        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*apuntadores.b_pointer[j],SEEK_SET);
                        fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                        for(int k = 0; k < 4; k++){
                            if(carpeta.b_content[k].b_inodo == -1){
                                libre = 1;
                                bloque = i;
                                pointer = j;
                                content = k;
                                break;
                            }
                        }
                    }
                    if(libre == 1)
                        break;
                }
            }else if(i == 13){

            }else if(i == 14){

            }else{
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                for(int j = 0; j < 4; j++){
                    if(carpeta.b_content[j].b_inodo == -1){
                        libre = 1;
                        bloque = i;
                        content = j;
                        break;
                    }
                }
            }
        }
        if(libre == 1)
            break;
    }
    return libre;
}

int buscarCarpetaArchivo(FILE *stream, char* path){
    SuperBloque super;
    InodoTable inodo;
    BloqueCarpeta carpeta;
    BloqueApuntadores apuntador;
    QList<string> lista = QList<string>();
    char *token = strtok(path,"/");
    int cont = 0;
    int numInodo = 0;
    while(token != nullptr){
        lista.append(token);
        cont++;
        token = strtok(nullptr,"/");
    }
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    numInodo = super.s_inode_start;

    for (int cont2 = 0; cont2 < cont; cont2++) {
        fseek(stream,numInodo,SEEK_SET);
        fread(&inodo,sizeof(InodoTable),1,stream);
        int siguiente = 0;
        for(int i = 0; i < 15; i++){
            if(inodo.i_block[i] != -1){
                int byteBloque = byteInodoBloque(stream,inodo.i_block[i],'2');
                fseek(stream,byteBloque,SEEK_SET);
                if(i < 12){
                    fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                    for (int j = 0; j < 4; j++) {
                        if((cont2 == cont - 1) && (strcasecmp(carpeta.b_content[j].b_name,lista.at(cont2).c_str()) == 0)){//Tendria que ser la carpeta
                            return carpeta.b_content[j].b_inodo;
                        }else if((cont2 != cont - 1) && (strcasecmp(carpeta.b_content[j].b_name,lista.at(cont2).c_str()) == 0)){
                            numInodo = byteInodoBloque(stream,carpeta.b_content[j].b_inodo,'1');
                            siguiente = 1;
                            break;
                        }
                    }
                }else if(i == 12){
                    fread(&apuntador,sizeof(BloqueApuntadores),1,stream);
                    for(int j = 0; j < 16; j++){
                        if(apuntador.b_pointer[j] != -1){
                            byteBloque = byteInodoBloque(stream,apuntador.b_pointer[j],'2');
                            fseek(stream,byteBloque,SEEK_SET);
                            fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                            for (int k = 0; k < 4; k++) {
                                if((cont2 == cont - 1) && (strcasecmp(carpeta.b_content[k].b_name,lista.at(cont2).c_str()) == 0)){//Tendria que ser la carpeta
                                    return carpeta.b_content[k].b_inodo;
                                }else if((cont2 != cont - 1) && (strcasecmp(carpeta.b_content[k].b_name,lista.at(cont2).c_str()) == 0)){
                                    numInodo = byteInodoBloque(stream,carpeta.b_content[k].b_inodo,'1');
                                    siguiente = 1;
                                    break;
                                }
                            }
                            if(siguiente == 1)
                                break;
                        }
                    }
                }else if(i == 13){

                }else if(i == 14){

                }
                if(siguiente == 1)
                    break;
            }
        }
    }

    return -1;
}

int buscarExisteGrupo(QString name){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    char cadena[400] = "\0";
    SuperBloque super;
    InodoTable inodo;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }
    fclose(fp);
    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        char group[12];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){
            token2 = strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"G") == 0){
                strcpy(group,end_token);
                if(strcmp(group,name.toStdString().c_str()) == 0)
                    return atoi(id);
            }
        }
        token = strtok_r(nullptr,"\n",&end_str);
    }
    return -1;
}

bool buscarExisteUsuario(QString name){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    char cadena[400] = "\0";
    SuperBloque super;
    InodoTable inodo;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }

    fclose(fp);
    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        char user[12];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){//Verificar que no sea un U/G eliminado
            token2 = strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"U") == 0){
                token2 = strtok_r(nullptr,",",&end_token);
                token2 = strtok_r(nullptr,",",&end_token);
                strcpy(user,token2);

                if(strcmp(user,name.toStdString().c_str()) == 0)
                    return true;
            }
        }
        token = strtok_r(nullptr,"\n",&end_str);
    }

    return false;
}

void eliminarGrupo(QString name){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    SuperBloque super;
    InodoTable inodo;
    BloqueArchivo archivo;
    int col = 1;
    char actual;
    int posicion = 0;
    int numBloque = 0;
    int id = -1;
    char tipo = '\0';
    string grupo = "";
    string palabra = "";
    bool Bandera = false;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);

    for (int i = 0; i < 12; i++) {
        if(inodo.i_block[i] != -1){
            fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*inodo.i_block[i],SEEK_SET);
            fread(&archivo,sizeof(BloqueArchivo),1,fp);
            for(int j = 0; j < 63; j++){
                actual = archivo.b_content[j];
                if(actual=='\n'){
                    if(tipo == 'G'){
                        grupo = palabra;
                        if(strcmp(grupo.c_str(),name.toStdString().c_str()) == 0){
                            fseek(fp,super.s_block_start+static_cast<int>(sizeof(BloqueArchivo))*numBloque,SEEK_SET);
                            fread(&archivo,sizeof(BloqueCarpeta),1,fp);
                            archivo.b_content[posicion] = '0';
                            fseek(fp,super.s_block_start+static_cast<int>(sizeof(BloqueArchivo))*numBloque,SEEK_SET);
                            fwrite(&archivo,sizeof(BloqueArchivo),1,fp);
                            cout << "Grupo eliminado con exito" << endl;
                            Bandera = true;
                            break;
                        }
                    }
                    col = 1;
                    palabra = "";
                }else if(actual != ','){
                    palabra += actual;
                    col++;
                }else if(actual == ','){
                    if(col == 2){
                        id = atoi(palabra.c_str());
                        posicion = j-1;
                        numBloque = inodo.i_block[i];
                    }
                    else if(col == 4)
                        tipo = palabra[0];
                    col++;
                    palabra = "";
                }
            }
            if(Bandera)
                break;
        }
    }

    fclose(fp);
}

void eliminarUsuario(QString name){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    SuperBloque super;
    InodoTable inodo;
    BloqueArchivo archivo;
    int col = 1;
    char actual;
    string palabra = "";
    int posicion = 0;
    int numBloque = 0;
    int id = -1;
    char tipo = '\0';
    string grupo = "";
    string usuario = "";
    bool Bandera = false;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for (int i = 0; i < 12; i++) {
        if(inodo.i_block[i] != -1){
            fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*inodo.i_block[i],SEEK_SET);
            fread(&archivo,sizeof(BloqueArchivo),1,fp);
            for(int j = 0; j < 63; j++){
                actual = archivo.b_content[j];
                if(actual=='\n'){
                    if(tipo == 'U'){
                        if(strcmp(usuario.c_str(),name.toStdString().c_str()) == 0){
                            fseek(fp,super.s_block_start+static_cast<int>(sizeof(BloqueArchivo))*numBloque,SEEK_SET);
                            fread(&archivo,sizeof(BloqueArchivo),1,fp);
                            archivo.b_content[posicion] = '0';
                            fseek(fp,super.s_block_start+static_cast<int>(sizeof(BloqueArchivo))*numBloque,SEEK_SET);
                            fwrite(&archivo,sizeof(BloqueArchivo),1,fp);
                            cout << "Usuario eliminado con exito" << endl;
                            Bandera = true;
                            break;
                        }
                        usuario = "";
                        grupo = "";
                    }
                    col = 1;
                    palabra = "";
                }else if(actual != ','){
                    palabra += actual;
                    col++;
                }else if(actual == ','){
                    if(col == 2){
                        id = atoi(palabra.c_str());
                        posicion = j-1;
                        numBloque = inodo.i_block[i];
                    }
                    else if(col == 4)
                        tipo = palabra[0];
                    else if(grupo == "")
                        grupo = palabra;
                    else if(usuario == "")
                        usuario = palabra;
                    col++;
                    palabra = "";
                }
            }
            if(Bandera)
                break;
        }
    }
    fclose(fp);
}

void eliminarCarpetaArchivo(FILE *stream, int n){
    SuperBloque super;
    InodoTable inodo;
    char buffer = '0';
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*n,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);
    if(inodo.i_type == '0'){
        for (int i = 0; i < 12; i++) {
            if(inodo.i_block[i] != -1){
                char byte = '0';
                fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
                byte = static_cast<char>(fgetc(stream));
                if(byte == '1'){
                    BloqueCarpeta carpeta;
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                    fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                    for (int j = 0; j < 4; j++) {
                        if(carpeta.b_content[j].b_inodo != -1){
                            if(strcmp(carpeta.b_content[j].b_name,".")!=0 && strcmp(carpeta.b_content[j].b_name,"..")!=0)
                                eliminarCarpetaArchivo(stream,carpeta.b_content[j].b_inodo);
                        }
                    }
                    fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
                    fputc(buffer,stream);
                }
            }
        }
        fseek(stream,super.s_bm_inode_start + n,SEEK_SET);
        fputc(buffer,stream);
    }else{
        for (int i = 0; i < 15; i++) {
            if(inodo.i_block[i] != -1){
                fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
                fputc(buffer,stream);
            }
        }
        fseek(stream,super.s_bm_inode_start + n,SEEK_SET);
        fputc(buffer,stream);
    }

}

int getID_usr(){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    char cadena[400] = "\0";
    int res = 0;
    SuperBloque super;
    InodoTable inodo;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }

    fclose(fp);
    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){
            token2 = strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"U") == 0)
                res++;
        }
        token = strtok_r(nullptr,"\n",&end_str);
    }
    return ++res;
}

int getID_grp(){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    char cadena[400] = "\0";
    int aux_id = -1;
    SuperBloque super;
    InodoTable inodo;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }

    fclose(fp);

    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){
            token2 = strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"G") == 0){
                aux_id = atoi(id);
            }

        }
        token = strtok_r(nullptr,"\n",&end_str);
    }
    return ++aux_id;
}

int buscarBit(FILE *fp, char tipo, char fit){
    SuperBloque super;
    int inicio_bm = 0;
    char tempBit = '0';
    int bit_libre = -1;
    int tam_bm = 0;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    if(tipo == 'I'){
        tam_bm = super.s_inodes_count;
        inicio_bm = super.s_bm_inode_start;
    }else if(tipo == 'B'){
        tam_bm = super.s_blocks_count;
        inicio_bm = super.s_bm_block_start;
    }
    if(fit == 'F'){
        for(int i = 0; i < tam_bm; i++){
            fseek(fp,inicio_bm + i,SEEK_SET);
            tempBit = static_cast<char>(fgetc(fp));
            if(tempBit == '0'){
                bit_libre = i;
                return bit_libre;
            }
        }

        if(bit_libre == -1)
            return -1;
    }else if(fit == 'B'){
        int libres = 0;
        int auxLibres = -1;

        for(int i = 0; i < tam_bm; i++){
            fseek(fp,inicio_bm + i,SEEK_SET);
            tempBit = static_cast<char>(fgetc(fp));
            if(tempBit == '0'){
                libres++;
                if(i+1 == tam_bm){
                    if(auxLibres == -1 || auxLibres == 0)
                        auxLibres = libres;
                    else{
                        if(auxLibres > libres)
                            auxLibres = libres;
                    }
                    libres = 0;
                }
            }else if(tempBit == '1'){
                if(auxLibres == -1 || auxLibres == 0)
                    auxLibres = libres;
                else{
                    if(auxLibres > libres)
                        auxLibres = libres;
                }
                libres = 0;
            }
        }
        for(int i = 0; i < tam_bm; i++){
            fseek(fp,inicio_bm + i, SEEK_SET);
            tempBit = static_cast<char>(fgetc(fp));
            if(tempBit == '0'){
                libres++;
                if(i+1 == tam_bm)
                    return ((i+1)-libres);
            }else if(tempBit == '1'){
                if(auxLibres == libres)
                    return ((i+1) - libres);
                libres = 0;
            }
        }

        return -1;
    }else if(fit == 'W'){
        int libres = 0;
        int auxLibres = -1;

        for (int i = 0; i < tam_bm; i++) {
            fseek(fp,inicio_bm + i, SEEK_SET);
            tempBit = static_cast<char>(fgetc(fp));
            if(tempBit == '0'){
                libres++;
                if(i+1 == tam_bm){
                    if(auxLibres == -1 || auxLibres == 0)
                        auxLibres = libres;
                    else {
                        if(auxLibres < libres)
                            auxLibres = libres;
                    }
                    libres = 0;
                }
            }else if(tempBit == '1'){
                if(auxLibres == -1 || auxLibres == 0)
                    auxLibres = libres;
                else{
                    if(auxLibres < libres)
                        auxLibres = libres;
                }
                libres = 0;
            }
        }

        for (int i = 0; i < tam_bm; i++) {
            fseek(fp,inicio_bm + i, SEEK_SET);
            tempBit = static_cast<char>(fgetc(fp));
            if(tempBit == '0'){
                libres++;
                if(i+1 == tam_bm)
                    return ((i+1) - libres);
            }else if(tempBit == '1'){
                if(auxLibres == libres)
                    return ((i+1) - libres);
                libres = 0;
            }
        }
        return -1;
    }
    return 0;
}

void agregarUsersTXT(QString datos){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(), "rb+");
    SuperBloque super;
    InodoTable inodo;
    BloqueArchivo archivo;
    int blockIndex = 0;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);
    for(int i = 0; i < 12; i++){
        if(inodo.i_block[i] != -1)
            blockIndex = inodo.i_block[i];
    }
    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*blockIndex,SEEK_SET);
    fread(&archivo,sizeof(BloqueArchivo),1,fp);
    int enUso = static_cast<int>(strlen(archivo.b_content));
    int libre = 63 - enUso;
    if(datos.length() <= libre){
        strcat(archivo.b_content,datos.toStdString().c_str());
        fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*blockIndex,SEEK_SET);
        fwrite(&archivo,sizeof(BloqueArchivo),1,fp);
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        fread(&inodo,sizeof(InodoTable),1,fp);
        inodo.i_size = inodo.i_size + datos.length();
        inodo.i_mtime = time(nullptr);
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        fwrite(&inodo,sizeof(InodoTable),1,fp);
    }else{
        QString aux = "";
        QString aux2 = "";
        int i = 0;
        for(i = 0; i < libre; i++)
            aux += datos.at(i);
        for(; i < datos.length(); i++)
            aux2  += datos.at(i);
        strcat(archivo.b_content,aux.toStdString().c_str());
        fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*blockIndex,SEEK_SET);
        fwrite(&archivo,sizeof(BloqueArchivo),1,fp);
        BloqueArchivo auxArchivo;
        strcpy(auxArchivo.b_content,aux2.toStdString().c_str());
        int bit = buscarBit(fp,'B',currentSession.fit);
        fseek(fp,super.s_bm_block_start + bit,SEEK_SET);
        fputc('2',fp);
        fseek(fp,super.s_block_start + (static_cast<int>(sizeof(BloqueArchivo))*bit),SEEK_SET);
        fwrite(&auxArchivo,sizeof(BloqueArchivo),1,fp);
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        fread(&inodo,sizeof(InodoTable),1,fp);
        inodo.i_size = inodo.i_size + datos.length();
        inodo.i_mtime = time(nullptr);
        inodo.i_block[blockIndex] = bit;
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        fwrite(&inodo,sizeof(InodoTable),1,fp);
        super.s_first_blo = super.s_first_blo + 1;
        super.s_free_blocks_count = super.s_free_blocks_count - 1;
        fseek(fp,currentSession.inicioSuper,SEEK_SET);
        fwrite(&super,sizeof(SuperBloque),1,fp);
    }
    fclose(fp);
}

void guardarJournal(char* operacion,int tipo,int permisos,char *nombre,char *content){
    SuperBloque super;
    Journal registro;
    strcpy(registro.journal_operation_type,operacion);
    registro.journal_type = tipo;
    strcpy(registro.journal_name,nombre);
    strcpy(registro.journal_content,content);
    registro.journal_date = time(nullptr);
    registro.journal_owner = currentSession.id_user;
    registro.journal_permissions = permisos;
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    Journal registroAux;
    bool ultimo = false;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    int inicio_journal = currentSession.inicioSuper + static_cast<int>(sizeof(SuperBloque));
    int final_journal = super.s_bm_inode_start;
    fseek(fp,inicio_journal,SEEK_SET);
    while((ftell(fp) < final_journal) && !ultimo){
        fread(&registroAux,sizeof(Journal),1,fp);
        if(registroAux.journal_type != 0 && registroAux.journal_type != 1)
            ultimo = true;
    }
    fseek(fp,ftell(fp)- static_cast<int>(sizeof(Journal)),SEEK_SET);
    fwrite(&registro,sizeof(Journal),1,fp);
    fclose(fp);
}

int nuevaCarpeta(FILE *stream, char fit, bool BanderaPR, char *path, int index){
    SuperBloque super;
    InodoTable inodo,inodoNuevo;
    BloqueCarpeta carpeta, carpetaNueva, carpetaAux;
    BloqueApuntadores apuntadores;
    QList<string> lista = QList<string>();
    char copiaPath[500];
    char directorio[500];
    char nombreCarpeta[80];
    strcpy(copiaPath,path);
    strcpy(directorio,dirname(copiaPath));
    strcpy(copiaPath,path);
    strcpy(nombreCarpeta,basename(copiaPath));
    char *token = strtok(path,"/");
    int cont = 0;
    int numInodo = index;
    int response = 0;
    while(token != nullptr){
        cont++;
        lista.append(token);
        token = strtok(nullptr,"/");
    }
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    if(cont == 1){
        int content = 0;
        int bloque = 0;
        int pointer = 0;
        int libre = buscarContentenido(stream,numInodo,inodo,carpeta,apuntadores,content,bloque,pointer);
        if(libre == 1){
            if(bloque == 12){
                bool permissions = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permissions || (currentSession.id_user == 1 && currentSession.id_grp == 1) ){
                    char buffer = '1';
                    int bitInodo = buscarBit(stream,'I',fit);
                    carpeta.b_content[content].b_inodo = bitInodo;
                    strcpy(carpeta.b_content[content].b_name,nombreCarpeta);
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*apuntadores.b_pointer[pointer],SEEK_SET);
                    fwrite(&carpeta,sizeof(BloqueCarpeta),1,stream);
                    inodoNuevo = crearInodo(0,'0',664);
                    int bitBloque = buscarBit(stream,'B',fit);
                    inodoNuevo.i_block[0] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    carpetaNueva = crearBloqueCarpeta();
                    carpetaNueva.b_content[0].b_inodo = bitInodo;
                    carpetaNueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpetaNueva.b_content[0].b_name,".");
                    strcpy(carpetaNueva.b_content[1].b_name,"..");
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_free_blocks_count = super.s_free_blocks_count - 1;
                    super.s_first_ino = super.s_first_ino + 1;
                    super.s_first_blo = super.s_first_blo + 1;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }else
                    return 2;
            }else if(bloque == 13){

            }else if(bloque == 14){

            }else{
                bool permisos = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permisos || (currentSession.id_user == 1 && currentSession.id_grp == 1) ){
                    char buffer = '1';
                    int bitInodo = buscarBit(stream,'I',fit);
                    carpeta.b_content[content].b_inodo = bitInodo;
                    strcpy(carpeta.b_content[content].b_name,nombreCarpeta);
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[bloque],SEEK_SET);
                    fwrite(&carpeta,sizeof(BloqueCarpeta),1,stream);
                    inodoNuevo = crearInodo(0,'0',664);
                    int bitBloque = buscarBit(stream,'B',fit);
                    inodoNuevo.i_block[0] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    carpetaNueva = crearBloqueCarpeta();
                    carpetaNueva.b_content[0].b_inodo = bitInodo;
                    carpetaNueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpetaNueva.b_content[0].b_name,".");
                    strcpy(carpetaNueva.b_content[1].b_name,"..");
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_free_blocks_count = super.s_free_blocks_count - 1;
                    super.s_first_ino = super.s_first_ino + 1;
                    super.s_first_blo = super.s_first_blo + 1;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }else
                    return 2;
            }
        }else if(libre == 0){
            bool flag = false;
            pointer = -1;
            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
            fread(&inodo,sizeof(InodoTable),1,stream);
            for (int i = 0; i < 15; i++) {
                if(i == 12){
                    if(inodo.i_block[i] == -1){
                        bloque = 12;
                        flag = true;
                        break;
                    }else{
                        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                        fread(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                        for(int j = 0; j < 16; j++){
                            if(apuntadores.b_pointer[j] == -1){
                                bloque = 12;
                                pointer = j;
                                break;
                            }
                        }
                    }
                    if(flag || pointer!= -1)
                        break;
                }else if(i == 13){

                }else if(i == 14){

                }else{
                    if(inodo.i_block[i] == -1){
                        bloque = i;
                        break;
                    }
                }
            }

            if(bloque == 12 && flag){
                bool permissions = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permissions || (currentSession.id_user == 1 && currentSession.id_grp == 1) ){
                    char buffer = '1';
                    char buffer3 = '3';
                    int bitBloque = buscarBit(stream,'B',fit);
                    inodo.i_block[bloque] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                    fwrite(&inodo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer3,sizeof(char),1,stream);
                    int bitBloqueCarpeta = buscarBit(stream,'B',fit);
                    apuntadores.b_pointer[0] = bitBloqueCarpeta;
                    for(int i = 1; i < 16; i++)
                        apuntadores.b_pointer[i] = -1;
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*bitBloque,SEEK_SET);
                    fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                    int bitInodo = buscarBit(stream,'I',fit);
                    carpetaAux = crearBloqueCarpeta();
                    carpetaAux.b_content[0].b_inodo = bitInodo;
                    strcpy(carpetaAux.b_content[0].b_name,nombreCarpeta);
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloqueCarpeta,SEEK_SET);
                    fwrite(&carpetaAux,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloqueCarpeta,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    inodoNuevo = crearInodo(0,'0',664);
                    bitBloque = buscarBit(stream,'B',fit);
                    inodoNuevo.i_block[0] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    carpetaNueva = crearBloqueCarpeta();
                    carpetaNueva.b_content[0].b_inodo = bitInodo;
                    carpetaNueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpetaNueva.b_content[0].b_name,".");
                    strcpy(carpetaNueva.b_content[1].b_name,"..");
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_free_blocks_count = super.s_free_blocks_count - 3;
                    super.s_first_ino = super.s_first_ino + 1;
                    super.s_first_blo = super.s_first_blo + 3;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }else
                    return 2;
            }else if(bloque == 12 && !flag){
                char buffer = '1';
                int bitBloque = buscarBit(stream,'B',fit);
                apuntadores.b_pointer[pointer] = bitBloque;
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                int bitInodo = buscarBit(stream,'I',fit);
                carpetaAux = crearBloqueCarpeta();
                carpetaAux.b_content[0].b_inodo = bitInodo;
                strcpy(carpetaAux.b_content[0].b_name,nombreCarpeta);
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                fwrite(&carpetaAux,sizeof(BloqueCarpeta),1,stream);
                fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                fwrite(&buffer,sizeof(char),1,stream);
                inodoNuevo = crearInodo(0,'0',664);
                inodoNuevo.i_uid = currentSession.id_user;
                inodoNuevo.i_gid = currentSession.id_grp;
                bitBloque = buscarBit(stream,'B',fit);
                inodoNuevo.i_block[0] = bitBloque;
                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                fwrite(&buffer,sizeof(char),1,stream);
                carpetaNueva = crearBloqueCarpeta();
                carpetaNueva.b_content[0].b_inodo = bitInodo;
                carpetaNueva.b_content[1].b_inodo = numInodo;
                strcpy(carpetaNueva.b_content[0].b_name,".");
                strcpy(carpetaNueva.b_content[1].b_name,"..");
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                fwrite(&buffer,sizeof(char),1,stream);
                super.s_free_inodes_count = super.s_free_inodes_count - 1;
                super.s_free_blocks_count = super.s_free_blocks_count - 2;
                super.s_first_ino = super.s_first_ino + 1;
                super.s_first_blo = super.s_first_blo + 2;
                fseek(stream,currentSession.inicioSuper,SEEK_SET);
                fwrite(&super,sizeof(SuperBloque),1,stream);
                return 1;
            }
            else if(bloque == 13){

            }else if(bloque == 14){

            }else{
                bool permissions = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permissions || (currentSession.id_user == 1 && currentSession.id_grp == 1) ){
                    char buffer = '1';
                    int bitBloque = buscarBit(stream,'B',fit);
                    inodo.i_block[bloque] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                    fwrite(&inodo,sizeof(InodoTable),1,stream);
                    int bitInodo = buscarBit(stream,'I',fit);
                    carpetaAux = crearBloqueCarpeta();
                    carpetaAux.b_content[0].b_inodo = bitInodo;
                    strcpy(carpetaAux.b_content[0].b_name,nombreCarpeta);
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaAux,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    inodoNuevo = crearInodo(0,'0',664);
                    bitBloque = buscarBit(stream,'B',fit);
                    inodoNuevo.i_block[0] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    carpetaNueva = crearBloqueCarpeta();
                    carpetaNueva.b_content[0].b_inodo = bitInodo;
                    carpetaNueva.b_content[1].b_inodo = numInodo;
                    strcpy(carpetaNueva.b_content[0].b_name,".");
                    strcpy(carpetaNueva.b_content[1].b_name,"..");
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_free_blocks_count = super.s_free_blocks_count - 2;
                    super.s_first_ino = super.s_first_ino + 1;
                    super.s_first_blo = super.s_first_blo + 2;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }else
                    return 2;
            }

        }
    }else{
        int existe = buscarCarpetaArchivo(stream,directorio);
        if(existe == -1){
            if(BanderaPR){
                int index = 0;
                string aux = "";
                for(int i = 0; i < cont; i++){
                    aux += "/"+lista.at(i);
                    char dir[500];
                    char auxDir[500];
                    strcpy(dir,aux.c_str());
                    strcpy(auxDir,aux.c_str());
                    int carpeta = buscarCarpetaArchivo(stream,dir);
                    if(carpeta == -1){
                        response = nuevaCarpeta(stream,fit,false,auxDir,index);
                        if(response == 2)
                            break;
                        strcpy(auxDir,aux.c_str());
                        index = buscarCarpetaArchivo(stream,auxDir);
                    }else
                        index = carpeta;
                }
            }else
                return 3;
        }else{
            char dir[100] = "/";
            strcat(dir,nombreCarpeta);
            return nuevaCarpeta(stream,fit,false,dir,existe);
        }
    }

    return response;
}

int crearCarpeta(QString path, bool p){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");

    SuperBloque super;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);

    string aux = path.toStdString();
    char auxPath[500];
    strcpy(auxPath,aux.c_str());
    int existe = buscarCarpetaArchivo(fp,auxPath);
    strcpy(auxPath,aux.c_str());
    int response = -1;

    if(existe != -1)
        response = 0;
    else
        response = nuevaCarpeta(fp,currentSession.fit,p,auxPath,0);

    fclose(fp);

    return response;
}

int nuevoArchivo(FILE *stream, char fit, bool BanderaPR, char *path, int size, QString contenido, int index,char *auxPath){
    SuperBloque super;
    InodoTable inodo,inodoNuevo;
    BloqueCarpeta carpeta, carpetaNueva;
    BloqueApuntadores apuntadores;
    QList<string> lista = QList<string>();
    char copiaPath[500];
    char directorio[500];
    char nombreCarpeta[80];
    string content = "";
    string contentSize = "0123456789";
    strcpy(copiaPath,path);
    strcpy(directorio,dirname(copiaPath));
    strcpy(copiaPath,path);
    strcpy(nombreCarpeta,basename(copiaPath));
    strcpy(copiaPath,path);
    char *token = strtok(path,"/");
    int cont = 0;
    int numInodo = index;
    int finalSize = size;
    while(token != nullptr){
        cont++;
        lista.append(token);
        token = strtok(nullptr,"/");
    }
    if(contenido.length() != 0){
        FILE *archivoCont;
        if((archivoCont = fopen(contenido.toStdString().c_str(),"r"))){
            fseek(archivoCont,0,SEEK_END);
            finalSize = static_cast<int>(ftell(archivoCont));
            fseek(archivoCont,0,SEEK_SET);
            for (int i = 0; i < finalSize; i++)
                content += static_cast<char>(fgetc(archivoCont));
        }else
            return 3;
    }
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    if(cont == 1){
        int bloque = 0;
        int pointer = 0;
        int b_content = 0;
        int libre = buscarContentenido(stream,numInodo,inodo,carpeta,apuntadores,b_content,bloque,pointer);
        if(libre == 1){
            bool permisos = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
            if(permisos || (currentSession.id_user == 1 && currentSession.id_grp == 1)){
                char buffer = '1';
                char buffer2 = '2';
                char buffer3 = '3';
                int bitInodo = buscarBit(stream,'I',fit);
                carpeta.b_content[b_content].b_inodo = bitInodo;
                strcpy(carpeta.b_content[b_content].b_name,nombreCarpeta);
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[bloque],SEEK_SET);
                fwrite(&carpeta,sizeof(BloqueCarpeta),1,stream);
                inodoNuevo = crearInodo(0,'1',664);
                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                fwrite(&buffer,sizeof(char),1,stream);
                if(finalSize != 0){
                    double n = static_cast<double>(finalSize)/static_cast<double>(63);
                    int numBloques = static_cast<int>(ceil(n));
                    int caracteres = finalSize;
                    size_t charNum = 0;
                    size_t contChar = 0;
                    numInodo = buscarCarpetaArchivo(stream,auxPath);
                    for (int i = 0; i < numBloques; i++) {
                        BloqueArchivo archivo;
                        memset(archivo.b_content,0,sizeof(archivo.b_content));
                        if(i == 12){
                            int bitBloqueA = buscarBit(stream,'B',fit);
                            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                            fread(&inodo,sizeof(InodoTable),1,stream);
                            inodo.i_block[i] = bitBloqueA;
                            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                            fwrite(&inodo,sizeof(InodoTable),1,stream);
                            fseek(stream,super.s_bm_block_start + bitBloqueA,SEEK_SET);
                            fwrite(&buffer3,sizeof(char),1,stream);
                            int bitBloque = buscarBit(stream,'B',fit);
                            apuntadores.b_pointer[0] = bitBloque;
                            for(int i = 1; i < 16; i++)
                                apuntadores.b_pointer[i] = -1;
                            fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*bitBloqueA,SEEK_SET);
                            fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                            fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                            fwrite(&buffer2,sizeof(char),1,stream);
                            if(caracteres > 63){
                                for(int j = 0; j < 63; j++){
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                caracteres -= 63;
                            }else{
                                for (int j = 0; j < caracteres; j++) {
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                            }
                        }else if(i > 12 && i < 28){
                            int libre = 0;
                            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                            fread(&inodo,sizeof(InodoTable),1,stream);
                            fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                            fread(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                            for (int a = 0; a < 16; a++) {
                                if(apuntadores.b_pointer[a] == -1){
                                    libre = a;
                                    break;
                                }
                            }
                            int bitBloque = buscarBit(stream,'B',fit);
                            apuntadores.b_pointer[libre] = bitBloque;
                            fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                            fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                            fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                            fwrite(&buffer2,sizeof(char),1,stream);
                            if(caracteres > 63){
                                for(int j = 0; j < 63; j++){
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{//-size
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                caracteres -= 63;
                            }else{
                                for (int j = 0; j < caracteres; j++) {
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                            }
                        }else if(i == 29){

                        }else{
                            int bitBloque = buscarBit(stream,'B',fit);
                            fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                            fwrite(&buffer2,sizeof(char),1,stream);
                            if(caracteres > 63){
                                for(int j = 0; j < 63; j++){
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{//-size
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fread(&inodo,sizeof(InodoTable),1,stream);
                                inodo.i_block[i] = bitBloque;
                                inodo.i_size = finalSize;
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fwrite(&inodo,sizeof(InodoTable),1,stream);
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                caracteres -= 63;
                            }else{
                                for (int j = 0; j < caracteres; j++) {
                                    if(content.length() != 0){
                                        archivo.b_content[j] = content[contChar];
                                        contChar++;
                                    }else{
                                        archivo.b_content[j] = contentSize[charNum];
                                        charNum++;
                                        if(charNum == 10)
                                            charNum = 0;
                                    }
                                }
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fread(&inodo,sizeof(InodoTable),1,stream);
                                inodo.i_block[i] = bitBloque;
                                inodo.i_size = finalSize;
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fwrite(&inodo,sizeof(InodoTable),1,stream);
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                            }
                        }
                    }
                    super.s_free_blocks_count = super.s_free_blocks_count - numBloques;
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_first_ino = super.s_first_ino + 1;
                    super.s_first_blo = super.s_first_blo + numBloques;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }
                super.s_free_inodes_count = super.s_free_inodes_count - 1;
                super.s_first_ino = super.s_first_ino + 1;
                fseek(stream,currentSession.inicioSuper,SEEK_SET);
                fwrite(&super,sizeof(SuperBloque),1,stream);
                return 1;
            }else
                return 2;
        }else{
            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
            fread(&inodo,sizeof(InodoTable),1,stream);
            for (int i = 0; i < 15; i++) {
                if(inodo.i_block[i] == -1){
                    bloque = i;
                    break;
                }
            }
            if(bloque == 12){
                bool permissions = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permissions || (currentSession.id_user == 1 && currentSession.id_grp == 1) ){

                }else
                    return 2;
            }else if(bloque == 13){

            }else if(bloque == 14){

            }else{
                bool permisos = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                if(permisos || (currentSession.id_user == 1 && currentSession.id_grp == 1)){
                    char buffer = '1';
                    char buffer2 = '2';
                    char buffer3 = '3';
                    int bitBloque = buscarBit(stream,'B',fit);
                    inodo.i_block[bloque] = bitBloque;
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                    fwrite(&inodo,sizeof(InodoTable),1,stream);
                    int bitInodo = buscarBit(stream,'I',fit);
                    carpetaNueva.b_content[0].b_inodo = bitInodo;
                    carpetaNueva.b_content[1].b_inodo = -1;
                    carpetaNueva.b_content[2].b_inodo = -1;
                    carpetaNueva.b_content[3].b_inodo = -1;
                    strcpy(carpetaNueva.b_content[0].b_name,nombreCarpeta);
                    strcpy(carpetaNueva.b_content[1].b_name,"");
                    strcpy(carpetaNueva.b_content[2].b_name,"");
                    strcpy(carpetaNueva.b_content[3].b_name,"");
                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
                    fwrite(&carpetaNueva,sizeof(BloqueCarpeta),1,stream);
                    fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    inodoNuevo = crearInodo(0,'1',664);
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*bitInodo,SEEK_SET);
                    fwrite(&inodoNuevo,sizeof(InodoTable),1,stream);
                    fseek(stream,super.s_bm_inode_start + bitInodo,SEEK_SET);
                    fwrite(&buffer,sizeof(char),1,stream);
                    if(finalSize != 0){
                        double n = static_cast<double>(finalSize)/static_cast<double>(63);
                        int numBloques = static_cast<int>(ceil(n));
                        int caracteres = finalSize;
                        size_t charNum = 0;
                        size_t contChar = 0;
                        numInodo = buscarCarpetaArchivo(stream,auxPath);
                        for (int i = 0; i < numBloques; i++) {
                            BloqueArchivo archivo;
                            memset(archivo.b_content,0,sizeof(archivo.b_content));
                            if(i == 12){
                                int bitBloqueA = buscarBit(stream,'B',fit);
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fread(&inodo,sizeof(InodoTable),1,stream);
                                inodo.i_block[i] = bitBloqueA;
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fwrite(&inodo,sizeof(InodoTable),1,stream);
                                fseek(stream,super.s_bm_block_start + bitBloqueA,SEEK_SET);
                                fwrite(&buffer3,sizeof(char),1,stream);
                                int bitBloque = buscarBit(stream,'B',fit);
                                apuntadores.b_pointer[0] = bitBloque;
                                for(int i = 1; i < 16; i++)
                                    apuntadores.b_pointer[i] = -1;
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*bitBloqueA,SEEK_SET);
                                fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                                fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                                fwrite(&buffer2,sizeof(char),1,stream);
                                if(caracteres > 63){
                                    for(int j = 0; j < 63; j++){
                                        if(content.length() != 0){
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{//-size
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                    caracteres -= 63;
                                }else{
                                    for (int j = 0; j < caracteres; j++) {
                                        if(content.length() != 0){
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                }
                            }else if(i > 12 && i < 28){
                                int libre = 0;
                                fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                fread(&inodo,sizeof(InodoTable),1,stream);
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                                fread(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                                for (int a = 0; a < 16; a++) {
                                    if(apuntadores.b_pointer[a] == -1){
                                        libre = a;
                                        break;
                                    }
                                }
                                int bitBloque = buscarBit(stream,'B',fit);
                                apuntadores.b_pointer[libre] = bitBloque;
                                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueApuntadores))*inodo.i_block[12],SEEK_SET);
                                fwrite(&apuntadores,sizeof(BloqueApuntadores),1,stream);
                                fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                                fwrite(&buffer2,sizeof(char),1,stream);
                                if(caracteres > 63){
                                    for(int j = 0; j < 63; j++){
                                        if(content.length() != 0){
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                    caracteres -= 63;
                                }else{
                                    for (int j = 0; j < caracteres; j++) {
                                        if(content.length() != 0){
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                }

                            }else if(i == 29){

                            }else{
                                int bitBloque = buscarBit(stream,'B',fit);
                                fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
                                fwrite(&buffer2,sizeof(char),1,stream);
                                if(caracteres > 63){
                                    for(int j = 0; j < 63; j++){
                                        if(content.length() != 0){//-cont
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{//-size
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                    fread(&inodo,sizeof(InodoTable),1,stream);
                                    inodo.i_block[i] = bitBloque;
                                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                    fwrite(&inodo,sizeof(InodoTable),1,stream);
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                    caracteres -= 63;
                                }else{
                                    for (int j = 0; j < caracteres; j++) {
                                        if(content.length() != 0){
                                            archivo.b_content[j] = content[contChar];
                                            contChar++;
                                        }else{
                                            archivo.b_content[j] = contentSize[charNum];
                                            charNum++;
                                            if(charNum == 10)
                                                charNum = 0;
                                        }
                                    }
                                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                    fread(&inodo,sizeof(InodoTable),1,stream);
                                    inodo.i_block[i] = bitBloque;
                                    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*numInodo,SEEK_SET);
                                    fwrite(&inodo,sizeof(InodoTable),1,stream);
                                    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueArchivo))*bitBloque,SEEK_SET);
                                    fwrite(&archivo,sizeof(BloqueArchivo),1,stream);
                                }
                            }
                        }
                        super.s_free_blocks_count = super.s_free_blocks_count - numBloques;
                        super.s_free_inodes_count = super.s_free_inodes_count - 1;
                        super.s_first_ino = super.s_first_ino + 1;
                        super.s_first_blo = super.s_first_blo + numBloques;
                        fseek(stream,currentSession.inicioSuper,SEEK_SET);
                        fwrite(&super,sizeof(SuperBloque),1,stream);
                        return 1;
                    }
                    super.s_free_inodes_count = super.s_free_inodes_count - 1;
                    super.s_first_ino = super.s_first_ino + 1;
                    fseek(stream,currentSession.inicioSuper,SEEK_SET);
                    fwrite(&super,sizeof(SuperBloque),1,stream);
                    return 1;
                }else
                    return 2;
            }
        }
    }else{
        int existe = buscarCarpetaArchivo(stream,directorio);
        if(existe == -1){
            if(BanderaPR){
                int index = 0;
                string aux = "";
                for (int i = 0; i < cont; i++) {
                    if(i == cont -1){
                        char dir[100] = "/";
                        strcat(dir,nombreCarpeta);
                        return nuevoArchivo(stream,fit,false,dir,size,contenido,index,auxPath);
                    }else{
                        aux += "/"+lista.at(i);
                        char dir[500];
                        char auxDir[500];
                        strcpy(dir,aux.c_str());
                        strcpy(auxDir,aux.c_str());
                        int carpeta = buscarCarpetaArchivo(stream,dir);
                        if(carpeta == -1){
                            nuevaCarpeta(stream,fit,false,auxDir,index);
                            strcpy(auxDir,aux.c_str());
                            index = buscarCarpetaArchivo(stream,auxDir);
                        }else
                            index = carpeta;
                    }
                }
            }else
                return 4;
        }else{
            char dir[100] = "/";
            strcat(dir,nombreCarpeta);
            return nuevoArchivo(stream,fit,false,dir,size,contenido,existe,auxPath);
        }
    }

    return 0;
}

void bloqueCarpetaArchivo(FILE *stream, char* path, int &block, int &posicion,int &pointer,int &posPointer){
    SuperBloque super;
    InodoTable inodo;
    BloqueCarpeta carpeta;
    BloqueApuntadores apuntador;
    QList<string> lista = QList<string>();
    char *token = strtok(path,"/");
    int cont = 0;
    int numInodo = 0;
    while(token != nullptr){
        lista.append(token);
        cont++;
        token = strtok(nullptr,"/");
    }
    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    numInodo = super.s_inode_start;

    for (int cont2 = 0; cont2 < cont; cont2++) {
        fseek(stream,numInodo,SEEK_SET);
        fread(&inodo,sizeof(InodoTable),1,stream);
        int siguiente = 0;
        for(int i = 0; i < 12; i++){
            if(inodo.i_block[i] != -1){
                int byteBloque = byteInodoBloque(stream,inodo.i_block[i],'2');
                fseek(stream,byteBloque,SEEK_SET);
                if(i < 12){
                    fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                    for (int j = 0; j < 4; j++) {
                        if((cont2 == cont - 1) && (strcasecmp(carpeta.b_content[j].b_name,lista.at(cont2).c_str()) == 0)){//Tendria que ser la carpeta
                            block = inodo.i_block[i];
                            posicion = j;
                        }else if((cont2 != cont - 1) && (strcasecmp(carpeta.b_content[j].b_name,lista.at(cont2).c_str()) == 0)){
                            numInodo = byteInodoBloque(stream,carpeta.b_content[j].b_inodo,'1');
                            siguiente = 1;
                            break;
                        }
                    }
                }else if(i == 12){
                    fread(&apuntador,sizeof(BloqueApuntadores),1,stream);
                    for(int j = 0; j < 16; j++){
                        if(apuntador.b_pointer[j] != -1){
                            byteBloque = byteInodoBloque(stream,apuntador.b_pointer[j],'2');
                            fseek(stream,byteBloque,SEEK_SET);
                            fread(&carpeta,sizeof(BloqueCarpeta),1,stream);
                            for (int k = 0; k < 4; k++) {
                                if((cont2 == cont - 1) && (strcasecmp(carpeta.b_content[k].b_name,lista.at(cont2).c_str()) == 0)){//Tendria que ser la carpeta
                                    pointer = inodo.i_block[i];
                                    posPointer = j;
                                    block = apuntador.b_pointer[j];
                                    posicion = k;
                                }else if((cont2 != cont - 1) && (strcasecmp(carpeta.b_content[k].b_name,lista.at(cont2).c_str()) == 0)){
                                    numInodo = byteInodoBloque(stream,carpeta.b_content[k].b_inodo,'1');
                                    siguiente = 1;
                                    break;
                                }
                            }
                            if(siguiente == 1)
                                break;
                        }
                    }
                }
                if(siguiente == 1)
                    break;
            }
        }
    }
}

int crearArchivo(QString path, bool p, int size, QString cont){
    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
    SuperBloque super;
    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    char auxPath[500];
    char auxPath2[500];
    strcpy(auxPath,path.toStdString().c_str());
    strcpy(auxPath2,path.toStdString().c_str());
    int existe = buscarCarpetaArchivo(fp,auxPath);
    strcpy(auxPath,path.toStdString().c_str());
    int response = -1;
    if(existe != -1)
        response = 0;
    else
        response = nuevoArchivo(fp,currentSession.fit,p,auxPath,size,cont,0,auxPath2);

    fclose(fp);
    return response;
}

void moverCarpetaArchivo(FILE *stream, int carpeta, char* path,int destino){
    SuperBloque super;
    InodoTable inodo;

    fseek(stream,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,stream);
    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*carpeta,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);

    if(inodo.i_type == '0'){
        BloqueCarpeta carp;
        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[0],SEEK_SET);
        fread(&carp,sizeof(BloqueCarpeta),1,stream);
        carp.b_content[1].b_inodo = destino;
        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[0],SEEK_SET);
        fwrite(&carp,sizeof(BloqueCarpeta),1,stream);
    }

    BloqueCarpeta carpet;
    int bloque = 0;
    int posicion = 0;
    int pointer = -1;
    int posPointer = 0;
    bloqueCarpetaArchivo(stream,path,bloque,posicion,pointer,posPointer);
    char tempNombre[15];
    int aux = 0;
    bool flag = false;

    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bloque,SEEK_SET);
    fread(&carpet,sizeof(BloqueCarpeta),1,stream);
    strcpy(tempNombre,carpet.b_content[posicion].b_name);
    aux = carpet.b_content[posicion].b_inodo;
    memset(carpet.b_content[posicion].b_name,0,sizeof(carpet.b_content[posicion].b_name));
    carpet.b_content[posicion].b_inodo = -1;
    fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bloque,SEEK_SET);
    fwrite(&carpet,sizeof(BloqueCarpeta),1,stream);

    fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*destino,SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,stream);

    for(int i = 0; i < 12; i++){
        if(inodo.i_block[i] != -1){
            BloqueCarpeta carp;
            char byte = '0';
            fseek(stream,super.s_bm_block_start + inodo.i_block[i],SEEK_SET);
            byte = static_cast<char>(fgetc(stream));
            if(byte == '1'){
                fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                fread(&carp,sizeof(BloqueCarpeta),1,stream);
                for(int j = 0; j < 4; j++){
                    if(carp.b_content[j].b_inodo == -1){
                        strcpy(carp.b_content[j].b_name,tempNombre);
                        carp.b_content[j].b_inodo = aux;
                        fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                        fwrite(&carp,sizeof(BloqueCarpeta),1,stream);
                        flag = true;
                        cout << "Carpeta/archivo movido con exito" << endl;
                        break;
                    }
                }
            }
            if(flag)
                break;
        }else{
            char buffer = '1';
            int bitBloque = buscarBit(stream,'B',currentSession.fit);
            inodo.i_block[i] = bitBloque;
            fseek(stream,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*destino,SEEK_SET);
            fwrite(&inodo,sizeof(InodoTable),1,stream);
            BloqueCarpeta nuevaCarpeta;
            nuevaCarpeta.b_content[0].b_inodo = aux;
            nuevaCarpeta.b_content[1].b_inodo = -1;
            nuevaCarpeta.b_content[2].b_inodo = -1;
            nuevaCarpeta.b_content[3].b_inodo = -1;
            strcpy(nuevaCarpeta.b_content[0].b_name,tempNombre);
            strcpy(nuevaCarpeta.b_content[1].b_name,"");
            strcpy(nuevaCarpeta.b_content[2].b_name,"");
            strcpy(nuevaCarpeta.b_content[3].b_name,"");
            fseek(stream,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bitBloque,SEEK_SET);
            fwrite(&nuevaCarpeta,sizeof(BloqueCarpeta),1,stream);
            fseek(stream,super.s_bm_block_start + bitBloque,SEEK_SET);
            fwrite(&buffer,sizeof(char),1,stream);
            super.s_free_blocks_count = super.s_free_blocks_count - 1;
            super.s_first_blo = super.s_first_blo + 1;
            fseek(stream,currentSession.inicioSuper,SEEK_SET);
            fwrite(&super,sizeof(SuperBloque),1,stream);
            cout << "Carpeta/archivo movido con exito" << endl;
            break;
        }
    }

}

void RMV(Nodo *raiz){
    bool BanderaPath = false;
    bool BanderaDest = false;
    bool Bandera = false;
    QString valPath = "";
    QString valDest = "";

    for(int i = 0; i < raiz->hijos.count(); i ++){
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case PATH:
        {
            if(BanderaPath){
                cout << "ERROR: parametro -path ya definido "<< endl;
                Bandera = true;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case DEST:
        {
            if(BanderaDest){
                cout << "ERROR: parametro -dest ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaDest = true;
            valDest = n.valor;
            valDest = valDest.replace("\"","");
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            if(BanderaDest){
                if(Bandera_login){
                    FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
                    char auxPath[500];
                    char auxDest[500];
                    strcpy(auxPath,valPath.toStdString().c_str());
                    strcpy(auxDest,valDest.toStdString().c_str());
                    int carpeta = buscarCarpetaArchivo(fp,auxPath);
                    int destino = buscarCarpetaArchivo(fp,auxDest);
                    if(carpeta != -1){
                        if(destino != -1){
                            bool permisos = permisosLecturaRecursivo(fp,carpeta);
                            if(permisos){
                                SuperBloque super;
                                InodoTable inodo;
                                fseek(fp,currentSession.inicioSuper,SEEK_SET);
                                fread(&super,sizeof(SuperBloque),1,fp);
                                fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*destino,SEEK_SET);
                                fread(&inodo,sizeof(InodoTable),1,fp);
                                bool permisos2 = permisosDeEscritura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
                                if(permisos2){
                                    char auxP[500];
                                    strcpy(auxP,valPath.toStdString().c_str());
                                    moverCarpetaArchivo(fp,carpeta,auxP,destino);
                                }else
                                    cout << "ERROR: El usuario actual no tiene permisos de escritura en la carpeta destino" << endl;
                            }else
                                cout << "ERROR: El usuario actual no tiene permisos en alguna de las carpetas hijas" << endl;
                        }else
                            cout << "ERROR: No existe la ruta a donde se movera la carpeta/archivo" << endl;
                    }else
                        cout << "ERROR: La carpeta/archivo a mover no existe" << endl;
                    fclose(fp);
                }else
                    cout << "ERROR: Para poder ejecutar este comando necesita iniciar sesion" << endl;
            }else
                cout << "ERROR: parametro -dest no definido "<< endl;
        }else
            cout << "ERROR: parametro -path no definido" << endl;
    }
}

void RMKDIR(Nodo *raiz){
    bool BanderaPath = false;
    bool BanderaP = false;
    bool Bandera = false;
    QString valPath = "";
    bool valP = false;
    for(int i = 0; i < raiz->hijos.count(); i++){
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case PATH:
        {
            if(BanderaPath){
                cout << "ERROR parametro -path ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case P:
        {
            if(BanderaP){
                cout << "ERROR parametro -p ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaP = true;
            valP = true;
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            QFileInfo fileName(valPath);
            QString name = fileName.fileName();
            if(name.length() <= 11){
                if(Bandera_login){
                    int result = crearCarpeta(valPath,valP);
                    if(result == 0)
                        cout << "ERROR: La carpeta ya existe" << endl;
                    else if(result == 1){
                        if(currentSession.tipo_sistema == 3){
                            char aux[500];
                            char operacion[8];
                            char content[5];
                            strcpy(aux,valPath.toStdString().c_str());
                            strcpy(operacion,"mkdir");
                            strcpy(content,"null");
                            guardarJournal(operacion,1,664,aux,content);
                        }
                        cout << "Carpeta creada con exito" << endl;
                    }else if(result == 2)
                        cout << "ERROR: No se tienen permisos de escritura" << endl;
                    else if(result == 3){
                        cout << "ERROR: No existe el directorio y no esta el parametro -p" << endl;
                    }
                }else
                    cout << "ERROR: necesita iniciar sesion para poder ejecutar este comando" << endl;
            }else
                cout << "ERROR: el nombre de la carpeta es mas grande de lo esperado" << endl;
        }else
            cout << "ERROR: parametro -path no definido" << endl;
    }
}

void RREM(Nodo *raiz){
    QString path = raiz->hijos.at(0).valor;
    path = path.replace("\"","");
    char auxPath[500];
    strcpy(auxPath,path.toStdString().c_str());
    if(Bandera_login){
        FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
        int existe = buscarCarpetaArchivo(fp,auxPath);
        strcpy(auxPath,path.toStdString().c_str());
        if(existe != -1){
            bool permisos = permisosEscrituraRecursivo(fp,existe);
            if(permisos){
                SuperBloque super;
                BloqueCarpeta carpeta;
                BloqueApuntadores apuntador;
                int bloque = 0;
                int posicion = 0;
                int pointer = -1;
                int posPointer = 0;
                fseek(fp,currentSession.inicioSuper,SEEK_SET);
                fread(&super,sizeof(SuperBloque),1,fp);
                bloqueCarpetaArchivo(fp,auxPath,bloque,posicion,pointer,posPointer);
                if(pointer == -1){
                    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bloque,SEEK_SET);
                    fread(&carpeta,sizeof(BloqueCarpeta),1,fp);
                    memset(carpeta.b_content[posicion].b_name,0,sizeof(carpeta.b_content[posicion].b_name));
                    carpeta.b_content[posicion].b_inodo = -1;
                    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*bloque,SEEK_SET);
                    fwrite(&carpeta,sizeof(BloqueCarpeta),1,fp);
                }else{
                    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*pointer,SEEK_SET);
                    fread(&apuntador,sizeof(BloqueApuntadores),1,fp);
                    apuntador.b_pointer[posPointer] = -1;
                    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*pointer,SEEK_SET);
                    fwrite(&apuntador,sizeof(BloqueApuntadores),1,fp);
                }
                eliminarCarpetaArchivo(fp,existe);
                cout << "Removido con exito" << endl;
            }else
                cout << "ERROR: alguna carpeta hija no posee permisos de escritura" << endl;
        }else
            cout << "ERROR: no se encuentra la direccion" << endl;
        fclose(fp);
    }else
        cout << "ERROR: para ejecutar este comando necesita iniciar sesion" << endl;
}

void RCAT(Nodo *raiz){
   QString path = raiz->hijos.at(0).valor;
   path = path.replace("\"","");
   char auxPath[500];
   strcpy(auxPath,path.toStdString().c_str());
   if(Bandera_login){
       FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
       int carpeta = buscarCarpetaArchivo(fp,auxPath);
       if(carpeta != -1){
           SuperBloque super;
           InodoTable inodo;
           string cadena = "";
           fseek(fp,currentSession.inicioSuper,SEEK_SET);
           fread(&super,sizeof(SuperBloque),1,fp);
           fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*carpeta,SEEK_SET);
           fread(&inodo,sizeof(InodoTable),1,fp);
           bool permisos = permisosDeLectura(inodo.i_perm,(inodo.i_uid == currentSession.id_user),(inodo.i_gid == currentSession.id_grp));
           if(permisos || (currentSession.id_user == 1 && currentSession.id_grp == 1)){
               for (int i = 0; i < 15; i++) {
                   if(inodo.i_block[i] != -1){
                       if(i < 12){
                           BloqueArchivo archivo;
                           fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                           fread(&archivo,sizeof(BloqueCarpeta),1,fp);
                           cadena += archivo.b_content;
                       }else if(i == 12){
                            BloqueApuntadores apuntador;
                            fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*inodo.i_block[i],SEEK_SET);
                            fread(&apuntador,sizeof(BloqueApuntadores),1,fp);
                            for(int j = 0; j < 16; j++){
                                if(apuntador.b_pointer[j] != -1){
                                    BloqueArchivo archivo;
                                    fseek(fp,super.s_block_start + static_cast<int>(sizeof(BloqueCarpeta))*apuntador.b_pointer[j],SEEK_SET);
                                    fread(&archivo,sizeof(BloqueArchivo),1,fp);
                                    cadena += archivo.b_content;
                                }
                            }
                       }
                   }
               }
               cout << cadena << endl;
           }else
               cout << "ERROR el usuario no tiene permisos de lectura" << endl;
       }else
           cout << "ERROR no se encuentra el archivo " << path.toStdString() << endl;
       fclose(fp);
   }else
       cout << "ERROR para ejecutar este comando necesita iniciar sesion" << endl;
}

void RMKFILE(Nodo *raiz){
    bool BanderaPath = false;
    bool BanderaR = false;
    bool BanderaSize = false;
    bool BanderaCont = false;
    bool Bandera = false;
    QString valPath = "";
    QString valCont = "";
    int valSize = 0;

    for(int i = 0; i < raiz->hijos.count(); i++){
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case PATH:
        {
            if(BanderaPath){
                cout << "ERROR parametro -path ya definido"<< endl;
                Bandera = true;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case R:
        {
            if(BanderaR){
                cout << "ERROR parametro -p ya definido "<< endl;
                Bandera = true;
                break;
            }
            BanderaR = true;
        }
            break;
        case SIZE:
        {
            if(BanderaSize){
                cout << "ERROR parametro -size ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaSize = true;
            valSize = n.valor.toInt();
            if(!(valSize > 0)){
                cout << "ERROR: parametro -size menor a cero" << endl;
                Bandera = true;
                break;
            }
        }
            break;
        case CONT:
        {
            if(BanderaCont){
                cout << "ERROR parametro -cont ya definido " << endl;
                Bandera = true;
                break;
            }
            BanderaCont = true;
            valCont = n.valor;
            valCont = valCont.replace("\"","");
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            QFileInfo fileName(valPath);
            QString name = fileName.fileName();
            if(name.length() <= 11){
                if(Bandera_login){
                    int result = crearArchivo(valPath,BanderaR,valSize,valCont);
                    if(result == 1){
                        if(currentSession.tipo_sistema == 3){
                            char aux[500];
                            char operacion[8];
                            char content[500];
                            strcpy(aux,valPath.toStdString().c_str());
                            strcpy(operacion,"mkfile");
                            strcpy(content,valCont.toStdString().c_str());
                            if(valCont.length() != 0)
                                guardarJournal(operacion,0,664,aux,content);
                            else{
                                strcpy(content,to_string(valSize).c_str());
                                guardarJournal(operacion,0,664,aux,content);
                            }
                        }
                        cout << "Archivo creado con exito" << endl;
                    }else if(result == 2)
                        cout << "ERROR el usuario actual no tiene permisos de escritura" << endl;
                    else if(result == 3)
                        cout << "ERROR el archivo contenido no existe" << endl;
                    else if(result == 4)
                        cout << "ERROR no existe la ruta y no se tiene el parametro -p" << endl;
                }else
                    cout << "ERROR necesita iniciar sesion para poder ejecutar este comando" << endl;
            }else
                cout << "ERROR el nombre del archivo es mas grande de lo esperado" << endl;

        }else
            cout << "ERROR parametro -path no definido" << endl;
    }
}

void RCHMOD(Nodo *raiz){
    bool BanderaPath = false;
    bool BanderaUgo = false;
    bool BanderaR = false;
    bool Bandera = false;
    QString valPath = "";
    QString valUgo = "";
    for(int i = 0; i < raiz->hijos.count(); i++){
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case PATH:
        {
            if(BanderaPath){
                cout << "ERROR parametro -path ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaPath = true;
            valPath = n.valor;
            valPath = valPath.replace("\"","");
        }
            break;
        case UGO:
        {
            if(BanderaUgo){
                cout << "ERRROR parametro -ugo ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaUgo = true;
            valUgo = n.valor;
        }
            break;
        case R:
        {
            if(BanderaR){
                cout << "ERROR parametro -r ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaR = true;
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaPath){
            if(BanderaUgo){
                if(Bandera_login){
                    int propietario =valUgo.at(0).unicode() - '0';
                    int grupo = valUgo.at(1).unicode() - '0';
                    int otros = valUgo.at(2).unicode() - '0';
                    if((propietario >= 0 && propietario <= 7) && (grupo >= 0 && grupo <= 7) && (otros >= 0 && otros <= 7)){
                        char auxPath[500];
                        strcpy(auxPath,valPath.toStdString().c_str());
                        FILE *fp = fopen(currentSession.direccion.toStdString().c_str(),"rb+");
                        SuperBloque super;
                        InodoTable inodo;
                        int existe = buscarCarpetaArchivo(fp,auxPath);
                        fseek(fp,currentSession.inicioSuper,SEEK_SET);
                        fread(&super,sizeof(SuperBloque),1,fp);
                        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*existe,SEEK_SET);
                        fread(&inodo,sizeof(InodoTable),1,fp);
                        if(existe != -1){
                            if((currentSession.id_user ==1 && currentSession.id_grp == 1) || currentSession.id_user == inodo.i_uid){
                                if(BanderaR)
                                    cambiarPermisosRecursivo(fp,existe,valUgo.toInt());
                                else{
                                    inodo.i_perm = valUgo.toInt();
                                    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable))*existe,SEEK_SET);
                                    fwrite(&inodo,sizeof(InodoTable),1,fp);
                                }
                                cout << "Permisos cambiados exitosamente" << endl;
                            }else
                                cout << "ERROR: Para cambiar los permisos debe ser el usuario root o ser dueno de la carpeta/archivo" << endl;
                        }else
                            cout << "ERROR: La ruta no existe" << endl;
                        fclose(fp);
                    }else
                        cout << "ERROR: alguno de los digitos se sale del rango predeterminado"<< endl;
                }else
                    cout << "ERROR: Se necesita iniciar sesion para poder ejecutar este comando" << endl;
            }else
                cout << "ERROR: Parametro -ugo no definido" << endl;
        }else
            cout << "ERROR: Parametro -path no definido" << endl;
    }
}

void RRMUSR(Nodo *raiz){
    QString userName = raiz->hijos.at(0).valor;
    userName = userName.replace("\"","");
    if(Bandera_login){
        if(currentSession.id_user == 1 && currentSession.id_grp == 1){
            if(buscarExisteUsuario(userName)){
                eliminarUsuario(userName);
            }else
                cout << "ERROR el usuario no existe" << endl;
        }else
           cout << "ERROR solo el usuario root puede ejecutar este comando" << endl;
    }else
        cout << "ERROR necesita iniciar sesion para poder ejecutar este comando" << endl;
}

void RMKUSR(Nodo *raiz){
    bool BanderaUser = false;
    bool BanderaPassword = false;
    bool BanderaGroup = false;
    bool Bandera = false;
    QString user = "";
    QString pass = "";
    QString group = "";
    for(int i = 0; i < raiz->hijos.count(); i++){
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case USER:
        {
            if(BanderaUser){
                cout << "ERROR parametro -usuario ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaUser = true;
            user = n.valor;
            user = user.replace("\"","");
        }
            break;
        case PASSWORD:
        {
            if(BanderaPassword){
                cout << "ERROR parametro -pwd ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaPassword = true;
            pass = n.valor;
            pass = pass.replace("\"","");
        }
            break;
        case GROUP:
        {
            if(BanderaGroup){
                cout << "ERROR parametro -grp ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaGroup = true;
            group = n.valor;
            group = group.replace("\"","");
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaUser){
            if(BanderaPassword){
                if(BanderaGroup){
                    if(user.length() <= 10){
                        if(pass.length() <= 10){
                            if(group.length() <= 10){
                                if(Bandera_login){
                                    if(currentSession.id_user == 1 && currentSession.id_grp == 1){
                                        if(buscarExisteGrupo(group) != -1){
                                            if(!buscarExisteUsuario(user)){
                                                int id = getID_usr();
                                                QString datos = QString::number(id) + ",U,"+group+","+user+","+pass+"\n";
                                                agregarUsersTXT(datos);
                                                cout << "Usuario creado con exito " << endl;
                                                if(currentSession.tipo_sistema ==3){
                                                    char aux[64];
                                                    char operacion[10];
                                                    char content[2];
                                                    strcpy(aux,datos.toStdString().c_str());
                                                    strcpy(operacion,"mkusr");
                                                    memset(content,0,2);
                                                    guardarJournal(operacion,0,0,aux,content);
                                                }
                                            }else
                                                cout << "ERROR el usuario ya existe" <<endl;
                                        }else
                                            cout << "ERROR no se encuentra el grupo al que pertenecera el usuario " << endl;
                                    }else
                                        cout << "ERROR solo el usuario root puede ejecutar este comando" << endl;
                                }else
                                    cout << "ERROR necesita iniciar sesion para poder ejecutar este comando" << endl;
                            }else
                                cout << "ERROR grupo del usuario excede de los 10 caracteres permitidos" << endl;
                        }else
                            cout << "ERROR contrasena de usuario excede de los 10 caracteres permitidos" << endl;
                    }else
                        cout << "ERROR nombre de usuario excede de los 10 caracteres permitidos" << endl;
                }else
                    cout << "ERROR parametro -grp no definido" << endl;
            }else
                cout << "ERROR parametro -pwd no definido"<< endl;
        }else
            cout << "ERROR parametro -usuario no definido " << endl;
    }

}

void RRMGRP(Nodo *raiz){
    QString grpName = raiz->hijos.at(0).valor;
    grpName = grpName.replace("\"","");
    if(Bandera_login){
        if(currentSession.id_user == 1 && currentSession.id_grp == 1){
            int grupo = buscarExisteGrupo(grpName);
            if(grupo != -1){
                eliminarGrupo(grpName);
            }else
                cout << "ERROR el grupo no existe" << endl;
        }else
           cout << "ERROR solo el usuario root puede ejecutar este comando" << endl;
    }else
        cout << "ERROR necesita iniciar sesion para poder ejecutar este comando" << endl;
}

void RMKGRP(Nodo *raiz){
    QString grpName = raiz->hijos.at(0).valor;
    grpName = grpName.replace("\"","");
    if(Bandera_login){
        if(currentSession.id_user == 1 && currentSession.id_grp == 1){
            if(grpName.length() <= 10){
                int grupo = buscarExisteGrupo(grpName);
                if(grupo == -1){
                    int idGrp = getID_grp();
                    QString nuevoGrupo = QString::number(idGrp)+",G,"+grpName+"\n";
                    agregarUsersTXT(nuevoGrupo);
                    cout << "Grupo creado con exito "<< endl;
                    if(currentSession.tipo_sistema ==3){
                        char aux[64];
                        char operacion[10];
                        char content[2];
                        strcpy(aux,nuevoGrupo.toStdString().c_str());
                        strcpy(operacion,"mkgrp");
                        memset(content,0,2);
                        guardarJournal(operacion,0  ,0,aux,content);
                    }
                }else
                    cout << "ERROR ya existe un grupo con ese nombre" << endl;
            }else
                cout << "ERROR el nombre del grupo no puede exceder los 10 caracters" << endl;
        }else
            cout << "ERROR solo el usuario root puede ejecutar este comando" << endl;
    }else
        cout << "ERROR necesita iniciar sesion para poder ejecutar este comando" << endl;
}

int verificarDatos(QString user, QString password, QString direccion){
    FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");

    char cadena[400] = "\0";
    SuperBloque super;
    InodoTable inodo;

    fseek(fp,currentSession.inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);

    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }

    fclose(fp);

    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        QString group;
        char user_[12];
        char password_[12];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){
            token2=strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"U") == 0){
                token2 = strtok_r(nullptr,",",&end_token);
                group = token2;
                token2 = strtok_r(nullptr,",",&end_token);
                strcpy(user_,token2);
                token2 = strtok_r(nullptr,",",&end_token);
                strcpy(password_,token2);
                if(strcmp(user_,user.toStdString().c_str()) == 0){
                    if(strcmp(password_,password.toStdString().c_str()) == 0){
                        currentSession.direccion = direccion;
                        currentSession.id_user = atoi(id);
                        currentSession.id_grp = buscarExisteGrupo(group);
                        return 1;
                    }else
                        return 2;
                }
            }
        }
        token = strtok_r(nullptr,"\n",&end_str);
    }

    return 0;
}

char LogicFit(QString direccion, QString nombre){
    string auxPath = direccion.toStdString();
    string auxName = nombre.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(),"rb+"))){
        int extendida = -1;
        MBR masterboot;
        fseek(fp,0,SEEK_SET);
        fread(&masterboot,sizeof(MBR),1,fp);
        for(int i = 0; i < 4; i++){
            if(masterboot.mbr_partition[i].part_type == 'E'){
                extendida = i;
                break;
            }
        }
        if(extendida != -1){
            EBR extendedBoot;
            fseek(fp, masterboot.mbr_partition[extendida].part_start,SEEK_SET);
            while(fread(&extendedBoot,sizeof(EBR),1,fp)!=0 && (ftell(fp) < masterboot.mbr_partition[extendida].part_start + masterboot.mbr_partition[extendida].part_size)){
                if(strcmp(extendedBoot.part_name, auxName.c_str()) == 0){
                    return extendedBoot.part_fit;
                }
            }
        }
        fclose(fp);
    }
    return -1;
}

int LOG_IN(QString direccion, QString nombre, QString user, QString password){
    int index = buscarParticion_Primaria_Extendida(direccion,nombre);
    if(index != -1){
        MBR masterboot;
        SuperBloque super;
        InodoTable inodo;
        FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");
        fread(&masterboot,sizeof(MBR),1,fp);
        fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
        fread(&super,sizeof(SuperBloque),1,fp);
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        fread(&inodo,sizeof(InodoTable),1,fp);
        fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
        inodo.i_atime = time(nullptr);
        fwrite(&inodo,sizeof(InodoTable),1,fp);
        fclose(fp);
        currentSession.inicioSuper = masterboot.mbr_partition[index].part_start;
        currentSession.fit = masterboot.mbr_partition[index].part_fit;
        currentSession.inicioJournal = masterboot.mbr_partition[index].part_start + static_cast<int>(sizeof(SuperBloque));
        currentSession.tipo_sistema = super.s_filesystem_type;
        return verificarDatos(user,password, direccion);
    }else{
        index = buscarParticion_Logica(direccion, nombre);
        if(index != -1){
            SuperBloque super;
            InodoTable inodo;
            FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");
            fseek(fp,index + static_cast<int>(sizeof(EBR)),SEEK_SET);
            fread(&super,sizeof(SuperBloque),1,fp);
            fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
            fread(&inodo,sizeof(InodoTable),1,fp);
            fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)),SEEK_SET);
            inodo.i_atime = time(nullptr);
            fwrite(&inodo,sizeof(InodoTable),1,fp);
            fclose(fp);
            currentSession.inicioSuper = index + static_cast<int>(sizeof(EBR));
            currentSession.fit = LogicFit(direccion,nombre);
            return verificarDatos(user,password,direccion);
        }
    }
    return 0;
}

void RLOGIN(Nodo *raiz){
    bool BanderaUser = false;
    bool BanderaPassword = false;
    bool BanderaID = false;
    bool Bandera = false;
    QString user = "";
    QString password = "";
    QString id = "";

    for(int i = 0; i < raiz->hijos.count(); i++) {
        Nodo n = raiz->hijos.at(i);
        switch (n.tipo_) {
        case USER:
        {
            if(BanderaUser){
                cout << "ERROR parametro -usuario ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaUser = true;
            user = n.valor;
            user = user.replace("\"","");
        }
            break;
        case PASSWORD:
        {
            if(BanderaPassword){
                cout << "ERROR parametro -password ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaPassword = true;
            password = n.valor;
            user = user.replace("\"","");
        }
            break;
        case ID:
        {
            if(BanderaID){
                cout << "ERROR parametro -id ya definido" << endl;
                Bandera = true;
                break;
            }
            BanderaID = true;
            id = n.valor;
        }
            break;
        }
    }

    if(!Bandera){
        if(BanderaUser){
            if(BanderaPassword){
                if(BanderaID){
                    if(!Bandera_login){
                        NodoListaMount *aux = lista->getNodo(id);
                        if(aux != nullptr){
                            int res = LOG_IN(aux->direccion,aux->nombre,user,password);
                            if(res == 1){
                                Bandera_login = true;
                                cout << "Sesion iniciada con exito" << endl;
                            }else if(res == 2)
                                cout << "ERROR Contrasena incorrecta" << endl;
                            else if(res == 0)
                                cout << "ERROR Usuario no encontrado" << endl;
                        }else
                            cout << "ERROR no se encuentra ninguna particion montada con ese id " << endl;
                    }else
                        cout << "ERROR sesion activa, cierre sesion para poder volver a iniciar sesion" << endl;
                }else
                    cout << "ERROR parametro -id no definido" << endl;
            }else
                cout << "ERROR parametro -password no definido" << endl;
        }else
            cout << "ERROR parametro -usuario no definido" << endl;
    }
}

//Recovery system

bool isNumber(string numero){
    bool response = false;
    for (int i = 0; i < static_cast<int>(numero.length()); i++) {
        if(isdigit(numero[i]))
            response = true;
        else{
            response = false;
            break;
        }
    }
    return response;
}

Usuario getUsuario(QString direccion,int inicioSuper, int usuario){
    FILE *fp = fopen(direccion.toStdString().c_str(),"rb+");

    char cadena[400] = "\0";
    SuperBloque super;
    InodoTable inodo;
    Usuario response;

    fseek(fp,inicioSuper,SEEK_SET);
    fread(&super,sizeof(SuperBloque),1,fp);
    fseek(fp,super.s_inode_start + static_cast<int>(sizeof(InodoTable)), SEEK_SET);
    fread(&inodo,sizeof(InodoTable),1,fp);

    for(int i = 0; i < 15; i++){
        if(inodo.i_block[i] != -1){
            BloqueArchivo archivo;
            fseek(fp,super.s_block_start,SEEK_SET);
            for(int j = 0; j <= inodo.i_block[i]; j++){
                fread(&archivo,sizeof(BloqueArchivo),1,fp);
            }
            strcat(cadena,archivo.b_content);
        }
    }

    fclose(fp);

    char *end_str;
    char *token = strtok_r(cadena,"\n",&end_str);
    while(token != nullptr){
        char id[2];
        char tipo[2];
        char user[12];
        char grupo[12];
        char *end_token;
        char *token2 = strtok_r(token,",",&end_token);
        strcpy(id,token2);
        if(strcmp(id,"0") != 0){
            token2 = strtok_r(nullptr,",",&end_token);
            strcpy(tipo,token2);
            if(strcmp(tipo,"U") == 0){
                token2 = strtok_r(nullptr,",",&end_token);
                strcpy(grupo,token2);
                token2 = strtok_r(nullptr,",",&end_token);
                strcpy(user,token2);
                int idAux = atoi(id);
                if(idAux == usuario){
                    QString groupName(grupo);
                    response.id_usr = atoi(id);
                    response.id_grp = buscarExisteGrupo(groupName);
                    strcpy(response.username,user);
                    strcpy(response.group,grupo);
                    return response;
                }

            }
        }
        token = strtok_r(nullptr,"\n",&end_str);
    }
    return response;
}

void systemLoss(QString id){
    NodoListaMount *n = lista->getNodo(id);
    if(n != nullptr){
        if(Bandera_login){
            SuperBloque super;
            FILE *fp = fopen(n->direccion.toStdString().c_str(),"rb+");
            int index = buscarParticion_Primaria_Extendida(n->direccion,n->nombre);
            if(index != -1){
                MBR masterboot;
                fread(&masterboot,sizeof(MBR),1,fp);
                fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
                fread(&super,sizeof(SuperBloque),1,fp);
            }else{
                index = buscarParticion_Logica(n->direccion,n->nombre);
                if(index != -1){
                    EBR extendedBoot;
                    fseek(fp,index,SEEK_SET);
                    fread(&extendedBoot,sizeof(EBR),1,fp);
                    fread(&super,sizeof(SuperBloque),1,fp);
                }
            }
            if(super.s_filesystem_type == 3){
                char buffer = '0';
                int inicio = super.s_bm_inode_start;
                int fin = super.s_block_start + super.s_block_size*(super.s_blocks_count);
                fseek(fp,inicio,SEEK_SET);
                for(int i = inicio; i < fin; i++)
                    fputc(buffer,fp);
                fclose(fp);
                cout << "STOP: {Fatal System Error}" << endl;
                cout << "Press enter to try to continue" << endl;
                getchar();
            }else
                cout << "ERROR: No se puede ejecutar en un sistema EXT2" << endl;
        }else
            cout << "ERROR: Se necesita iniciar sesion para poder ejecutar este comando" << endl;
    }else
        cout << "ERROR: No se encuentra la particion" << endl;
}

void systemRecovery(QString id){
    NodoListaMount *n = lista->getNodo(id);
    if(n != nullptr){
        if(Bandera_login){
            int id_usr = currentSession.id_user;
            int id_grp = currentSession.id_grp;
            Usuario user;
            SuperBloque super;
            int index = buscarParticion_Primaria_Extendida(n->direccion,n->nombre);
            int inicioJournal = 0;
            int inicioSuper = 0;
            FILE *fp = fopen(n->direccion.toStdString().c_str(),"rb+");
            if(index != -1){
                MBR masterboot;
                fread(&masterboot,sizeof(MBR),1,fp);
                fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
                inicioJournal = masterboot.mbr_partition[index].part_start + static_cast<int>(sizeof(SuperBloque));
                inicioSuper = masterboot.mbr_partition[index].part_start;
                fread(&super,sizeof(SuperBloque),1,fp);
                FEXT3(masterboot.mbr_partition[index].part_start, masterboot.mbr_partition[index].part_start + masterboot.mbr_partition[index].part_size,n->direccion);
            }else{
                index = buscarParticion_Logica(n->direccion,n->nombre);
                if(index != -1){
                    EBR extendedBoot;
                    fseek(fp,index,SEEK_SET);
                    fread(&extendedBoot,sizeof(EBR),1,fp);
                    fread(&super,sizeof(SuperBloque),1,fp);
                    fseek(fp,index,SEEK_SET);
                    int init = static_cast<int>(ftell(fp));
                    FEXT3(init,init+extendedBoot.part_size,n->direccion);
                }
            }
            Journal registro;
            if(super.s_filesystem_type == 3){
                string punto = "...";
                fseek(fp,inicioJournal,SEEK_SET);
                while(ftell(fp) < super.s_bm_inode_start){
                    fread(&registro,sizeof(Journal),1,fp);
                    if(strcmp(registro.journal_operation_type,"mkgrp") == 0 || strcmp(registro.journal_operation_type,"mkusr") == 0){
                        currentSession.id_user = 1;
                        currentSession.id_grp = 1;
                        QString datos(registro.journal_name);
                        agregarUsersTXT(datos);
                    }else if(strcmp(registro.journal_operation_type,"mkdir") == 0){
                        user = getUsuario(n->direccion,inicioSuper,registro.journal_owner);
                        currentSession.id_user = user.id_usr;
                        currentSession.id_grp = user.id_grp;
                        QString path(registro.journal_name);
                        crearCarpeta(path,true);
                    }else if(strcmp(registro.journal_operation_type,"mkfile") == 0){
                        user = getUsuario(n->direccion,inicioSuper,registro.journal_owner);
                        currentSession.id_user = user.id_usr;
                        currentSession.id_grp = user.id_grp;
                        QString path(registro.journal_name);
                        bool aux = isNumber(registro.journal_content);
                        if(aux){
                            int size = atoi(registro.journal_content);
                            crearArchivo(path,true,size,"");
                        }else
                            crearArchivo(path,true,0,registro.journal_content);
                    }
                    cout << "\rAnalizando Journaling ..." << flush;
                }
                cout << endl;
                for(int i = 1; i <= 20; i++){
                    cout << "\r[-----" << i*5 << "%-----] Restaurando el sistema" << flush;
                    this_thread::sleep_for(chrono::milliseconds(200));
                }
                cout << endl;
                cout << "El sistema se ha restaurado con exito" << endl;
                currentSession.id_user = id_usr;
                currentSession.id_grp = id_grp;
            }else
                cout << "ERROR: No se puede ejecutar en un sistema EXT2" << endl;
        }else
            cout << "ERROR: Se necesita iniciar sesion para poder ejecutar este comando" << endl;
    }else
        cout << "ERROR: No se encuentra la particion" << endl;
}

//Funciones extras

void Menu()
{

    cout << "*-------------- PROYECTO 1 --------------*" <<endl;
    cout << "*       Byron Estuardo Caal Catun        *" <<endl;
    cout << "*               201901907                *" <<endl;
    cout << "*----------------------------------------*" <<endl;
    cout << endl;
    cout << "Ingrese un Comando:" << endl;
}

void REXEC(Nodo *raiz)
{
    QString valPath = raiz->hijos.at(0).valor;
    string auxPath = valPath.toStdString();
    FILE *fp;
    if((fp = fopen(auxPath.c_str(),"r"))){
        char line[400]="";
        memset(line,0,sizeof(line));
        while(fgets(line,sizeof line,fp)){
            if(line[0]!='\n'){
                cout << ">> "<< line;
                leerComando(line);
            }
            memset(line,0,sizeof(line));
        }
        fclose(fp);
    }else
        cout << "ERROR: script no encontrado o no se pudo abrir correctamente" << endl;
}

void Reconocer_Comandos(Nodo *raiz)
{
    switch (raiz->tipo_)
        {
        case MKDISK:
        {
            Nodo n = raiz->hijos.at(0);
            RMKDISK(&n);
        }
            break;
        case RMDISK:
        {
            RRMDISK(raiz);
        }
            break;
        case FDISK:
        {
            Nodo n = raiz->hijos.at(0);
            RFDISK(&n);
        }
            break;
        case MOUNT:
        {
            Nodo n = raiz->hijos.at(0);
            RMOUNT(&n);
        }
            break;
        case UNMOUNT:
        {
            RUNMOUNT(raiz);
        }
            break;
        case MKFS:
        {
            Nodo n = raiz->hijos.at(0);
            RMKFS(&n);
        }
            break;
        case LOGIN:
        {
            Nodo n = raiz->hijos.at(0);
            RLOGIN(&n);
        }
            break;
        case LOGOUT:
        {
            LOG_OUT();
        }
            break;
        case MKGRP:
        {
            RMKGRP(raiz);
        }
            break;
        case RMGRP:
        {
            RRMGRP(raiz);
        }
            break;
        case MKUSR:
        {
            Nodo n = raiz->hijos.at(0);
            RMKUSR(&n);
        }
            break;
        case RMUSR:
        {
            RRMUSR(raiz);
        }
            break;
        case CHMOD:
        {
            Nodo n = raiz->hijos.at(0);
            RCHMOD(&n);
        }
            break;
        case MKFILE:
        {
            Nodo n = raiz->hijos.at(0);
            RMKFILE(&n);
        }
            break;
        case CAT:
        {
            RCAT(raiz);
        }
            break;
        case REM:
        {
            RREM(raiz);
        }
            break;
        case EDIT:
        {
            cout<< "nel amigo, no funciona el EDIT" <<endl;
        }
            break;
        case REN:
        {
            cout<< "nel amigo, no funciona el RENAME" <<endl;
        }
            break;
        case MKDIR:
        {
            Nodo n = raiz->hijos.at(0);
            RMKDIR(&n);
        }
            break;
        case CP:
        {
            cout<< "nel amigo, no funciona el COPY" <<endl;
        }
            break;
        case MV:
        {
            Nodo n = raiz->hijos.at(0);
            RMV(&n);
        }
            break;
        case CHOWN:
        {
            cout<< "nel amigo, no funciona el CHOWN" <<endl;
        }
            break;
        case CHGRP:
        {
            cout<< "nel amigo, no funciona el CHGRP" <<endl;
        }
            break;
        case PAUSE:{
            cout << "Presiona enter para continuar ";
            getchar();
        }
            break;
        case RECOVERY:
        {
            systemRecovery(raiz->hijos.at(0).valor);
        }
            break;
        case LOSS:
        {
            systemLoss(raiz->hijos.at(0).valor);
        }
            break;
        case REP:
        {

        }
            break;
        case EXEC:
        {
            REXEC(raiz);
        }
            break;
        default: printf("ERROR no se reconoce el comando, porfavor intentelo de nuevo");
        }

}

//home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/MIA_Proyecto1_201901907/main.cpp

void leerComando(char comando[400]){
    //cout<<"comando"<<endl;
    //cout<<comando<<endl;
    //cout<<sizeof (comando)<<endl;
    //for(int i=0; i<sizeof (comando); i++){
        //cout<<comando[i]<<endl;
    //}
    if(comando[0] != '#'){
        YY_BUFFER_STATE buffer = yy_scan_string(comando);//Esto no se pa que
        buffer = yy_scan_string(comando);
        //cout<<"yy_scan_string"<<endl;
        //cout <<yy_scan_string(comando)<<endl;
        //cout<<"Esta entrando if"<<endl;
        //cout<<"Revision Comentario"<<endl;
        if(yyparse() == 0){
            if(raiz!=nullptr){
                //cout<<"Entra raiz"<<endl;
                //cout<<raiz->tipo_<<endl;
                Reconocer_Comandos(raiz);
            }
        }else{
            cout << "Comando no reconocido" << endl;
        }
    }
}

int main()
{
    Menu();
    while(Bandera_Global){
        char texto[500];
        cout << ">> ";
        fgets(texto, sizeof(texto), stdin);
        //cout<<"Prueba de que pasa de aqui"<<endl;
        //cout<<texto<<endl;
        leerComando(texto);
        memset(texto,0,400);
    }
}
