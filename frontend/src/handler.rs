use crate::{api::Data, config::HandlerConfig, error::Error};
use actix_web::{web, HttpRequest, Responder};
use actix_web_lab::respond::Html;
use anyhow::{format_err, Context};
use askama::Template;
use log::{debug, info};
use prost::Message;
use std::result::Result;

#[cfg(not(feature = "local"))]
pub const BACKEND_URL: &str = "http://backend.kube-direct";

#[cfg(feature = "local")]
pub const BACKEND_URL: &str = "http://localhost:8081";

pub async fn index(req: HttpRequest, _: web::Data<HandlerConfig>) -> Result<impl Responder, Error> {
    debug!("{:#?}", req);

    let resp = reqwest::get(BACKEND_URL)
        .await
        .context("get backend data")?;

    if !resp.status().is_success() {
        return Err(format_err!(
            "bad backend status: {}",
            resp.text().await.context("get response text")?
        )
        .into());
    }

    let bytes = resp.bytes().await.context("parse response bytes")?;

    let data: Data =
        Message::decode(bytes.as_ref()).context(format!("decode message: {:?}", bytes))?;

    info!("Got backend data: {:?}", data);
    Ok(Html(data.render().context("render index")?))
}
