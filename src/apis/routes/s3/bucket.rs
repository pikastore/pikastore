use rocket::{get, post};

#[get("/")]
pub async fn list_buckets() -> &'static str {
    ""
}

#[post("/")]
pub async fn create_bucket() -> &'static str {
    ""
}