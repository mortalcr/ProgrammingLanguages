package main

import (
  "fmt"
  "sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var prod, err, _ = l.buscarProducto(nombre)
  if err != -1{
    prod.cantidad += cantidad
    if prod.precio != precio{
      prod.precio = precio
    }
  }else{
    *l = append(*l,producto{nombre: nombre, cantidad: cantidad, precio: precio})
  }
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var err = -1
  var p *producto = nil
  var i int
  index := 0
  for i = 0; i < len(*l); i++{
    if (*l)[i].nombre == nombre {
      p = &((*l)[i])
      err=0
      index = i
    }
  }
  return p,err, index
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod,err, index = l.buscarProducto(nombre)
	if err != -1 && prod.cantidad >= cantidad {
		prod.cantidad = prod.cantidad - cantidad
    if prod.cantidad == 0 {
      *l = append((*l)[:index], (*l)[index+1:]...)
    }
    }else{
      fmt.Println("No hay en stock")
  }
}

func (l *listaProductos) modificarPrecio(nombre string, precio int) {
	var prod, err, _ = l.buscarProducto(nombre)
  if err != -1 && prod.precio != precio{
    prod.precio = precio
    }
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
  var minProd listaProductos
  for _,prod := range *l {
    if prod.cantidad < existenciaMinima {
      minProd = append(minProd, prod)
    }
  }
    return minProd
}

func (l listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos){
  for _, prod := range l {
    for _, min := range listaMinimos {
      if prod.nombre == min.nombre {
        nombre := prod.nombre
        precio := prod.precio
        cantidad := existenciaMinima - prod.cantidad
        fmt.Println(cantidad)
        lProductos.agregarProducto(nombre,cantidad,precio)
      }
    }
  }
}

func (l *listaProductos) OrdenarPorPrecio() {
    sort.SliceStable(*l, func(i, j int) bool {
        return (*l)[i].precio < (*l)[j].precio
    })
}


func main() {
	llenarDatos()
	fmt.Println(lProductos)
  lProductos.agregarProducto("arroz",2,3000)
  fmt.Println(lProductos)
  lProductos.venderProducto("frijoles",5)
  fmt.Println(lProductos)
  lProductos.modificarPrecio("arroz", 1)
  fmt.Println(lProductos)
  listaMinimos := lProductos.listarProductosMínimos()
  lProductos.aumentarInventarioDeMinimos(listaMinimos)
  fmt.Println(lProductos)
  lProductos.OrdenarPorPrecio()
  fmt.Println(lProductos)
}
