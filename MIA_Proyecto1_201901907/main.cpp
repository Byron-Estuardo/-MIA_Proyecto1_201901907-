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
bool flag_global = true;

void crearParticionPrimaria(QString, QString, int, char, char, QString);
void crearParticionExtendida(QString, QString, int, char, char, QString);
void crearParticionLogica(QString, QString, int, char, char, QString) ;
void eliminarParticion(QString, QString, QString, QString);
void agregarParticion(QString, QString ,int ,char,QString);
bool existeParticion(QString, QString);
int buscarParticion_Primaria_Extendida(QString, QString);
int buscarParticion_Logica(QString, QString);
void Reconocer_Comandos(Nodo*);
void RMKDISK(Nodo*);
void RRMDISK(Nodo*);
void RFDISK(Nodo*);
void RMOUNT(Nodo*);
void RUNMOUNT(Nodo*);
void RMKFS(Nodo*);
void leerComando(char*);
void crearDisco(QString);
void RREP(Nodo*);
void REXEC(Nodo*);
QString getDirectorio(QString);
QString getExtension(QString);
QString getFileName(QString);
void formatearEXT2(int, int , QString);
void formatearEXT3(int, int, QString);

using namespace std;

//enum de la lista de comandos y parametros que reconoce el analizador
//y que nos sive para el recorrido del arbol que nos devuelve el analizador
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

void crearDisco(QString direccion){
    QString aux = getDirectorio(direccion);
    string comando = "sudo mkdir -p \'"+aux.toStdString()+"\'";//pa crear el archivo
    cout<<"direccion";
    cout<<aux.toStdString();
    system(comando.c_str());
    string comando2 = "sudo chmod -R 777 \'"+aux.toStdString()+"\'";//y pa darle los permisos al archivo
    cout<<"permisis direccion";
    cout<<aux.toStdString();
    system(comando2.c_str());
    string arch = direccion.toStdString();
    FILE *fp = fopen(arch.c_str(),"wb");//aqui abrimos el archivo para ver si se creo bien
    if((fp = fopen(arch.c_str(),"wb")))
        fclose(fp);
    else
        cout << "Error al crear el archivo" << endl;
}

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

void formatearEXT2(int inicio, int tamano, QString direccion){
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

void formatearEXT3(int inicio, int tamano, QString direccion){
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
            cout << "¿Seguro que desea eliminar la particion? Y/N : " ;
            getline(cin, opcion);
        }else
            opcion = "Y";
        if(opcion.compare("Y") == 0 || opcion.compare("y") == 0){
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
                        if(archivo == "principal")
                        cout << "Particion extendida eliminada con exito!" << endl;
                    }else{
                        masterboot.mbr_partition[index].part_status = '1';
                        strcpy(masterboot.mbr_partition[index].part_name,"");
                        fseek(fp,0,SEEK_SET);
                        fwrite(&masterboot,sizeof(MBR),1,fp);
                        fseek(fp,masterboot.mbr_partition[index].part_start,SEEK_SET);
                        fwrite(&buffer,1,masterboot.mbr_partition[index].part_size,fp);
                        if(archivo == "principal")
                        cout << "Particion extendida eliminada con exito!" << endl;
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
                        formatearEXT3(inicio,tamano,aux->direccion);
                    }else{
                        cout<< "a prro si jala 2" <<endl;
                        formatearEXT2(inicio,tamano,aux->direccion);
                    }
                    fclose(fp);
                }else{
                    index = buscarParticion_Logica(aux->direccion,aux->nombre);
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
                        int letra = lista->buscarLetra(valPath,valName);
                        if(letra == -1){
                            cout << "ERROR la particion ya se encuentra  montada, no se puede montar de nuevo" << endl;
                        }else{
                            int num = lista->buscarNumero(valPath, valName);
                            char auxLetra = static_cast<char>(letra);
                            string id = "vd";
                            id += auxLetra + to_string(num);
                            NodoListaMount *n = new NodoListaMount(valPath,valName,auxLetra,num);
                            lista->insertarNodo(n);
                            cout << "Particion montada con exito!" << endl;
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

                            int letra = lista->buscarLetra(valPath,valName);
                            if(letra == -1){
                                cout << "ERROR la particion ya se encuentra montada, no se puede montar de nuevo" << endl;
                            }else{
                                int num = lista->buscarNumero(valPath,valName);
                                char auxLetra = static_cast<char>(letra);
                                string id = "vd";
                                id += auxLetra + to_string(num);
                                NodoListaMount *n = new NodoListaMount(valPath, valName, auxLetra, num);
                                lista->insertarNodo(n);
                                cout << "Particion montada con exito!" << endl;
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

void Menu()
{

    cout << "*-------------- PROYECTO 1 --------------*" <<endl;
    cout << "*       Byron Estuardo Caal Catun        *" <<endl;
    cout << "*               201901907                *" <<endl;
    cout << "*----------------------------------------*" <<endl;
    cout << endl;
    cout << "Ingrese un Comando:" << endl;
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

        }
            break;
        case LOGOUT:
        {

        }
            break;
        case MKGRP:
        {

        }
            break;
        case RMGRP:
        {

        }
            break;
        case MKUSR:
        {

        }
            break;
        case RMUSR:
        {

        }
            break;
        case CHMOD:
        {

        }
            break;
        case MKFILE:
        {

        }
            break;
        case CAT:
        {

        }
            break;
        case REM:
        {

        }
            break;
        case EDIT:
        {

        }
            break;
        case REN:
        {

        }
            break;
        case MKDIR:
        {

        }
            break;
        case CP:
        {

        }
            break;
        case MV:
        {

        }
            break;
        case CHOWN:
        {

        }
            break;
        case CHGRP:
        {

        }
            break;
        case PAUSE:{
            cout << "Presiona enter para continuar ";

        }
            break;
        case RECOVERY:
        {

        }
            break;
        case LOSS:
        {

        }
            break;
        case REP:
        {

        }
            break;
        case EXEC:
        {

        }
            break;
        default: printf("ERROR no se reconoce el comando");
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
    while(flag_global){
        char texto[500];
        cout << ">> ";
        fgets(texto, sizeof(texto), stdin);
        //cout<<"Prueba de que pasa de aqui"<<endl;
        //cout<<texto<<endl;
        leerComando(texto);
        memset(texto,0,400);
    }
}
