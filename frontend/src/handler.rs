use crate::{config::HandlerConfig, template::Index};
use actix_web::{web::Data, HttpRequest, Responder, Result};
use actix_web_lab::respond::Html;
use askama::Template;
use log::debug;
use std::io::{Error as IOError, ErrorKind};

macro_rules! io_err {
    ($x:expr) => {
        $x.map_err(|e| IOError::new(ErrorKind::Other, e.to_string()))?
    };
}

pub async fn index(req: HttpRequest, cfg: Data<HandlerConfig>) -> Result<impl Responder> {
    debug!("{:#?}", req);

    let html = io_err!(Index { name: &cfg.name }.render());

    Ok(Html(html))
}
