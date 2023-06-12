use actix_web::web::{Data, ServiceConfig};
use log::info;

#[derive(Clone)]
pub struct HandlerConfig {
    pub title: String,
}

impl Default for HandlerConfig {
    fn default() -> HandlerConfig {
        HandlerConfig {
            title: "kube.direct".into(),
        }
    }
}

pub fn configure(cfg: &mut ServiceConfig) {
    info!("Configuring service");
    cfg.app_data(Data::new(HandlerConfig::default()));
}
