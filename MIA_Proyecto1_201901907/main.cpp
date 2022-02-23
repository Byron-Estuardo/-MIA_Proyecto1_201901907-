#include <iostream>
#include <stdio.h>
#include <sys/stat.h>
#include <fstream>

#include "nodo.h"
#include "scanner.h"
#include "parser.h"

//los extern son para indicar el almacenamiente y valor de una variable
//o una funcion que esta definido en otro archivo
extern int yyparse();
extern Nodo *raiz;

//Declaracion de variables globales
bool flag_global = true;

//Declaracion de funciones
void Menu();
void leerComando(char*);
void Reconocer_Comandos(Nodo*);

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
    COMENTARIO = 18,
};

//Declaracion de las estructuras que necesitamos
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
        cout<<"mkdisk"<<endl;
    }
        break;
    case RMDISK:
    {
        cout<<"rmdisk"<<endl;
    }
        break;
    case FDISK:
    {
        cout<<"fdisk"<<endl;
    }
        break;
    case MOUNT:
    {
        cout<<"mount"<<endl;
    }
        break;
    case UNMOUNT:
    {
        cout<<"unmount"<<endl;
    }
        break;
    case REP:
    {
        cout<<"reportes"<<endl;
    }
        break;
    case EXEC:
    {
        cout<<"exec"<<endl;
    }
        break;
    case COMENTARIO:
    {

    }
        break;
    default: printf("Error no se reconoce el comando");

    }

}

void leerComando(char comando[400]){
    cout<<"comando"<<endl;
    cout<<comando<<endl;
    cout<<sizeof (comando)<<endl;
    for(int i=0; i<sizeof (comando); i++){
        cout<<comando[i]<<endl;
    }
    if(comando[0] != '#'){
        YY_BUFFER_STATE buffer = yy_scan_string(comando);//Esto no se pa que
        cout<<"yy_scan_string"<<endl;
        cout <<yy_scan_string(comando)<<endl;
        //cout<<"Esta entrando if"<<endl;
        cout<<"Revision Comentario"<<endl;
        if(yyparse() == 0){
            if(raiz!=nullptr){
                cout<<"Entra raiz"<<endl;
                cout<<raiz->tipo_<<endl;
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
        printf(">> ");
        fgets(texto, sizeof(texto), stdin);
        cout<<"Prueba de que pasa de aqui"<<endl;
        cout<<texto<<endl;
        leerComando(texto);
        memset(texto,0,400);
    }
}
