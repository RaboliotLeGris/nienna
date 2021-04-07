#[macro_use]
extern crate log;
#[cfg(test)]
#[macro_use]
extern crate serial_test;
// extern crate ffmpeg_next as ffmpeg;

use crate::amqp::client::AMQP;
use crate::s3::client::S3Client;
use std::sync::Arc;

mod amqp;
mod worker_pool;
mod jobs;
mod s3;
mod video_processing;

fn main() {
    if std::env::var("RUST_LOG").is_err() {
        std::env::set_var("RUST_LOG", "INFO");
    }
    env_logger::init();

    info!("Starting Backburner service");

    debug!("Creating WorkerPool");
    let worker_count: usize = std::env::var("BACKBURNER_WORKER_COUNT").unwrap_or(String::from("10")).parse::<usize>().expect("BACKBURNER_WORKER_COUNT must be a valid NON NULL and POSITIVE integer");
    let worker_pool = worker_pool::worker_pool::WorkerPool::new(worker_count);

    let addr = std::env::var("RABBITMQ_URI").unwrap();
    async_global_executor::block_on(async {
        let mut amqp_client: AMQP = AMQP::new(addr).await;
        while let Ok(event) = amqp_client.next().await {
            match event.event.as_str() {
                "EventVideoReadyForProcessing" => worker_pool.submit(jobs::job_video_process::job_process_video(event, Arc::new(Box::new(S3Client::new())))),
                _ => {}
            }
        }
    })
}