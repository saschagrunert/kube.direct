#[cfg(not(feature = "gen-proto"))]
pub fn main() {}

#[cfg(feature = "gen-proto")]
fn main() -> std::io::Result<()> {
    let mut config = prost_build::Config::new();
    config.out_dir("src");
    config.type_attribute(".", "#[derive(serde::Deserialize, serde::Serialize)]");
    config.compile_protos(&["api.proto"], &["../api"])
}
