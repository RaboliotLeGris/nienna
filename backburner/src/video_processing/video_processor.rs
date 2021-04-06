use std::process::Command;
use crate::video_processing::errors::VideoProcessorError;
use std::io::Read;

#[cfg(test)]
#[path = "./video_processor_tests.rs"]
mod video_processor_tests;

pub struct VideoProcessor {}

impl VideoProcessor {
    pub fn new() -> Self {
        VideoProcessor {}
    }

    /// Returns the video mimetype if possible
    ///
    /// Require `file` binary on the system
    pub fn get_mimetype(self, filepath: String) -> Result<String, VideoProcessorError> {
        let output = Command::new("file")
            .args(&["--mime-type", filepath.as_str()]).output()?.stdout;
        let mut parsed = String::new();
        output.as_slice().read_to_string(&mut parsed)?;
        if let Some(stripped_output) = parsed.strip_suffix("\n") {
            let collected_split: Vec<&str> = stripped_output.split_whitespace().collect();
            if let Some(mimetype) = collected_split.get(1) {
                return Ok(String::from(*mimetype));
            }
        }
        return Err(VideoProcessorError::FailExtractMimetype);
    }
}