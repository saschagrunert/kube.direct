use crate::{api::Data, config::HandlerConfig, error::Error};
use actix_web::{web, HttpRequest, Responder};
use actix_web_lab::respond::Html;
use anyhow::Context;
use askama::Template;
use log::{debug, info};
use prost::Message;
use std::result::Result;

const BACKEND_URL: &str = "http://backend.kube-direct";

pub async fn index(req: HttpRequest, _: web::Data<HandlerConfig>) -> Result<impl Responder, Error> {
    debug!("{:#?}", req);

    let resp = reqwest::get(BACKEND_URL)
        .await
        .context("get backend data")?
        .bytes()
        .await
        .context("parse bytes")?;

    let data: Data = Message::decode(resp).context("decode message")?;
    info!("Got backend data: {:?}", data);

    Ok(Html(data.render().context("render index")?))
}
