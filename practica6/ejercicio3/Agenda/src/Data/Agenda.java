/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

 /*
 * Eager Singleton: Este patrón crea la instancia de la clase de forma anticipada, es decir, en el
 * momento de la carga de la clase. En Java, esto se logra mediante la inicialización estática.
 * 
 * Lazy Singleton: Este patrón crea la instancia de la clase solo cuando se solicita por primera 
 * vez. En Java, esto se implementa utilizando una verificación de nulidad en el método de obtención de la instancia.
 * 
 * Se implementó el patrón Singleton utilizando el enfoque Lazy Singleton en la clase Agenda
 * 
 * Se optó por un "Lazy Singleton" porque este enfoque retrasa la creación de la instancia única de la clase Agenda
 * hasta que se solicita por primera vez mediante el método obtenerInstancia()
 * Esto puede ser beneficioso si la creación de la instancia es costosa en términos de recursos y no se necesita hasta un punto posterior en la ejecución del programa.
 */

package Data;

import java.util.Date;
import java.util.LinkedList;
import java.util.Map;

/**
 *
 * @author oviquez
 */
public class Agenda {
    private static Agenda instanciaUnica;
    private LinkedList<Object> listaObjetos;
    private AgendaFactory fabrica;
    
    private LinkedList<Contacto> contactos;
    private LinkedList<Evento> eventos;

    public Agenda() {
        this.listaObjetos = new LinkedList<>();
        this.contactos = new LinkedList<>();
        this.eventos = new LinkedList<>();
        this.fabrica = new AgendaFactoryImpl();
    }

    public LinkedList<Contacto> getContactos() {
        return contactos;
    }

    public LinkedList<Evento> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }

    public void agregarContacto(String parentezco, String nombre, String telefono) {
        Contacto nuevoContacto = fabrica.crearContacto(parentezco, nombre, telefono);
        contactos.add(nuevoContacto);
        listaObjetos.add(nuevoContacto);
    }

    public void agregarEvento(int cantidadAsistentes, Date fecha, String nombre) {
        Evento nuevoEvento = fabrica.crearEvento(cantidadAsistentes, fecha, nombre);
        eventos.add(nuevoEvento);
        listaObjetos.add(nuevoEvento);
    }

    public static Agenda obtenerInstancia() {
        if (instanciaUnica == null) {
            instanciaUnica = new Agenda();
        }
        return instanciaUnica;
    }
    
    
    
}
