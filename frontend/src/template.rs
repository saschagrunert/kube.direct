use askama::Template;

#[derive(Template, Default)]
#[template(path = "index.html")]
pub struct Index<'a> {
    pub title: &'a str,
    pub nodes: u8,
    pub kubernetes_version: &'a str,
}
