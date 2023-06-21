use crate::api::Data;
use anyhow::{bail, Context, Result};
use log::info;
use prost::Message;
use serde::{Deserialize, Serialize};
use std::rc::Rc;
use yew::prelude::*;

#[cfg(not(debug_assertions))]
pub const BACKEND_URL: &str = "http://backend.kube-direct";

#[cfg(debug_assertions)]
pub const BACKEND_URL: &str = "http://localhost:8081";

#[function_component]
pub fn App() -> Html {
    info!("Loading server side content");
    let fallback = html! { <div>{"Loading…"}</div> };

    html! {
        <Suspense {fallback}>
            <Content />
        </Suspense>
    }
}

#[derive(Deserialize, Serialize)]
enum ReturnData {
    Data(Data),
    Error(String),
}

#[function_component]
fn Content() -> HtmlResult {
    let state = use_prepared_state!(
        async move |_| -> ReturnData {
            match fetch_data().await.context("fetch data") {
                Err(e) => ReturnData::Error(format!("{:#}", e)),
                Ok(d) => ReturnData::Data(d),
            }
        },
        ()
    )?
    .unwrap_or(Rc::new(ReturnData::Error(
        "unable to use prepared state".into(),
    )));

    Ok(match state.as_ref() {
        ReturnData::Error(e) => html! {
            <div class="uk-alert-danger" uk-alert="">
                <p>{ "Unable to " } { e }</p>
            </div>
        },
        ReturnData::Data(d) => html! {
            <div class="uk-card uk-card-default uk-card-body uk-width-1-3@s uk-position-center">
                <h1 class="uk-card-title">{ "Welcome to kube.direct!" }</h1>
                <p>{ "The cluster is up and running if you see this website." }</p>
                <p>
                    { "Checkout the " }
                    <a href="https://github.com/saschagrunert/kube.direct" target="_blank">{ "sources on GitHub" }</a>
                    { "." }
                </p>
                <h4>{"Cluster Details"}</h4>
                <table class="uk-table uk-table-small uk-table-justify">
                    <tbody>
                        <tr>
                            <td>{ "Nodes" }</td>
                            <td>{{ d.nodes }}</td>
                        </tr>
                        <tr>
                            <td>{ "Kubernetes version" }</td>
                            <td>{{ &d.kubernetes_version }}</td>
                        </tr>
                        <tr>
                            <td>{ "OS image" }</td>
                            <td>{{ &d.os_image }}</td>
                        </tr>
                        <tr>
                            <td>{ "Kernel version" }</td>
                            <td>{{ &d.kernel_version }}</td>
                        </tr>
                        <tr>
                            <td>{ "Container runtime version" }</td>
                            <td>{{ &d.container_runtime_version }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        },
    })
}

async fn fetch_data() -> Result<Data> {
    info!("Fetching backend data");
    let resp = reqwest::get(BACKEND_URL)
        .await
        .context("get backend data")?;

    if !resp.status().is_success() {
        bail!(
            "bad backend status: {}",
            resp.text().await.context("get response text")?
        );
    }

    let bytes = resp.bytes().await.context("parse response bytes")?;

    Message::decode(bytes.as_ref()).context(format!("decode message: {:?}", bytes))
}
