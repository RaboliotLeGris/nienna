use std::io::Read;
use std::process::Command;

use crate::video_processing::errors::VideoProcessorError;
use std::env;

#[cfg(test)]
#[path = "./video_processor_tests.rs"]
mod video_processor_tests;

pub struct VideoProcessor {}

impl VideoProcessor {
    /// Returns the video mimetype if possible
    ///
    /// Require `file` binary on the system
    pub fn extract_mimetype(filepath: String) -> Result<String, VideoProcessorError> {
        let output = Command::new("file")
            .args(&["--mime-type", filepath.as_str()]).output()?.stdout;
        let mut parsed = String::new();
        output.as_slice().read_to_string(&mut parsed)?;
        if let Some(stripped_output) = parsed.strip_suffix("\n") {
            let collected_split: Vec<&str> = stripped_output.split_whitespace().collect();
            if let Some(mimetype) = collected_split.get(1) {
                if *mimetype == "cannot" {
                    return Err(VideoProcessorError::FailExtractMimetype);
                }
                return Ok(String::from(*mimetype));
            }
        }
        return Err(VideoProcessorError::FailExtractMimetype);
    }

    pub fn process(filepath: &String) -> Result<(), VideoProcessorError> {
        // ffmpeg -i .dev/samples/SampleVideo_1280x720_30mb.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls part.m3u8
        let output = Command::new("ffmpeg")
            .args(&["-i", filepath.as_str(), "-profile:v", "baseline", "-level", "3.0", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", "part.m3u8"])
            .output()?;
        if !output.status.success() {
            return Err(VideoProcessorError::FailProcessVideo);
        }
        Ok(())
    }

    pub fn extract_miniature(filepath: &String) -> Result<String, VideoProcessorError> {
        // ffmpeg -i input.mp4 -ss 00:00:01.000 -vframes 1 miniature.png
        let output = Command::new("ffmpeg")
            .args(&["-i", filepath.as_str(), "-ss", "00:00:01.000", "-vframes", "1", "miniature.jpeg"])
            .output()?;
        if !output.status.success() {
            return Err(VideoProcessorError::FailProcessVideo);
        }
        Ok(String::from("miniature.jpeg"))
    }
}