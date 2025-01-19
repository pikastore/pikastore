use rocket::get;

#[get("/")]
pub async fn list_buckets() -> &'static str {
    "List of buckets"
}