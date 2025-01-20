/* ---ARGS---
    donate
    serve (additional args: --port <i64> --host <String> ) 
    secrets (additional args: generate ) (write these down you wont be able to see them again)
    create bucket additional args: -name <String> )
*/
use clap::{Parser, Subcommand};
use rocket::{tokio, Config};
use apis::server::server;
use services::cmd::secrets;
#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
pub struct Args {
    #[command(subcommand)]
    pub command: Commands,
}
#[derive(Subcommand, Debug)]
pub enum Commands {
    Serve {
        #[arg(long, default_value = "8000")]
        port: u16,
        #[arg(long, default_value = "127.0.0.1")]
        host: String,
    },
    Secrets {
        #[arg(long)]
        generate: bool,
    },
    Bucket {
        #[arg(long)]
        name: String,
    },
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = Args::parse();

    match &args.command {
        Commands::Serve { host, port } => {
            let config = Config {
                address: host.parse()?,
                port: *port, 
                ..Config::default()
            };
            server(config).await.launch().await?;
        }
        Commands::Secrets { generate } => {
            if *generate {
                secrets::generate_key_id().await;
                secrets::generate_access_key().await;
            }
        }
        Commands::Bucket { name } => {
            println!("Creating bucket with name: {}", name);
        }
    }

    Ok(())
}
