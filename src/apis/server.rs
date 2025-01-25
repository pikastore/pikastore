use axum::{response::Redirect, routing::get, Router};
pub async fn server(host: String, port: u16) {
    let addr = format!("{}:{}", host, port);
    //NETTSPENDDDDD
    let app = Router::new().route("/", get(|| async {Redirect::permanent("https://open.spotify.com/album/2j74DNrJ8TgnMEukERqnnm?si=JDnC66ORSJG6Seu1d3a5Tw") }));
    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();
    println!("🦀Server started on {}:{}", host, port);
    axum::serve(listener, app).await.unwrap();
}