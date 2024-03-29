use std::path::Path;
use crate::clients::s3::TS3Client;
use crate::clients::s3::errors::S3ClientError;

pub struct S3ClientStub {}

impl S3ClientStub {
    pub fn new() -> Self {
        S3ClientStub {}
    }
}

impl TS3Client for S3ClientStub {
    // a bit of hacking but we only want to copy test file into working directory
    fn get(&self, original_folder: &String, filename: &String) -> Result<(), S3ClientError> {
        let p = Path::new(&original_folder).join(Path::new(&filename));
        // TODO: change rsrc.mp4 to rsrc.EXTENSION
        std::fs::copy(p, std::env::current_dir().unwrap().join("rsrc.mp4")).unwrap();
        Ok(())
    }

    fn put(&self, path: &String, slug: &String, filename: &String) -> Result<(), S3ClientError> {
        unimplemented!()
    }
}