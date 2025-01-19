use rocket::{catch, catchers, Build, Rocket, Config};
use rocket::serde::json::{Value, json};

use crate::routes;

#[catch(404)]
fn not_found() -> Value {
    json!({
        "reason": "Resource was not found.",
        "success": false,
    })
}

#[catch(500)]
fn internal_error() -> Value {
    json!({
        "reason": "Internal server error.",
        "success": false,
    })
}
pub async fn server(config: Config) -> Rocket<Build> {
    let rocket = rocket::custom(config)
        .register("/", catchers![not_found, internal_error]);
    routes::mount(rocket)
}
