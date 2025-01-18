use rocket::{routes, Build, Rocket};

mod hw;

pub fn mount(rocket: Rocket<Build>) -> Rocket<Build> {
    rocket.mount("/hello", routes![hw::hello_world])
}
