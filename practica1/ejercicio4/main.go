package main

import "fmt"

type calzado struct {
	marca   string
	talla int
	precio   int
  cantidad int
}

type listaCalzados []calzado

var lCalzados listaCalzados

func (l *listaCalzados) agregarCalzado(marca string, talla int, precio int, cantidad int) {
  if talla < 34 || talla > 44 {
    return 
  }
	var calz, err, _ = l.buscarCalzado(marca, talla)  
  if err != -1{
    calz.cantidad += cantidad
    if calz.precio != precio{
      calz.precio = precio
    }
  }else{
    *l = append(*l,calzado{marca: marca, talla: talla, precio: precio, cantidad: cantidad})
  }
}

func (l* listaCalzados) buscarCalzado(marca string, talla int) (*calzado,int,int) {
  var err = -1
  var c *calzado = nil
  index:=0 
  for i:=0;i < len(*l); i++ {
    if (*l)[i].marca == marca && (*l)[i].talla == talla{
      c = &((*l)[i])
      err = 0
      index = i
    }
  }
  return c,err,index
}

func (l *listaCalzados) venderCalzado(marca string,talla int, cantidad int) {
	var calz,err, index = l.buscarCalzado(marca,talla)
  if err != -1 && calz.cantidad >= cantidad {
		calz.cantidad -= cantidad
    if calz.cantidad == 0 {
      *l = append((*l)[:index], (*l)[index+1:]...)
    }
  }
}

func llenarDatos() {
	lCalzados.agregarCalzado("Nike",40,64000,2)
	lCalzados.agregarCalzado("Nike", 41, 66000,4)
	lCalzados.agregarCalzado("Converse", 36, 48000, 3)
  lCalzados.agregarCalzado("Vans", 25, 23000,5) //Comprobación del rango de tallas, que no se registre
  lCalzados.agregarCalzado("Broncos", 34,22000,4)
  lCalzados.agregarCalzado("Lakai", 43, 43000, 5)
  lCalzados.agregarCalzado("Adidas", 40, 42000, 7)
  lCalzados.agregarCalzado("Adidas", 42, 45000, 8)
  lCalzados.agregarCalzado("Gucci", 42, 420000, 2)
  lCalzados.agregarCalzado("DC", 37, 34000, 4)
}

func main () {
  llenarDatos()
  fmt.Println(lCalzados) //Imprime el slice original
  lCalzados.venderCalzado("Nike",40,2) //Algunas ventas
  lCalzados.venderCalzado("Nike",41,2)
  lCalzados.venderCalzado("DC",37,2)
  lCalzados.venderCalzado("Gucci",42,2)
  lCalzados.venderCalzado("Adidas",42,5)
  lCalzados.venderCalzado("Broncos",34,4)
  fmt.Println(lCalzados) //Imprime slice final donde se han eliminado algunos calzados por falta de stock
}
