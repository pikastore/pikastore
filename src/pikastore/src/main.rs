/* ---ARGS---
    serve (additional args: -port <i64> --host <String> )
    
    create bucket (additional args: -name <String> )
*/
use clap::Parser;
use rocket::{tokio, Config};
use apis::server::server;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
  #[arg(short, long)]
   port: u16,
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = Args::parse();
    let config = Config {
        port: args.port,
        ..Config::default()
    };

    server(config).await.launch().await?;
 Ok(())
}