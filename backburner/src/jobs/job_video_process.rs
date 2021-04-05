use crate::amqp::serialization::EventSerialization;
use crate::worker_pool::worker_pool::Job;
use crate::s3::TS3Client;
use std::sync::Arc;
use crate::jobs::helpers;

#[cfg(test)]
#[path = "./job_video_process_tests.rs"]
mod job_video_process_tests;

/// return a job closure to process given video from raw video to DASH
pub fn job_process_video(event: EventSerialization, s3_client: Arc<Box<dyn TS3Client>>) -> Job {
    let shared_s3_client = s3_client.clone();
    Box::new(move || {
        let original_dir = std::env::current_dir().unwrap(); // TODO do not allow panic
        if let Ok(working_folder) = helpers::go_to_working_directory() {
            shared_s3_client.get(event.slug, event.filename);

            // Do the actual work
                // Check mimetype
                //

            // Clean working directory afterwards
            std::env::set_current_dir(&original_dir).unwrap();
            std::fs::remove_dir_all(working_folder).unwrap();
        }
    })
}