#[derive(Debug)]
pub enum S3ClientError{
    FailFetchResource,
}

impl std::error::Error for S3ClientError {}

impl std::fmt::Display for S3ClientError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            S3ClientError::FailFetchResource => write!(f, "Failed fetch resources from object storage"),
        }
    }
}