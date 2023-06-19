use anyhow::{Context, Result};
use frontend::Server;

#[tokio::main]
async fn main() -> Result<()> {
    Server::run().await.context("run server")
}
