mod config;
mod db;
mod routes;
mod services;
mod test;
mod utils;

use ferris_says::say;
use serde::{Deserialize, Serialize};
use std::env;
use std::io::{stdout, BufWriter};

#[macro_use]
extern crate serde_derive;

#[derive(Serialize, Deserialize)]
struct User {
    id: Option<String>,
    name: String,
    email: String,
}

use lazy_static::lazy_static;
lazy_static! {
    static ref DB_URL: String = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
}

const OK_RESPONSE: &str = "HTTP/1.1 200 OK\r\nContent-type: application/json\r\n\r\n";
const NOT_FOUND: &str = "HTTP/1.1 404 NOT FOUND\r\n\r\n";
const INTERNAL_SERVER_ERROR: &str = "HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n";

fn main() {
    let stdout = stdout();
    let message = String::from("Hello fellow Rustaceans!");
    let width = message.chars().count();

    let mut writer = BufWriter::new(stdout.lock());
    say(&message, width, &mut writer).unwrap();
}

fn set_database() -> Resul<(), PostgresError> {
let mut client = Client:
}
