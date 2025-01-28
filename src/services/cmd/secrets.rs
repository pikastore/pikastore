use rand::RngCore;
use sha1::Digest;
use sha256::digest;

pub async fn generate_key_id() {
    let mut rng = rand::rng();
    let mut random_bytes = [0u8; 10];
    rng.fill_bytes(&mut random_bytes);
    let hash = hex::encode(sha1::Sha1::digest(&random_bytes));

    println!("Key ID: {}", hash);
}

pub async fn generate_access_key() {
    let mut rng = rand::rng();
    let mut random_bytes = [0u8; 10];
    rng.fill_bytes(&mut random_bytes);
    let hash = digest(&random_bytes);

    println!("Secret Access Key: {}", hash);
}
