let rec desplazar direccion n lista =
    // Función para desplazar a la izquierda
    let desplazarIzquierda n lista =
        List.append (List.skip n lista) (List.replicate n 0)

    // Función para desplazar a la derecha
    let desplazarDerecha n lista =
        List.append (List.replicate n 0) (List.take (List.length lista - n) lista)

    match direccion with
    | "izq" -> desplazarIzquierda n lista
    | "der" -> desplazarDerecha n lista
    | _ -> failwith "Dirección no válida"

// Ejemplos
let resultado1 = desplazar "izq" 3 [1; 2; 3; 4; 5] // [4; 5; 0; 0; 0]
let resultado2 = desplazar "der" 2 [1; 2; 3; 4; 5] // [0; 0; 1; 2; 3]
let resultado3 = desplazar "izq" 5 [1; 2; 3; 4; 5] // [0; 0; 0; 0; 0]

printfn "%A" resultado1
printfn "%A" resultado2
printfn "%A" resultado3
