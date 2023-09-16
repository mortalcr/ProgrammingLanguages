let n_esimo n lista =
    let indices = [0..List.length lista - 1]
    let elementos = List.map2 (fun i x -> (i, x)) indices lista
    match List.tryFind (fun (i, _) -> i = n) elementos with
    | Some (_, elem) -> 
            printfn "%d" elem
    | None -> printfn "False"

let resultado1 = n_esimo 2 [1; 2; 3; 4; 5] // 3
let resultado2 = n_esimo 3 [1; 2; 3; 4; 5] // 4
let resultado3 = n_esimo 6 [1; 2; 3; 4; 5] // False
