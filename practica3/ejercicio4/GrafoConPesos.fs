module GrafoConPesos

open System.Collections.Generic

let grafoConPesos = [
    ("i", [("a", 3); ("b", 6)]);
    ("a", [("i", 3); ("c", 6); ("d", 5)]);
    ("b", [("i", 6); ("c", 2); ("d", 4)]);
    ("c", [("a", 6); ("b", 2); ("x", 7)]);
    ("d", [("a", 5); ("b", 4); ("f", 8)]);
    ("f", [("d", 8)]);
    ("x", [("c", 7)])
]

let rec dfsConPesos ini fin grafo =
    let rec dfsAux ruta pesoActual nodoActual =
        if List.isEmpty ruta then
            []
        elif nodoActual = fin then
            // Si llegamos al destino, retornamos la ruta
            [(List.rev ruta)]
        else
            // Extender la ruta con vecinos no visitados
            let vecinos =
                match List.tryFind (fun (nodo, _) -> nodo = nodoActual) grafo with
                | Some (_, vecinos) -> vecinos
                | None -> []
            
            let rutasExtendidas =
                List.filter (fun (vecino, _) -> not (List.contains vecino ruta)) vecinos
                |> List.collect (fun (vecino, pesoArista) ->
                    let nuevoPeso = pesoActual + pesoArista
                    let nuevaRuta = vecino :: ruta
                    dfsAux nuevaRuta nuevoPeso vecino)
            
            rutasExtendidas

    // Iniciar la búsqueda en profundidad desde el nodo inicial
    let rutas = dfsAux [ini] 0 ini

    // Encontrar la ruta más corta basándonos en la suma de pesos
    match rutas with
    | [] -> None // No hay ruta
    | _ ->
        let rutaMasCorta = List.minBy List.length rutas
        Some (List.rev rutaMasCorta) // Invertir la ruta para que sea del inicio al fin
