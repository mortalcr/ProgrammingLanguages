/*
 * Clase concreta AgendaFactoryImpl que implementa la interfaz AgendaFactory y proporciona implementaciones concretas para crear instancias específicas de Contacto y Evento.
 */

package Data;

import java.util.Date;

public class AgendaFactoryImpl implements AgendaFactory {
    @Override
    public Contacto crearContacto(String parentezco, String nombre, String telefono) {
        return new ContactoFamiliar(parentezco, nombre, telefono);
    }

    @Override
    public Evento crearEvento(int cantidadAsistentes, Date fecha, String nombre) {
        return new EventoReunion(cantidadAsistentes, fecha, nombre);
    }
}