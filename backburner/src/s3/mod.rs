use crate::s3::errors::S3ClientError;

#[cfg(test)]
pub mod client_stub;

pub mod client;
mod errors;

pub trait TS3Client: Sync + Send {
    /// get: fetch resources on the S3 and write it in the current folder
    fn get(&self, slug: String, filename: String) -> Result<(), S3ClientError>;
    // fn get_stream(self, slug: String) -> String;
    // fn put(&self, path: String, slug: String) -> String;
    // fn put_stream(self, slug: String) -> String;
}