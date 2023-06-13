use askama::Template;

#[derive(Template, Default)]
#[template(path = "index.html")]
pub struct Index<'a> {
    pub title: &'a str,
    pub nodes: usize,
    pub kubernetes_version: &'a str,
    pub os_image: &'a str,
    pub kernel_version: &'a str,
    pub container_runtime_version: &'a str,
}
