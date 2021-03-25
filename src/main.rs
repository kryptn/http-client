use reqwest;
use reqwest::header::{HeaderMap, HeaderName};
use serde::{Deserialize, Serialize};
use serde_json;
use std::collections::HashMap;
use std::error::Error;
use std::path::Path;

#[macro_use]
extern crate clap;

#[derive(Deserialize, Debug)]
struct Request {
    address: String,
    method: String,
    headers: HashMap<String, String>,
    data: serde_json::Value,
}

#[derive(Serialize, Debug)]
struct Response {
    // request: Request,
    status_code: u16,
    // headers: HashMap<String, String>
    data: serde_json::Value,
}

fn request_from_file<P: AsRef<Path>>(path: P) -> Result<Request, Box<dyn Error>> {
    let deserializer = serde_dhall::from_file(path);
    let thing: Result<Request, serde_dhall::Error> = deserializer.parse();
    match thing {
        Ok(req) => Ok(req),
        Err(err) => {
            eprintln!("{}", err);
            Err(Box::new(err))
        }
    }
}

fn run_req(req: &Request) -> Result<Response, Box<dyn Error>> {
    let client = reqwest::blocking::Client::new();

    let address = req.address.clone();

    let client = match req.method.to_lowercase().as_str() {
        "head" => client.head(address),
        "post" => client.post(address),
        "put" => client.put(address),
        "delete" => client.delete(address),
        "patch" => client.patch(address),
        "get" | _ => client.get(address),
    };

    let mut headers = HeaderMap::new();
    for (k, v) in req.headers.clone().into_iter() {
        let header_name = HeaderName::from_bytes(k.as_bytes()).unwrap();
        headers.insert(header_name, v.parse().unwrap());
    }

    let client = client.headers(headers);

    let client = client.json(&req.data);

    let res = client.send()?;

    // println!("{:#?}", res);
    // println!("{:?}", res.bytes()?);

    let resp = Response {
        // request: req,
        status_code: res.status().as_u16(),
        data: res.json()?,
    };

    // let res = client.post(req.address)
    // //.headers(req.headers)
    // .json(&req.data)
    // .send()?;

    Ok(resp)
}

fn store_result<P: AsRef<Path>>(
    _req: &Request,
    resp: &Response,
    path: P,
) -> Result<(), Box<dyn std::error::Error>> {
    // eprintln!("would try to store now");

    let path = std::fs::canonicalize(path)?;
    let filename = path.file_name().unwrap();
    let mut p = path.parent().unwrap().to_path_buf();
    p.push("resp");

    match std::fs::create_dir_all(&p) {
        Ok(..) => eprintln!("created dir {:?}", "resp_root"),
        Err(error) => eprintln!("error on create dir {}", error),
    }

    match serde_dhall::serialize(&resp.data).to_string() {
        Ok(o) => {
            p.push(filename);
            eprintln!("trying to output to {:?}", p);
            std::fs::write(p, o)?;
        }
        Err(error) => eprintln!("Error on write: {}", error),
    }
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let matches = clap_app!(myapp =>
        (version: crate_version!())
        (author: "David Bibb <kryptn@gmail.com>")
        (about: "CLI request tool")
        (@arg CONFIG: -c --config +takes_value "Sets a custom config file")
        (@arg debug: -d ... "Sets the level of debugging information")
        (@subcommand req =>
            (about: "runs a request")
            (@arg FILENAME: +required "input file")
            (@arg verbose: -v --verbose "Print test information verbosely")
        )
        (@subcommand resp =>
            (about: "returns the response from the last request")
            (@arg FILENAME: +required "input file")
            (@arg verbose: -v --verbose "Print test information verbosely")
        )
    )
    .get_matches();

    match matches.subcommand() {
        ("req", Some(sub_match)) => {
            //println!("matched req");
            let filename = sub_match.value_of("FILENAME").unwrap();
            let request = request_from_file(filename)?;
            let response = run_req(&request)?;

            store_result(&request, &response, filename)?;

            println!("{}", serde_json::to_string_pretty(&response.data).unwrap());
        }
        ("resp", Some(sub_match)) => {
            //println!("matched resp");
            let filename = sub_match.value_of("FILENAME").unwrap();

            let resp_path = Path::new("./resp").join(filename);

            let value: serde_json::Value = serde_dhall::from_file(resp_path).parse()?;
            println!("{}", serde_json::to_string_pretty(&value).unwrap());

            //println!("filename: {:?}", filename);
        }
        _ => {
            println!("no command found");
        }
    }

    Ok(())

    // println!("{:#?}", matches)
}
