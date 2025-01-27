use axum::Json;
use serde_json::{Value, json};
pub async fn list_buckets() -> Json<Value> {
    Json(json!({ "message": "idk"}))
}
