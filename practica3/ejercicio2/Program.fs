let sub_cadenas (subcadena: string) (lista: string list) =
    lista
    |> List.filter (fun cadena -> cadena.Contains(subcadena))

// Ejemplo.
let resultado = sub_cadenas "la" ["la casa"; "el perro"; "pintando la cerca"]

printfn "%A" resultado
