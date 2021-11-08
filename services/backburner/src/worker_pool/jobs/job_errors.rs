use std::io::Error;

#[derive(Debug)]
pub enum JobsError {
    FailCreateWorkingFolder,
}

impl std::error::Error for JobsError {}

impl std::fmt::Display for JobsError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            JobsError::FailCreateWorkingFolder => write!(f, "Failed to creating working folder"),
        }
    }
}

impl From<std::io::Error> for JobsError {
    fn from(_: Error) -> Self {
        JobsError::FailCreateWorkingFolder
    }
}
