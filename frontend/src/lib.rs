pub use app::App;

mod api;
mod app;

#[cfg(not(target_arch = "wasm32"))]
pub use server::Server;

#[cfg(not(target_arch = "wasm32"))]
mod server;
