fn main() {
    #[cfg(target_arch = "wasm32")]
    {
        wasm_logger::init(wasm_logger::Config::default());
        yew::Renderer::<frontend::App>::new().hydrate();
    }
}
