use rust_s3::S3Error;

#[derive(Debug)]
pub enum S3ClientError{
    FailFetchResource,
    FailWriteResource,
}

impl std::error::Error for S3ClientError {}

impl std::fmt::Display for S3ClientError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            S3ClientError::FailFetchResource => write!(f, "Failed fetch resources from object storage"),
            S3ClientError::FailWriteResource => write!(f, "Failed to write resources"),
        }
    }
}

impl From<S3Error> for S3ClientError {
    fn from(e: S3Error) -> Self {
        match e {
            _ => S3ClientError::FailFetchResource
        }
    }
}

impl From<std::io::Error> for S3ClientError {
    fn from(e: std::io::Error) -> Self {
        match e {
           _  => S3ClientError::FailWriteResource
        }
    }
}