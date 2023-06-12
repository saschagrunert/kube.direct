use anyhow::Result;
use frontend::Server;

#[actix_web::main]
async fn main() -> Result<()> {
    Server::start().await
}
