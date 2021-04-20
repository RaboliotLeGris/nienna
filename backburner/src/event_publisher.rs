use std::sync::{mpsc, Arc};
use std::thread;
use crate::clients::amqp::client::AMQP;
use crate::clients::amqp::serialization::EventSerialization;

pub enum JobEventResult {
    Success(String),
    Failure(String, String),
}

const JOB_SUCCESS: &str = "EventVideoProcessingSucceed";
const JOB_FAILURE: &str = "EventVideoProcessingFail";

pub fn launch_job_event_publisher(amqp_addr: String) -> mpsc::Sender<JobEventResult> {
    let (sender, receiver) = mpsc::channel();

    thread::spawn(move || {
        async_global_executor::block_on(async {
            let mut amqp_client: AMQP = AMQP::new(amqp_addr, String::from("nienna_jobs_result")).await;
            loop {
                if let Ok(event) = receiver.recv() {
                    let job_event = match event {
                        JobEventResult::Success(slug) => {
                            EventSerialization::new(String::from(JOB_SUCCESS), slug, String::from("")).to_json()
                        }
                        JobEventResult::Failure(slug, reason) => {
                            EventSerialization::new(String::from(JOB_FAILURE), slug, reason).to_json()
                        }
                    };
                    if let Ok(json) = job_event {
                        amqp_client.publish(json).await;
                    } else {
                        warn!("failed to publish event to amqp");
                    }
                }
            }
        });
    });

    return sender;
}