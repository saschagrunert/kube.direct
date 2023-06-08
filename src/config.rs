use actix_web::web::{Data, ServiceConfig};
use log::info;

#[derive(Clone)]
pub struct HandlerConfig {
    pub name: String,
}

impl Default for HandlerConfig {
    fn default() -> HandlerConfig {
        HandlerConfig {
            name: "world".into(),
        }
    }
}

pub fn configure(cfg: &mut ServiceConfig) {
    info!("Configuring service");
    cfg.app_data(Data::new(HandlerConfig::default()));
}
