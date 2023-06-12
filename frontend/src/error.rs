use actix_web::error::ResponseError;
use anyhow::Error as AnyhowError;
use std::fmt::{Display, Formatter, Result};

#[derive(Debug)]
pub struct Error {
    err: AnyhowError,
}

impl Display for Error {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        self.err.fmt(f)
    }
}

impl ResponseError for Error {}

impl<T> From<T> for Error
where
    T: Into<AnyhowError>,
{
    fn from(e: T) -> Self {
        Error { err: e.into() }
    }
}
