#ifndef NODOLISTAMOUNT_H
#define NODOLISTAMOUNT_H

#include <qstring.h>

class NodoListaMount
{
public:
    NodoListaMount(QString, QString, char, int);
    QString direccion;
    QString nombre;
    char letra;
    int num;
    NodoListaMount *siguiente;
};

#endif // NODOLISTAMOUNT_H
