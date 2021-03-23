let Map = https://prelude.dhall-lang.org/v15.0.0/Map/Type

let withVariables = \(address : Text) -> \(headers : Map Text Text) -> \(document : Text) -> \(variablesType : Type) -> \(variables : variablesType) -> {
    method = "POST",
    headers = headers,
    address = address,
    data = {
        query = document,
        variables = variables,
    }
}

let query = \(address : Text) -> \(headers : Map Text Text) -> \(document : Text) -> {
    method = "POST",
    headers = headers,
    address = address,
    data = {
        query = document,
    }
}

in {
    queryWithVariables = withVariables,
    query = query,
}