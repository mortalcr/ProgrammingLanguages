/*
 * Creamos una interfaz AgendaFactory que define métodos para crear instancias de Contacto y Evento.
 */

package Data;

import java.util.Date;

public interface AgendaFactory {
    Contacto crearContacto(String parentezco, String nombre, String telefono);
    Evento crearEvento(int cantidadAsistentes, Date fecha, String nombre);
}

