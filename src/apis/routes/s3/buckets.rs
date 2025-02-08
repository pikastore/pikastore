use axum::{http::StatusCode, response::IntoResponse};
use serde::Serialize;
use services::utils::xml::Xml;

#[derive(Serialize)]
#[serde(rename = "ListAllMyBucketsResult")]
pub struct ListAllMyBucketsResult {
    #[serde(rename = "Buckets")]
    pub buckets: Buckets,
    #[serde(rename = "Owner")]
    pub owner: Owner,
}

#[derive(Serialize)]
pub struct Buckets {
    #[serde(rename = "Bucket")]
    pub bucket_list: Vec<Bucket>,
}

#[derive(Serialize)]
pub struct Bucket {
    #[serde(rename = "CreationDate")]
    pub creation_date: String,
    #[serde(rename = "Name")]
    pub name: String,
}

#[derive(Serialize)]
pub struct Owner {
    #[serde(rename = "DisplayName")]
    pub display_name: String,
    #[serde(rename = "ID")]
    pub id: String,
}

pub async fn list_buckets() -> impl IntoResponse {
    //fake data from the s3 docs 
    let response = ListAllMyBucketsResult {
        buckets: Buckets {
            bucket_list: vec![
                Bucket {
                    creation_date: "2019-12-11T23:32:47+00:00".to_string(),
                    name: "amzn-s3-demo-bucket".to_string(),
                },
                Bucket {
                    creation_date: "2019-11-10T23:32:13+00:00".to_string(),
                    name: "amzn-s3-demo-bucket1".to_string(),
                },
            ],
        },
        owner: Owner {
            display_name: "Account+Name".to_string(),
            id: "AIDACKCEVSQ6C2EXAMPLE".to_string(),
        },
    };

    (StatusCode::OK, Xml(response))
}
