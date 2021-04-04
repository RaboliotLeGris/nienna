use crate::s3::TS3Client;
use crate::s3::errors::S3ClientError;

pub struct S3ClientStub {

}

impl S3ClientStub {
    pub fn new() -> Self {
        S3ClientStub{}
    }
}

impl TS3Client for S3ClientStub {
    fn get(&self, slug: String, filename: String) ->  Result<(), S3ClientError> {
        todo!()
    }
}