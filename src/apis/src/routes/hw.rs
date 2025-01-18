use rocket::get;

#[get("/")]
pub fn hello_world() -> &'static str {
    "Hello, world!"
}