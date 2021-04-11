use crate::amqp::serialization::EventSerialization;
use crate::worker_pool::worker_pool::Job;
use crate::s3::TS3Client;
use std::sync::Arc;
use crate::jobs::helpers;
use crate::video_processing::video_processor::VideoProcessor;

#[cfg(test)]
#[path = "./job_video_process_tests.rs"]
mod job_video_process_tests;

/// return a job closure to process given video from raw video to DASH
pub fn job_process_video(event: EventSerialization, s3_client: Arc<Box<dyn TS3Client>>) -> Job {
    let shared_s3_client = s3_client.clone();
    Box::new(move || {
        let original_dir = std::env::current_dir().unwrap(); // TODO do not allow panic
        if let Ok(working_folder) = helpers::go_to_working_directory() {
            shared_s3_client.get(&event.slug, &event.filename); // Is OK ?

            let video_processor = VideoProcessor::new();
            let _mimetype = video_processor.extract_mimetype(event.filename);
            //
            // video_processor.process(&event.filename);
            //
            // let dir = std::fs::read_dir(".").unwrap();
            // for entry in dir {
            //     let entry = entry.unwrap();
            //     let path = entry.path();
            //     // shared_s3_client.put(path, &event.slug)
            // }


            // Clean working directory afterwards
            let _ = std::env::set_current_dir(&original_dir);
            let _ = std::fs::remove_dir_all(working_folder);
        }
    })
}