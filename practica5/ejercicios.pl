persona(juan, perez, [1, 1, 0, 0, 1, 0, 0, 1, 1, 0]).
persona(ana, gomez, [0, 1, 1, 0, 0, 1, 0, 1, 0, 1]).
persona(jose, rodriguez, [1, 0, 1, 0, 1, 1, 0, 0, 1, 1]).
persona(maria, lopez, [0, 0, 1, 1, 1, 0, 0, 1, 0, 1]).
persona(luis, fernandez, [1, 0, 0, 0, 1, 1, 0, 1, 0, 1]).
persona(ana, martinez, [0, 1, 0, 0, 1, 1, 0, 1, 1, 0]).
persona(juan, soto, [1, 0, 1, 1, 0, 0, 1, 0, 1, 0]).
persona(david, castillo, [0, 1, 1, 0, 1, 0, 1, 0, 0, 1]).
persona(sofia, garcia, [1, 1, 0, 0, 0, 1, 1, 0, 1, 0]).
persona(pepito, ramirez, [0, 0, 1, 1, 0, 1, 1, 0, 0, 1]).

conexiones(pepito, ana, 3).  
conexiones(ana, david, 2).
conexiones(pepito, juan).
conexiones(juan, david, 4).

ruta_corta(X, Y, (Ruta, Peso)) :-
    ruta_corta_aux(X, Y, [X], Ruta, 0, Peso).

ruta_corta_aux(X, X, Ruta, Ruta, Peso, Peso).
ruta_corta_aux(X, Y, Visitados, Ruta, PesoActual, Peso) :-
    conexiones(X, Z, P1),
    not(member(Z, Visitados)),
    PesoNuevo is PesoActual + P1,
    ruta_corta_aux(Z, Y, [Z | Visitados], Ruta, PesoNuevo, Peso).



muestra([1,0,1,0,0,1,1,1,0,1]).

similitud_cromosomas(C1, C2, Similitud) :-
    hamming_distance(C1, C2, Distancia),
    Similitud is (100 - Distancia * 10).

hamming_distance([], [], 0).
hamming_distance([X|Xs], [Y|Ys], Distancia) :-
    hamming_distance(Xs, Ys, RestoDistancia),
    Distancia is RestoDistancia + abs(X - Y).

encontrar_persona_similar(Persona, Porcentaje) :-
    muestra(Muestra),
    findall((Nombre, Similitud), (persona(Nombre, _, Cromosomas), similitud_cromosomas(Cromosomas, Muestra, Similitud)), ListaSimilitudes),
    max_member((Persona, Porcentaje), ListaSimilitudes).
