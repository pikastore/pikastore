use axum::{
    extract::FromRequest,
    body::Bytes,
    http::{header::CONTENT_TYPE, StatusCode, Request, HeaderMap, HeaderValue},
    response::{IntoResponse, Response},
};
use serde::de::DeserializeOwned;
use quick_xml::de::Deserializer;
use serde::de::Deserialize;
use serde::Serialize;


#[derive(Debug, Clone, Copy, Default)]
pub struct Xml<T>(pub T);

impl<T> IntoResponse for Xml<T> where T: Serialize,{
    fn into_response(self) -> axum::response::Response {
        match quick_xml::se::to_string(&self.0) {
            Ok(xml) => {
                let mut headers = HeaderMap::new();
                headers.insert(
                    CONTENT_TYPE,
                    HeaderValue::from_static("application/xml; charset=utf-8"),
                );
                (headers, xml).into_response()
            }
            Err(err) => {
                (
                    StatusCode::INTERNAL_SERVER_ERROR,
                    format!("Failed to serialize XML: {}", err),
                )
                    .into_response()
            }
        }
    }   
}

impl<T, S> FromRequest<S> for Xml<T> where T: DeserializeOwned, S: Send + Sync, {
    type Rejection = Response;

    async fn from_request(req: Request<axum::body::Body>, state: &S) -> Result<Self, Self::Rejection> {
        if !has_xml_content_type(req.headers()) {
            return Err(StatusCode::UNSUPPORTED_MEDIA_TYPE.into_response());
        }
        let bytes = Bytes::from_request(req, state)
            .await
            .map_err(|err| (StatusCode::BAD_REQUEST, err.to_string()).into_response())?;

             let mut deserializer = Deserializer::from_reader(&bytes[..]);
             let value = T::deserialize(&mut deserializer)
             .map_err(|err| {
                (
                    StatusCode::BAD_REQUEST,
                    format!("Failed to parse XML: {}", err),
                )
                    .into_response()
            })?;
        Ok(Xml(value))
    }
}

fn has_xml_content_type(headers: &HeaderMap) -> bool {
    let content_type = if let Some(content_type) = headers.get(CONTENT_TYPE) {
        content_type.to_str().unwrap_or("")
    } else {
        ""
    };

    content_type.starts_with("application/xml")
        || content_type.starts_with("text/xml")
        || content_type.ends_with("+xml")
}
