struct Server {
  pub  ServerName: String,
  pub   KeyID: String,
  pub   AccessKey: String, 
  pub   Region: String
}

struct Bucket {
  pub  BucketName: String,
  pub CreatedAt: std::time::SystemTime,
}

struct Object {

}