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
    pub fn extract_mimetype(self, filepath: String) -> Result<String, VideoProcessorError> {
        let output = Command::new("file")
            .args(&["--mime-type", filepath.as_str()]).output()?.stdout;
        let mut parsed = String::new();
        output.as_slice().read_to_string(&mut parsed)?;
        if let Some(stripped_output) = parsed.strip_suffix("\n") {
            let collected_split: Vec<&str> = stripped_output.split_whitespace().collect();
            if let Some(mimetype) = collected_split.get(1) {
                if *mimetype == "cannot" {
                    return Err(VideoProcessorError::FailExtractMimetype)
                }
                return Ok(String::from(*mimetype));
            }
        }
        return Err(VideoProcessorError::FailExtractMimetype);
    }

    pub fn process(self, filepath: String) -> Result<(), VideoProcessorError> {
        /*
        Something like that sort of works:
        ``` bash
            ffmpeg -i filepath \
            -filter_complex "[0:v:0]split=2[split1][split2];[split2]scale=width=-2:height=432:flags=fast_bilinear[scale2]" \
            -codec:v "libx264" -crf:v 23 -profile:v "high" -pix_fmt:v "yuv420p" -force_key_frames:v expr:'gte(t,n_forced*2.000)' -preset:v "faster" -b-pyramid:v "strict" \
            -map [split1] \
            -map [scale2] \
            -codec:a aac -ac:a 2 -b:a 96000 \
            -map 0:a:0 \
            -map 0:a:0 \
            -f hls \
            -hls_time 6 \
            -hls_playlist_type "vod" \
            -hls_segment_type "mpegts" \
            -hls_segment_filename '%v_%05d.ts' \
            -master_pl_name "master.m3u8" \
            -var_stream_map "v:1,a:1 v:0,a:0" "%v.m3u8"
        ```
         */


        Ok(())
    }
}