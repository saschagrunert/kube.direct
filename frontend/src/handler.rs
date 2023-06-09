use crate::{config::HandlerConfig, template::Index};
use actix_web::{web::Data, HttpRequest, Responder, Result};
use actix_web_lab::respond::Html;
use askama::Template;
use log::debug;
use std::io::{self, ErrorKind};

pub async fn index(req: HttpRequest, _: Data<HandlerConfig>) -> Result<impl Responder> {
    debug!("{:#?}", req);

    let html = Index {
        name: "kube.direct",
    }
    .render()
    .map_err(|e| io::Error::new(ErrorKind::Other, e.to_string()))?;

    Ok(Html(html))
}
