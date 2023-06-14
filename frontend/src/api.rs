#[derive(::askama::Template)]
#[template(path = "index.html")]
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Data {
    #[prost(uint32, tag = "1")]
    pub nodes: u32,
    #[prost(string, tag = "2")]
    pub kubernetes_version: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub os_image: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub kernel_version: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub container_runtime_version: ::prost::alloc::string::String,
}
