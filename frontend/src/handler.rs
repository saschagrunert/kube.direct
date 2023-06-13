use crate::{config::HandlerConfig, error::Error, template::Index};
use actix_web::{web::Data, HttpRequest, Responder};
use actix_web_lab::respond::Html;
use anyhow::Context;
use askama::Template;
use log::{debug, info};
use serde::Deserialize;
use std::result::Result;

const BACKEND_URL: &str = "http://backend.kube-direct";

#[derive(Debug, Deserialize)]
pub struct BackendData {
    nodes: usize,
    kubernetes_version: String,
    os_image: String,
    kernel_version: String,
    container_runtime_version: String,
}

pub async fn index(req: HttpRequest, cfg: Data<HandlerConfig>) -> Result<impl Responder, Error> {
    debug!("{:#?}", req);

    let resp = reqwest::get(BACKEND_URL)
        .await
        .context("get backend data")?
        .json::<BackendData>()
        .await
        .context("parse JSON")?;
    info!("Got backend data: {:?}", resp);

    let index = Index {
        title: &cfg.title,
        nodes: resp.nodes,
        kubernetes_version: &resp.kubernetes_version,
        os_image: &resp.os_image,
        kernel_version: &resp.kernel_version,
        container_runtime_version: &resp.container_runtime_version,
    };

    Ok(Html(index.render().context("render index")?))
}
