pub mod job_helpers;
pub mod job_errors;

pub type Job = Box<dyn FnOnce() + Send + 'static>;