use std::fs::File;

use rust_s3::bucket::Bucket;
use rust_s3::creds::Credentials;
use rust_s3::region::Region;

use crate::clients::s3::errors::S3ClientError;
use crate::clients::s3::TS3Client;

#[cfg(test)]
#[path = "./client_tests.rs"]
mod client_tests;

#[derive(Clone)]
pub struct S3Client {
    region: Region,
    credentials: Credentials,
    bucket: String,
}

impl S3Client {
    pub fn new(endpoint: String, bucket_name: String, access_key: String, secret_key: String) -> Self {
        S3Client {
            region: Region::Custom {
                region: "eu-west-1".into(),
                endpoint,
            },
            credentials: Credentials {
                access_key: Some(access_key),
                secret_key: Some(secret_key),
                security_token: None,
                session_token: None,
            },
            bucket: bucket_name,
        }
    }
}

impl TS3Client for S3Client {
    fn get(&self, slug: &String, filename: &String) -> Result<(), S3ClientError> {
        let bucket = Bucket::new_with_path_style("nienna-1", self.region.clone(), self.credentials.clone())?;
        let mut file = File::create(filename).unwrap();
        let code = bucket.get_object_stream_blocking(format!("{}/{}", slug, filename).as_str(), &mut file)?;
        if code != 200 {
            return Err(S3ClientError::FailFetchResource);
        }
        Ok(())
    }

    fn put(&self, path: &String, slug: &String, filename: &String) -> Result<(), S3ClientError> {
        let bucket = Bucket::new_with_path_style("nienna-1", self.region.clone(), self.credentials.clone())?;
        bucket.put_object_stream_blocking(path, format!("{}/{}", slug, filename))?;
        Ok(())
    }
}