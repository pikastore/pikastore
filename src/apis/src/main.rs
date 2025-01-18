use rocket::{launch, Build, Rocket};
mod routes;

pub async fn server() -> Rocket<Build> {
    let rocket = rocket::build();
    routes::mount(rocket)
}

#[launch]
async fn rocket() -> _ {
    server().await
}