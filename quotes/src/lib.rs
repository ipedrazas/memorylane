#![feature(proc_macro)]

#[macro_use] extern crate diesel;
#[macro_use] extern crate diesel_codegen;
extern crate dotenv;

pub mod schema;
pub mod models;

use diesel::prelude::*;
use diesel::pg::PgConnection;
use dotenv::dotenv;
use std::env;

use self::models::{Quote, NewQuote};

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .expect(&format!("Error connecting to {}", database_url))
}

pub fn create_quote(conn: &PgConnection, author: &str, body: &str) -> Quote {
    use schema::quotes;

    let new_quote = NewQuote {
        author: author,
        body: body,
    };

    diesel::insert(&new_quote).into(quotes::table)
        .get_result(conn)
        .expect("Error saving new quote")
}