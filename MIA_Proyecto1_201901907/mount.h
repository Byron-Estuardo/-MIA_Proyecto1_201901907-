#ifndef MOUNT_H
#define MOUNT_H

#include <iostream>

#include "nodolistamount.h"

using namespace std;

class ListaMount
{
public:
    NodoListaMount *primero;
    ListaMount();
    int buscarLetra(QString direccion, QString nombre);
    int buscarNumero(QString direccion, QString nombre);
    void mostrarLista();
    void insertarNodo(NodoListaMount*);
    int eliminarNodo(QString);
    bool buscarNodo(QString, QString);
    QString getDireccion(QString);
    NodoListaMount* getNodo(QString);
};

#endif // MOUNT_H
