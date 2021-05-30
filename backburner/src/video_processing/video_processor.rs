use std::io::Read;
use std::process::Command;

use crate::video_processing::errors::VideoProcessorError;
use std::num::ParseIntError;

#[cfg(test)]
#[path = "./video_processor_tests.rs"]
mod video_processor_tests;

type Resolution = (u32, u32);

pub struct VideoProcessor {}

impl VideoProcessor {
    /// Returns the video mimetype if possible
    ///
    /// Require `file` binary on the system
    #[allow(dead_code)]
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
        Err(VideoProcessorError::FailExtractMimetype)
    }

    pub fn process(filepath: &String) -> Result<(), VideoProcessorError> {
        let resolution = VideoProcessor::extract_resolution(filepath)?;
        let cmd_args = VideoProcessor::generate_command(filepath, &resolution)?;

        debug!("Video processing generated command: {:?}\n\n\n", cmd_args);

        let output = Command::new("ffmpeg")
            .args(&cmd_args)
            .output()?;
        if !output.status.success() {
            return Err(VideoProcessorError::FailProcessVideo);
        }
        Ok(())
    }

    fn extract_resolution(filename: &String) -> Result<Resolution, VideoProcessorError> {
        let output = Command::new("ffprobe")
            .args(&["-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=s=x:p=0", filename.as_str()])
            .output()?;
        if !output.status.success() {
            return Err(VideoProcessorError::FailProcessVideo);
        }
        if let Ok(raw_output) = String::from_utf8(output.stdout) {
            let split = raw_output.trim().split("x");
            let vec: Vec<&str> = split.collect();

            let x: Result<u32, ParseIntError> = vec[0].parse();
            let y: Result<u32, ParseIntError> = vec[1].parse();
            if x.is_err() {
                return Err(VideoProcessorError::FailExtractResolution);
            }
            if y.is_err() {
                return Err(VideoProcessorError::FailExtractResolution);
            }
            return Ok((x.unwrap(), y.unwrap()));
        }
        Ok((0u32, 0u32))
    }


    /// Generate args for the ffmpeg commands args.
    fn generate_command<'a>(filename: &'a String, resolution: &'a Resolution) -> Result<std::vec::Vec<&'a str>, VideoProcessorError> {
        // we don't support video below 640x480 since it will meant to upscale it
        if resolution.lt(&(640u32, 480u32)) {
            return Err(VideoProcessorError::FailVideoTooSmall);
        }
        let mut command = vec!["-y", "-i", filename.as_str(), "-pix_fmt", "yuv420p", "-vcodec", "libx264", "-preset", "slow", "-g", "48", "-sc_threshold", "0"];
        let mut map = vec!["-map", "0:0", "-map", "0:1"];
        let mut gen_res = vec!["-s:v:0", "640x480", "-c:v:0", "libx264", "-b:v:0", "1000k"];
        let mut var_stream = vec!["-var_stream_map"];

        // Inverting the condition and constructing if progressively will be cleaner but I had some issue with the borrow checker
        if resolution.ge(&(3840u32, 2160u32)) {
            map.append(&mut vec!["-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1"]);
            gen_res.append(&mut vec!["-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k"]);
            gen_res.append(&mut vec!["-s:v:2", "1920x1080", "-c:v:2", "libx264", "-b:v:2", "4000k"]);
            gen_res.append(&mut vec!["-s:v:3", "3840x2160", "-c:v:3", "libx264", "-b:v:3", "8000k"]);
            var_stream.push("v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3");
        } else if resolution.ge(&(1920u32, 1080u32)) {
            map.append(&mut vec!["-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1"]);
            gen_res.append(&mut vec!["-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k"]);
            gen_res.append(&mut vec!["-s:v:2", "1920x1080", "-c:v:2", "libx264", "-b:v:2", "4000k"]);
            var_stream.push("v:0,a:0 v:1,a:1 v:2,a:2");
        } else if resolution.ge(&(1280u32, 720u32)) {
            map.append(&mut vec!["-map", "0:0", "-map", "0:1"]);
            gen_res.append(&mut vec!["-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k"]);
            var_stream.push("v:0,a:0 v:1,a:1");
        } else {
            var_stream.push("v:0,a:0");
        }

        command.append(&mut map);
        command.append(&mut gen_res);
        command.append(&mut vec!["-c:a", "aac", "-b:a", "128k", "-ac", "2"]);
        command.append(&mut var_stream);
        command.append(&mut vec!["-master_pl_name", "master.m3u8"]);
        command.append(&mut vec!["-f", "hls", "-hls_time", "6", "-hls_list_size", "0"]);
        command.append(&mut vec!["-hls_segment_filename", "v%v/part%d.ts", "v%v/part_index.m3u8"]);

        /*
            The command generated when processing 4K video:
            ffmpeg -y -i <filepath> \
              -pix_fmt yuv420p -vcodec libx264 -preset slow -g 48 -sc_threshold 0 \
              -map 0:0 -map 0:1 -map 0:0 -map 0:1 -map 0:0 -map 0:1 -map 0:0 -map 0:1 \
              -s:v:0 640x480 -c:v:0 libx264 -b:v:0 1000k \
              -s:v:1 1280x720 -c:v:1 libx264 -b:v:1 2000k  \
              -s:v:2 1920x1080 -c:v:2 libx264 -b:v:2 4000k  \
              -s:v:3 3840x2160 -c:v:3 libx264 -b:v:3 8000k  \
              -c:a aac -b:a 128k -ac 2 \
              -var_stream_map "v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3" \
              -master_pl_name master.m3u8 \
              -f hls -hls_time 6 -hls_list_size 0 \
              -hls_segment_filename "v%v/part%d.ts" \
              v%v/part_index.m3u8
         */

        Ok(command)
    }

    pub fn extract_miniature(filepath: &String) -> Result<String, VideoProcessorError> {
        // ffmpeg -i input.mp4 -ss 00:00:01.000 -vframes 1 miniature.png
        let output = Command::new("ffmpeg")
            .args(&["-i", filepath.as_str(), "-ss", "00:00:01.000", "-vframes", "1", "miniature.jpeg"])
            .output()?;
        if !output.status.success() {
            return Err(VideoProcessorError::FailExtractMiniature);
        }
        Ok(String::from("miniature.jpeg"))
    }
}