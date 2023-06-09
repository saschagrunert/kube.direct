use crate::api::greeter_client::GreeterClient;
use actix_web::{middleware::Logger, web, App, HttpResponse, HttpServer};
use actix_web_static_files::ResourceFiles;
use anyhow::{Context, Result};
use env_logger::{Builder, Env};
use log::info;
use std::env;

mod api;
mod config;
mod handler;
mod template;

include!(concat!(env!("OUT_DIR"), "/generated.rs"));

#[actix_web::main]
async fn main() -> Result<()> {
    Builder::from_env(Env::default().default_filter_or("info")).init();

    let port = match env::var("PORT") {
        Ok(v) => v.parse::<u16>().context("parse PORT env variable")?,
        Err(_) => 8080,
    };

    info!("Connecting to gRPC server");
    let client = GreeterClient::connect("http://backend:50051")
        .await
        .context("connect to gRPC server")?;

    info!("Starting web server");
    HttpServer::new(move || {
        App::new()
            .wrap(Logger::default())
            .service(ResourceFiles::new("/static", generate()))
            .configure(|sc| config::configure(sc, client.clone()))
            .route("/", web::get().to(handler::index))
            .route("/", web::post().to(handler::index))
            .route(
                "/health/{_:(readiness|liveness)}",
                web::get().to(HttpResponse::Ok),
            )
    })
    .bind(("0.0.0.0", port))?
    .workers(1)
    .run()
    .await
    .context("run http server")
}
