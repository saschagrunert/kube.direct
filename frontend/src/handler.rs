use crate::{config::HandlerConfig, template::Index};
use actix_web::{web::Data, HttpRequest, Responder, Result};
use actix_web_lab::respond::Html;
use askama::Template;
use log::info;
use std::io::{self, ErrorKind};

pub async fn index(req: HttpRequest, _: Data<HandlerConfig>) -> Result<impl Responder> {
    info!("{:#?}", req);

    let html = Index
        .render()
        .map_err(|e| io::Error::new(ErrorKind::Other, e.to_string()))?;

    Ok(Html(html))
}
