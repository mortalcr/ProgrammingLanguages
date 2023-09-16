module Laberinto

type Habitacion = int
type Grafo = Map<Habitacion, Habitacion list>

// Representación del laberinto como un grafo
let laberintoGrafo : Grafo = 
    Map.ofList [
        (1, [4; 5]);
        (2, [1; 1]);
        (3, [3; 5]);
        (4, [2; 6]);
        (5, [5; 7]);
        (6, [4; 3]);
        (7, [5]);
        (8, [2]);
    ]

// Función para encontrar todas las rutas en el laberinto
let encontrarRutas grafo inicio fin =
    let rec dfs nodo ruta rutas visitados =
        if nodo = fin then
            List.rev (nodo::ruta) :: rutas
        else
            let vecinos = Map.find nodo grafo |> List.filter (fun n -> not (List.contains n visitados))
            List.collect (fun vecino -> dfs vecino (nodo::ruta) rutas (vecino::visitados)) vecinos

    dfs inicio [] [] [inicio]

// Función para encontrar la ruta más corta entre dos puntos
let encontrarRutaMasCorta grafo inicio fin =
    let rutas = encontrarRutas grafo inicio fin
    match rutas with
    | [] -> None // No hay ruta
    | _ ->
        let rutasOrdenadas = List.sortBy List.length rutas
        Some (List.head rutasOrdenadas)


