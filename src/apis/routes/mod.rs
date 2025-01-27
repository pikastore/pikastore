use axum::{routing::get, Router};

mod s3;

pub fn mount() -> Router {
    Router::new()
        .route("/v1/bucket", get(|| async { s3::buckets::list_buckets().await }))
}