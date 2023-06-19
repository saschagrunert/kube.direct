use crate::App;
use anyhow::{Context, Result};
use bytes::Bytes;
use futures::stream::{self, Stream, StreamExt};
use std::{env, error::Error, path::PathBuf};
use warp::{self, Filter};
use yew::ServerRenderer;

type BoxedError = Box<dyn Error + Send + Sync + 'static>;

pub struct Server;

impl Server {
    pub async fn run() -> Result<()> {
        const TARGET: &str = "server";
        env::set_var("RUST_LOG", TARGET);
        env_logger::init();

        const BODY: &str = "<body>";
        let dir = PathBuf::from("public");

        let index_html_s = tokio::fs::read_to_string(dir.join("index.html"))
            .await
            .context("read index")?;

        let (index_html_before, index_html_after) =
            index_html_s.split_once(BODY).context("split index")?;
        let mut index_html_before = index_html_before.to_owned();
        index_html_before.push_str(BODY);
        let index_html_after = index_html_after.to_owned();

        let html = warp::path::end().then(move || {
            let index_html_before = index_html_before.clone();
            let index_html_after = index_html_after.clone();

            async move { warp::reply::html(Self::render(index_html_before, index_html_after).await) }
        });

        let readiness = warp::path!("health" / "readiness").map(warp::reply);
        let liveness = warp::path!("health" / "liveness").map(warp::reply);
        let log = warp::log(TARGET);
        let routes = html
            .or(warp::fs::dir(dir))
            .or(readiness)
            .or(liveness)
            .with(log);

        println!("Serving at: http://localhost:8080/");
        warp::serve(routes).run(([0, 0, 0, 0], 8080)).await;

        Ok(())
    }

    async fn render(
        index_html_before: String,
        index_html_after: String,
    ) -> Box<dyn Stream<Item = Result<Bytes, BoxedError>> + Send> {
        let renderer = ServerRenderer::<App>::new();

        Box::new(
            stream::once(async move { index_html_before })
                .chain(renderer.render_stream())
                .chain(stream::once(async move { index_html_after }))
                .map(|m| Result::<_, BoxedError>::Ok(m.into())),
        )
    }
}
