#include "mount.h"

ListaMount::ListaMount(){
    primero = nullptr;
}

void ListaMount::insertarNodo(NodoListaMount *nuevo){
    NodoListaMount *aux = primero;
    if(primero == nullptr){
        primero = nuevo;
    }else{
        while(aux->siguiente!=nullptr){
            aux = aux->siguiente;
        }
        aux->siguiente = nuevo;
    }

}

int ListaMount::eliminarNodo(QString ID){
    NodoListaMount *aux = primero;
    QString tempID = "07";
    tempID += QString::number(aux->num) + aux->letra ;
    if(ID == tempID){
        primero = aux->siguiente;
        free(aux);
        return 1;
    }else{
        NodoListaMount *aux2 = nullptr;
        while(aux!=nullptr){
            tempID = "07";
            tempID += QString::number(aux->num) + aux->letra ;
            if(ID == tempID){
                aux2->siguiente = aux->siguiente;
                return 1;
            }
            aux2 = aux;
            aux = aux->siguiente;
        }
    }
    return 0;
}

int ListaMount::buscarLetra(QString direccion, QString nombre){
    NodoListaMount *aux = primero;
    int retorno = 'a';
    while(aux!=NULL){
        if((direccion == aux->direccion) && (nombre == aux->nombre)){
            return -1;
        }else{
            if(direccion == aux->direccion){
                return aux->letra;
            }else if(retorno <= aux->letra){
                retorno++;
            }
        }
        aux = aux->siguiente;
    }
    return retorno;
}

int ListaMount::buscarNumero(QString direccion, QString nombre){
    int retorno = 1;
    NodoListaMount *aux = primero;
    while(aux!=nullptr){
        if((direccion == aux->direccion) && (retorno == aux->num)){
            retorno++;
        }
        aux = aux->siguiente;
    }
    return retorno;
}

bool ListaMount::buscarNodo(QString direccion, QString nombre){
    NodoListaMount *aux = primero;
    while(aux!=nullptr){
        if((aux->direccion == direccion) && (aux->nombre == nombre)){
            return true;
        }
        aux = aux->siguiente;
    }
    return false;
}

NodoListaMount* ListaMount::getNodo(QString id){
    NodoListaMount *aux = primero;
    while(aux!=nullptr){
        QString tempID = "07";
        tempID += QString::number(aux->num) + aux->letra ;
        if(id == tempID){
            return aux;
        }
        aux = aux->siguiente;
    }
    return nullptr;
}

void ListaMount::mostrarLista(){
    NodoListaMount *aux = primero;
    while(aux!=nullptr){
        cout << "|     "<< aux->nombre.toStdString() << "          " << "07"<<aux->letra<<aux->num << endl;
        aux = aux->siguiente;
    }
}
