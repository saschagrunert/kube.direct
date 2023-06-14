use crate::{
    config,
    handler::{self, BACKEND_URL},
};
use actix_web::{middleware::Logger, web, App, HttpResponse, HttpServer};
use actix_web_static_files::ResourceFiles;
use anyhow::{Context, Result};
use env_logger::{Builder, Env};
use log::info;
use std::env;

include!(concat!(env!("OUT_DIR"), "/generated.rs"));

#[derive(Debug)]
/// The main server structure.
pub struct Server;

impl Server {
    /// Start the Server
    pub async fn start() -> Result<()> {
        Builder::from_env(Env::default().default_filter_or("info")).init();

        let port = match env::var("PORT") {
            Ok(v) => v.parse::<u16>().context("parse PORT env variable")?,
            Err(_) => 8080,
        };

        #[cfg(feature = "local")]
        info!(
            "Running in local mode, expecting backend on {}",
            BACKEND_URL
        );

        info!("Serving on http://localhost:{}", port);
        HttpServer::new(|| {
            App::new()
                .wrap(Logger::default())
                .service(ResourceFiles::new("/static", generate()))
                .configure(config::configure)
                .route("/", web::get().to(handler::index))
                .route("/", web::post().to(handler::index))
                .route(
                    "/health/{_:(readiness|liveness)}",
                    web::get().to(HttpResponse::Ok),
                )
        })
        .bind(("0.0.0.0", port))
        .context("bind server IP and port")?
        .workers(1)
        .run()
        .await
        .context("run http server")
    }
}
