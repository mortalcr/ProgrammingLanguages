from abc import ABC, abstractmethod

class Representable(ABC):
    @abstractmethod
    def dibujar(self):
        pass

class Texto(Representable):
    def __init__(self, contenido):
        self._contenido = contenido

    def dibujar(self):
        return f"Texto: {self._contenido}"

    def get_contenido(self):
        return self._contenido

    def set_contenido(self, nuevo_contenido):
        self._contenido = nuevo_contenido

class ObjetoGeometrico(Representable, ABC):
    @abstractmethod
    def area(self):
        pass

class Circulo(ObjetoGeometrico):
    def __init__(self, radio):
        self._radio = radio

    def dibujar(self):
        return f"Círculo de radio {self._radio}"

    def area(self):
        return 3.14 * self._radio ** 2

    def get_radio(self):
        return self._radio

    def set_radio(self, nuevo_radio):
        self._radio = nuevo_radio

class Elipse(ObjetoGeometrico):
    def __init__(self, semieje_mayor, semieje_menor):
        self._semieje_mayor = semieje_mayor
        self._semieje_menor = semieje_menor

    def dibujar(self):
        return f"Elipse con semiejes {self._semieje_mayor} y {self._semieje_menor}"

    def area(self):
        return 3.14 * self._semieje_mayor * self._semieje_menor

    def get_semieje_mayor(self):
        return self._semieje_mayor

    def set_semieje_mayor(self, nuevo_semieje_mayor):
        self._semieje_mayor = nuevo_semieje_mayor

    def get_semieje_menor(self):
        return self._semieje_menor

    def set_semieje_menor(self, nuevo_semieje_menor):
        self._semieje_menor = nuevo_semieje_menor

class Rectangulo(ObjetoGeometrico):
    def __init__(self, base, altura):
        self._base = base
        self._altura = altura

    def dibujar(self):
        return f"Rectángulo con base {self._base} y altura {self._altura}"

    def area(self):
        return self._base * self._altura

    def get_base(self):
        return self._base

    def set_base(self, nueva_base):
        self._base = nueva_base

    def get_altura(self):
        return self._altura

    def set_altura(self, nueva_altura):
        self._altura = nueva_altura

class Linea(ObjetoGeometrico):
    def __init__(self, longitud):
        self._longitud = longitud

    def dibujar(self):
        return f"Línea de longitud {self._longitud}"

    def get_longitud(self):
        return self._longitud

    def set_longitud(self, nueva_longitud):
        self._longitud = nueva_longitud

class Cuadrado(Rectangulo):
    def __init__(self, lado):
        super().__init__(lado, lado)

    def dibujar(self):
        return f"Cuadrado de lado {self.get_base()}"

class Grupo(Representable):
    def __init__(self):
        self._elementos = []

    def agregar_elemento(self, elemento):
        self._elementos.append(elemento)

    def dibujar(self):
        return f"Grupo con {len(self._elementos)} elementos"

    def get_elementos(self):
        return self._elementos

class Pagina:
    def __init__(self):
        self._elementos = []

    def agregar_elemento(self, elemento):
        self._elementos.append(elemento)

    def dibujar(self):
        resultado = "Elementos en la página:\n"
        for elemento in self._elementos:
            resultado += elemento.dibujar() + "\n"
        return resultado

class Documento:
    def __init__(self):
        self._paginas = []

    def agregar_pagina(self, pagina):
        self._paginas.append(pagina)

    def dibujar(self):
        resultado = "Contenido del documento:\n"
        for i, pagina in enumerate(self._paginas, start=1):
            resultado += f"Página {i}:\n"
            resultado += pagina.dibujar()
        return resultado

# Uso del programa
if __name__ == "__main__":
    documento = Documento()

    pagina1 = Pagina()
    texto = Texto("Hola, mundo!")
    circulo = Circulo(5)
    pagina1.agregar_elemento(texto)
    pagina1.agregar_elemento(circulo)

    pagina2 = Pagina()
    cuadrado = Cuadrado(6)
    elipse = Elipse(3, 4)
    pagina2.agregar_elemento(cuadrado)
    pagina2.agregar_elemento(elipse)

    documento.agregar_pagina(pagina1)
    documento.agregar_pagina(pagina2)

    print(documento.dibujar())
