let withVariables = \(address : Text) -> \(document : Text) -> \(variablesType : Type) -> \(variables : variablesType) -> {
    method = "POST",
    headers = {=},
    address = address,
    data = {
        query = document,
        variables = variables,
    }
}

let query = \(address : Text) -> \(document : Text) -> {
    method = "POST",
    headers = {=},
    address = address,
    data = {
        query = document,
    }
}

in {
    queryWithVariables = withVariables,
    query = query,
}