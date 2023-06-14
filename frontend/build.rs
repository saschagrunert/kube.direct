use static_files::resource_dir;

fn main() -> std::io::Result<()> {
    #[cfg(feature = "gen-proto")]
    {
        let mut config = prost_build::Config::new();
        config.out_dir("src");
        config.type_attribute(".", "#[derive(::askama::Template)]");
        config.type_attribute(".", "#[template(path = \"index.html\")]");
        config.compile_protos(&["api.proto"], &["../api"])?;
    }

    resource_dir("./templates/static").build()
}
