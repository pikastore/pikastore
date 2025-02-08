use crate::routes;
pub async fn server(host: String, port: u16) {
    let addr = format!("{}:{}", host, port);
    let app = routes::mount();
    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();
    println!("🦀 Server started on {}:{}", host, port);
    axum::serve(listener, app).await.unwrap();
}