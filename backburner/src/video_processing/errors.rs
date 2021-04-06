use std::io::Error;

#[derive(Debug)]
pub enum VideoProcessorError{
    FailExtractMimetype,
}

impl std::error::Error for VideoProcessorError {}

impl std::fmt::Display for VideoProcessorError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            VideoProcessorError::FailExtractMimetype => write!(f, "Fail to extract mimetype"),
        }
    }
}

impl From<std::io::Error> for VideoProcessorError {
    fn from(e: Error) -> Self {
        match e {
            Error { .. } => VideoProcessorError::FailExtractMimetype
        }
    }
}