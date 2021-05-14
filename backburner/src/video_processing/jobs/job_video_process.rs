use std::sync::{Arc, mpsc};

use crate::clients::amqp::serialization::EventSerialization;
use crate::clients::s3::TS3Client;
use crate::video_processing::{errors::VideoProcessorError, video_processor::VideoProcessor};
use crate::worker_pool::jobs::{Job, job_helpers};
use crate::event_publisher::JobEventResult;

#[cfg(test)]
#[path = "./job_video_process_tests.rs"]
mod job_video_process_tests;

/// return a jobs closure to process given video from raw video to DASH
pub fn job_process_video(event: EventSerialization, s3_client: Arc<Box<dyn TS3Client>>, status_publisher: mpsc::Sender<JobEventResult>) -> Job {
    let shared_s3_client = s3_client.clone();
    Box::new(move || {
        let _ = match wrapper(&event, shared_s3_client) {
            Ok(_) => {
                status_publisher.send(JobEventResult::Success(event.slug.clone()))
            }
            Err(e) => {
                status_publisher.send(JobEventResult::Failure(event.slug.clone(), e.to_string()))
            }
        };
    })
}

fn wrapper(event: &EventSerialization, s3_client: Arc<Box<dyn TS3Client>>) -> Result<(), VideoProcessorError> {
    let original_dir = std::env::current_dir()?;
    let working_folder = job_helpers::go_to_working_directory()?;

    s3_client.get(&event.slug, &event.content)?;

    VideoProcessor::process(&event.content)?;
    let paths = std::fs::read_dir(".")?;
    for path in paths {
        let filename = path?.path().to_str().unwrap().to_string();
        if filename.ends_with(".ts") || filename.ends_with(".m3u8") {
            s3_client.put(&filename, &event.slug, &format!("HLS/{}", &filename))?;
        }
    }

    let miniature_path = VideoProcessor::extract_miniature(&event.content)?;
    s3_client.put(&miniature_path, &event.slug, &miniature_path)?;

    let _ = std::env::set_current_dir(&original_dir);
    let _ = std::fs::remove_dir_all(working_folder);
    Ok(())
}