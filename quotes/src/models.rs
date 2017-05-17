use schema::posts;

#[derive(Queryable)]
pub struct Quote {
    pub id: i32,
    pub author: String,
    pub body: String
}

#[derive(Insertable)]
#[table_name="quotes"]
pub struct NewQuote<'a> {
    pub author: &'a str,
    pub body: &'a str,
}