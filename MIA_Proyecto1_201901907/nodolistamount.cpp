#include "nodolistamount.h"

NodoListaMount::NodoListaMount(QString direccion, QString nombre, char letra, int num)
{
    this->direccion = direccion;
    this->nombre = nombre;
    this->letra = letra;
    this->num = num;
    this->siguiente = nullptr;
}
