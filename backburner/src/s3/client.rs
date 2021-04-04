use crate::s3::TS3Client;
use crate::s3::errors::S3ClientError;

pub struct S3Client {

}

impl S3Client {
    pub fn new() -> Self {
        S3Client{}
    }
}

impl TS3Client for S3Client {
    fn get(&self, slug: String, filename: String) -> Result<(), S3ClientError> {
        todo!()
    }
}