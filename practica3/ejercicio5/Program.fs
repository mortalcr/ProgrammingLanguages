open Laberinto

let rutaMasCorta = encontrarRutaMasCorta laberintoGrafo 1 7

match rutaMasCorta with
| Some ruta -> printfn "Ruta más corta: %A" ruta
| None -> printfn "No hay ruta"
