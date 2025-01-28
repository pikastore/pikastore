use axum::response::IntoResponse;
use serde::Serialize;

#[derive(Debug, Clone, Copy, Default)]
pub struct Xml<T>(pub T);

impl<T> IntoResponse for Xml<T> where T: Serialize,{
    fn into_response(self) -> axum::response::Response {
        match quick_xml::se::to_string(&self.0) {
            Ok(Xml) => {
                
            }
        }
    }   
}
