#[macro_use]
extern crate log;
extern crate s3 as rust_s3;
#[cfg(test)]
#[macro_use]
extern crate serial_test;

use std::sync::Arc;

use crate::clients::amqp::client::AMQP;
use crate::clients::s3::client::S3Client;
use crate::video_processing::jobs::job_video_process::job_process_video;

mod clients;
mod worker_pool;
mod video_processing;
mod event_publisher;

#[tokio::main]
pub async fn main() -> Result<(), ()> {
    if std::env::var("RUST_LOG").is_err() {
        std::env::set_var("RUST_LOG", "INFO");
    }
    env_logger::init();

    info!("Starting Backburner service");
    info!("Create WorkerPool");
    let worker_count: usize = std::env::var("BACKBURNER_WORKER_COUNT").unwrap_or_else(|_| String::from("10")).parse::<usize>().expect("BACKBURNER_WORKER_COUNT must be a valid NON NULL and POSITIVE integer");
    let worker_pool = worker_pool::pool::WorkerPool::new(worker_count);

    debug!("Create S3Client");
    let s3_client = S3Client::new(std::env::var("S3_URI").unwrap(), "nienna-1".into(), std::env::var("S3_ACCESS_KEY").unwrap(), std::env::var("S3_SECRET_KEY").unwrap());

    debug!("Create event publisher");
    let addr = std::env::var("AMQP_URI").unwrap();
    let sender = event_publisher::launch_job_event_publisher(addr.clone());

    let mut amqp_client: AMQP = AMQP::new(addr, String::from("nienna_backburner"), true).await;
    loop {
        match amqp_client.next().await {
            Ok(event) => {
                match event.event.as_str() {
                    "EventVideoReadyForProcessing" => worker_pool.submit(job_process_video(event, Arc::new(Box::new(s3_client.clone())), sender.clone())),
                    _ => {}
                }
            }
            Err(e) => error!("Failed to deserialize event with err: {:?}", e)
        }
    };
}