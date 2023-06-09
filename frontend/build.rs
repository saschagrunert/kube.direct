use static_files::resource_dir;

fn main() -> std::io::Result<()> {
    #[cfg(feature = "gen-proto")]
    tonic_build::configure()
        .out_dir("src")
        .build_server(false)
        .compile(&["api.proto"], &["../api"])?;
    resource_dir("./templates/static").build()
}
