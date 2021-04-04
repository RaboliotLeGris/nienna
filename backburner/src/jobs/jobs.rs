use crate::amqp::serialization::EventSerialization;
use crate::worker_pool::worker_pool::Job;
use crate::jobs::helpers;
use crate::s3::TS3Client;
use std::sync::Arc;

#[cfg(test)]
#[path = "./jobs_tests.rs"]
mod jobs_tests;

/// return a job closure to process given video from raw video to DASH
pub fn job_process_video(event: EventSerialization, s3_client: Arc<Box<dyn TS3Client>>) -> Job {
    println!("RUNNING: job_video_to_dash");
    let shared_s3_client = s3_client.clone();
    Box::new(move || {
        println!("RUNNING: job_video_to_dash JOB");
        if let Ok(working_folder) = helpers::go_to_working_directory() {
            println!("Current folder {:?}", std::env::current_dir());
            shared_s3_client.get(event.slug, String::from("TODO"));
            // go to working folder
            // get data
            // process
        }
    })
}