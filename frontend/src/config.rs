use crate::api::greeter_client::GreeterClient;
use actix_web::web::{Data, ServiceConfig};
use log::info;
use tokio::sync::RwLock;
use tonic::transport::Channel;

pub struct HandlerConfig {
    pub name: String,
    pub client: RwLock<Option<GreeterClient<Channel>>>,
}

pub fn configure(cfg: &mut ServiceConfig, client: Option<GreeterClient<Channel>>) {
    info!("Configuring service");
    cfg.app_data(Data::new(HandlerConfig {
        name: "kube.direct".into(),
        client: RwLock::new(client),
    }));
}
