from datetime import date

class Socio:
    def __init__(self, numero_socio, nombre, direccion):
        self.numero_socio = numero_socio
        self.nombre = nombre
        self.direccion = direccion
        self.libros_prestados = []

    def prestar_libro(self, libro, fecha):
        if libro.disponible:
            libro.disponible = False
            self.libros_prestados.append(Prestamo(libro.codigo, self.numero_socio, fecha))
            print(f"{self.nombre} ha prestado '{libro.titulo}'")
        else:
            print(f"Lo siento, '{libro.titulo}' no está disponible en este momento.")

    def libros_prestados_count(self):
        return len(self.libros_prestados)

class Libro:
    def __init__(self, codigo, titulo, autor, disponible=True):
        self.codigo = codigo
        self.titulo = titulo
        self.autor = autor
        self.disponible = disponible

    def mostrar_estado(self):
        estado = "Disponible" if self.disponible else "Prestado"
        print(f"Código: {self.codigo}, Título: {self.titulo}, Autor: {self.autor}, Estado: {estado}")

class Prestamo:
    def __init__(self, codigo_libro, numero_socio, fecha):
        self.codigo_libro = codigo_libro
        self.numero_socio = numero_socio
        self.fecha = fecha

#Creacion de los libros y socios
libro1 = Libro(codigo=1, titulo="La Sombra del Viento", autor="Carlos Ruiz Zafón")
libro2 = Libro(codigo=2, titulo="1984", autor="George Orwell")
libro3 = Libro(codigo=3, titulo="Cien años de soledad", autor="Gabriel García Márquez")

socio1 = Socio(numero_socio=101, nombre="Juan Pérez", direccion="Calle A")
socio2 = Socio(numero_socio=102, nombre="Ana Gómez", direccion="Calle B")

# Estado de los libros
libro1.mostrar_estado()
libro2.mostrar_estado()
libro3.mostrar_estado()

# Libros que tienen los socios
print(f"\nLibros prestados por {socio1.nombre}: {socio1.libros_prestados_count()}")
print(f"Libros prestados por {socio2.nombre}: {socio2.libros_prestados_count()}")

# Juan Pérez toma un libro
hoy = date.today()
socio1.prestar_libro(libro1, hoy)

# Mostrar estado del libro que tomó Juan Pérez
libro1.mostrar_estado()

# Mostrar estado de los socios después del préstamo
print(f"\nLibros prestados por {socio1.nombre}: {socio1.libros_prestados_count()}")
print(f"Libros prestados por {socio2.nombre}: {socio2.libros_prestados_count()}")
