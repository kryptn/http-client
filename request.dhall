let Map = https://prelude.dhall-lang.org/v15.0.0/Map/Type


let post = \(address : Text) -> \(headers : Map Text Text) -> \(payloadType : Type) -> \(payload : payloadType) -> {
    address = address,
    method = "POST",
    data = payload,
    headers = headers,
    extra = "words",
}

in {
    post = post
}