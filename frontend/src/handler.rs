use crate::{api::HelloRequest, config::HandlerConfig, template::Index};
use actix_web::{web::Data, HttpRequest, Responder, Result};
use actix_web_lab::respond::Html;
use askama::Template;
use log::{debug, info};
use std::io::{Error as IOError, ErrorKind};
use tonic::Request;

macro_rules! io_err {
    ($x:expr) => {
        $x.map_err(|e| IOError::new(ErrorKind::Other, e.to_string()))?
    };
}

pub async fn index(req: HttpRequest, cfg: Data<HandlerConfig>) -> Result<impl Responder> {
    debug!("Got request: {:#?}", req);

    let request = Request::new(HelloRequest { name: "foo".into() });
    let mut client = cfg.client.write().await;
    if let Some(c) = client.as_mut() {
        let response = io_err!(c.say_hello(request).await);
        info!("Got response {:?}", response);
    }

    let html = io_err!(Index { name: &cfg.name }.render());
    Ok(Html(html))
}
