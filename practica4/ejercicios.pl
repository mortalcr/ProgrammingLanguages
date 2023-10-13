sumlist([], 0).
sumlist([X|Xs], S) :-
    sumlist(Xs, S1),
    S is S1 + X.

subconj([], _).
subconj([X|S1], S) :-
    member(X, S),
    subconj(S1, S).

aplanar([], []).
aplanar([X|Xs], L2) :-
    is_list(X),           
    aplanar(X, X1),       
    aplanar(Xs, Xs1),     
    append(X1, Xs1, L2).  
aplanar([X|Xs], [X|Xs1]) :-
    not(is_list(X)),     
    aplanar(Xs, Xs1).

partir([], _, [], []).
partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Resto, Umbral, Menores, Mayores).
partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Resto, Umbral, Menores, Mayores).

sub_cadenas(_, [], []).
sub_cadenas(Subcadena, [Cadena|Resto], Filtradas) :-
    sub_atom(Cadena, _, _, _, Subcadena), % Verifica si Subcadena es una subcadena de Cadena
    sub_cadenas(Subcadena, Resto, FiltradasResto),
    Filtradas = [Cadena|FiltradasResto].
sub_cadenas(Subcadena, [_|Resto], Filtradas) :-
    sub_cadenas(Subcadena, Resto, Filtradas).
