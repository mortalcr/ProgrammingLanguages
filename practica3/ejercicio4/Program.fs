open Recipientes
open RutaCorta
open GrafoConPesos

let pasosRecipientes = prof [2; 2] [2; 0]
printfn "%A" pasosRecipientes

let ruta = prof2 "i" "f" grafo
printfn "%A" ruta

let result = dfsConPesos "i" "x" grafoConPesos

match result with
| Some ruta2 -> printfn "Ruta más corta: %A" ruta2
| None -> printfn "No hay ruta"

0 // Salida del programa

