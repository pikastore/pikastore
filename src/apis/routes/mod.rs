use rocket::{routes, Build, Rocket};

mod s3;

pub fn mount(rocket: Rocket<Build>) -> Rocket<Build> {
    rocket.mount("/v1/bucket", routes![s3::bucket::list_buckets])
}
