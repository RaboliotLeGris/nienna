use std::io::Error;

use crate::clients::s3::errors::S3ClientError;
use crate::worker_pool::jobs::job_errors::JobsError;

#[derive(Debug)]
#[allow(dead_code)]
pub enum VideoProcessorError {
    FailExtractMimetype,
    FailProcessVideo,
    FailExtractMiniature,
}

impl std::error::Error for VideoProcessorError {}

impl std::fmt::Display for VideoProcessorError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            VideoProcessorError::FailExtractMimetype => write!(f, "Fail to extract mimetype"),
            VideoProcessorError::FailProcessVideo => write!(f, "Fail to process video"),
            VideoProcessorError::FailExtractMiniature => write!(f, "Fail to extract miniature from video")
        }
    }
}

impl From<std::io::Error> for VideoProcessorError {
    fn from(e: Error) -> Self {
        match e {
            Error { .. } => VideoProcessorError::FailProcessVideo
        }
    }
}

impl From<JobsError> for VideoProcessorError {
    fn from(_e: JobsError) -> Self {
        VideoProcessorError::FailProcessVideo
    }
}

impl From<S3ClientError> for VideoProcessorError {
    fn from(_e: S3ClientError) -> Self {
        VideoProcessorError::FailProcessVideo
    }
}