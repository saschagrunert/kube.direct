use actix_web::web::{Data, ServiceConfig};
use log::info;

pub struct HandlerConfig;

pub fn configure(cfg: &mut ServiceConfig) {
    info!("Configuring service");
    cfg.app_data(Data::new(HandlerConfig));
}
