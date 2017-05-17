extern crate rustweb;
extern crate diesel;

use self::rustweb::*;
use self::rustweb::models::*;
use self::diesel::prelude::*;

fn main() {
    use rustweb::schema::quotes::dsl::*;

    let connection = establish_connection();
    let results = quotes
        .limit(20)
        .load::<Quote>(&connection)
        .expect("Error loading quotes");

    println!("Displaying the {} most recent quotes", results.len());
    for quote in results {
        println!("==========================================================");
        println!("Author: {}", quote.author);
        println!("==========================================================\n");
        println!("{}", quote.body);
    }
}