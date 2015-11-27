//Arbol AVL en GO
// Noviembre 2015 , Javier Sauma
//Basado en AVl-Salvador Pozo / C con Clase: http://c.conclase.net

package main

import "fmt"

const (
	LEFT  = 1
	RIGHT = 2
)

type nodo struct {
	dato, FE           int
	left, right, padre *nodo
}
type Avl struct {
	raiz             *nodo
	actual           *nodo
	contador, altura int
}

func newAvl() *Avl {
	return new(Avl)
}
func Vacio(r *nodo) bool { return r == nil }

func (a *Avl) Insert(x int) {
	aux := a.actual
	aux = nil
	a.actual = a.raiz
	fmt.Println("insertar", x)
	//buscar el dato en el arbol manteniendo un puntero al padre
	for !Vacio(a.actual) && x != a.actual.dato {
		aux = a.actual
		if x > a.actual.dato {
			a.actual = a.actual.right
		} else if x < a.actual.dato {
			a.actual = a.actual.left
		}
	}
	//si se encontro el elemento regresa sin insertar
	if !Vacio(a.actual) {
		return
	}
	// si el aux (padre) es nil entonces el arbol estaba vacio
	if Vacio(aux) {
		a.raiz = &nodo{dato: x}
		//si el dato es menor al nodo inserta en la rama izq
	} else if x < aux.dato {
		aux.left = &nodo{dato: x,
			padre: aux}
		a.Equilibrar(aux, 1, true)
		// si el dato es mayor inserta en la rama der
	} else if x > aux.dato {
		aux.right = &nodo{dato: x,
			padre: aux}
		a.Equilibrar(aux, 2, true)
	}
}
func (a *Avl) Equilibrar(nodo *nodo, rama int, nuevo bool) {
	salir := false
	//recorrer el camino inverso actualizando FE
	for nodo != nil && !salir {
		if nuevo {
			if rama == 1 {
				nodo.FE = nodo.FE - 1
			} else {
				nodo.FE = nodo.FE + 1
			}
		} else {
			if rama == 1 {
				nodo.FE = nodo.FE + 1
			} else {
				nodo.FE = nodo.FE - 1
			}
		}
		//la altura no varia , salir de equilivrar
		if nodo.FE == 0 {
			salir = true
			//rotar  a la derecha y salir
		} else if nodo.FE == -2 {
			if nodo.left.FE == 1 {
				a.RDD(nodo)
			} else {
				a.RSD(nodo)
			}
			salir = true
			//rotar a la izquierda y salir
		} else if nodo.FE == 2 {
			if nodo.right.FE == -1 {
				a.RDI(nodo)
			} else {
				a.RSI(nodo)
			}
			salir = true
		}
		if nodo.padre != nil {
			if nodo.padre.right == nodo {
				rama = 2
			} else {
				rama = 1
			}
		}
		nodo = nodo.padre //calular FE, siguiente nodo del camino
	}
}

// rotacion doble a la derecha
func (a *Avl) RDD(nodo *nodo) {
	fmt.Println("RDD")
	Padre := nodo.padre
	P := nodo
	Q := P.left
	R := Q.right
	B := R.left
	C := R.right
	if Padre != nil {
		if Padre.right == nodo {
			Padre.right = R
		} else {
			Padre.left = R
		}
	} else {
		a.raiz = R
	}
	//Reconstruir arbol
	Q.right = B
	P.left = C
	R.left = Q
	R.right = P
	// Reasignar padres
	R.padre = Padre
	P.padre = R
	Q.padre = R
	if B != nil {
		B.padre = Q
	}
	if C != nil {
		C.padre = P
	}
	//ajustar valores de FE
	switch R.FE {
	case -1:
		Q.FE = 0
		P.FE = 1
	case 0:
		Q.FE = 0
		P.FE = 0
	case 1:
		Q.FE = -1
		P.FE = 0
	}
	R.FE = 0
}

//rotacion doble izq
func (a *Avl) RDI(nodo *nodo) {
	fmt.Println("RDI")
	Padre := nodo.padre
	P := nodo
	Q := P.right
	R := Q.left
	B := R.left
	C := R.right
	if Padre != nil {
		if Padre.right == nodo {
			Padre.right = R
		} else {
			Padre.left = R
		}
	} else {
		a.raiz = R
	}
	//Reconstruir arbol
	P.right = B
	Q.left = C
	R.left = P
	R.right = Q
	//Reasignar padres
	R.padre = Padre
	P.padre = R
	Q.padre = R
	if B != nil {
		B.padre = P
	}
	if C != nil {
		C.padre = Q
	}
	//ajustar valores FE
	switch R.FE {
	case -1:
		P.FE = 0
		Q.FE = 1
	case 0:
		Q.FE = 0
		P.FE = 0
	case 1:
		P.FE = -1
		Q.FE = 0
	}
	R.FE = 0
}

//rotacion simple derecha
func (a *Avl) RSD(nodo *nodo) {
	fmt.Println("RSD")
	Padre := nodo.padre
	P := nodo
	Q := P.left
	B := Q.right
	if Padre != nil {
		if Padre.right == P {
			Padre.right = Q
		} else {
			Padre.left = Q
		}
	} else {
		a.raiz = Q
	}
	//recountruir arbol
	P.left = B
	Q.right = P
	//reasignar padres
	P.padre = Q
	Q.padre = Padre
	if B != nil {
		B.padre = P
	}
	//ajustar FE
	P.FE = 0
	Q.FE = 0
}

//rotacion simple izquierda
func (a *Avl) RSI(nodo *nodo) {
	fmt.Println("RSI")
	Padre := nodo.padre
	P := nodo
	Q := P.right
	B := Q.left
	if Padre != nil {
		if Padre.right == P {
			Padre.right = Q
		} else {
			Padre.left = Q
		}
	} else {
		a.raiz = Q
	}
	//recountruir arbol
	P.right = B
	Q.left = P
	//reasignar padres
	P.padre = Q
	Q.padre = Padre
	if B != nil {
		B.padre = P
	}
	//ajustar FE
	P.FE = 0
	Q.FE = 0
} //funciones de busqueda
func (a *Avl) Search(x int) bool {
	return a.raiz.Buscar(x)
}
func (n *nodo) Search(x int) bool {
	if n == nil {
		return false
	}
	if x == n.dato {
		return true
	}
	if x <= n.dato {
		return n.left.Search(x)
	}
	return n.right.Search(x)
}
func preorder(nodo *nodo) {
	if nodo == nil {
		return
	}
	fmt.Println("nodo: ", nodo.dato)
	preorder(nodo.left)
	preorder(nodo.right)
}
func main() {
	a := newAvl()
	for i := 1; i < 17; i++ {
		a.Insertar(i)
	}
	preorder(a.raiz)

}
