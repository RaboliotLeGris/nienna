#[macro_use]
extern crate log;

use futures_lite::StreamExt;
use lapin::{
    BasicProperties, Connection, ConnectionProperties, options::*, publisher_confirm::Confirmation,
    Result, types::FieldTable,
};

use crate::amqp::serialization::EventSerialization;

mod amqp;
mod worker_pool;

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
        let conn = Connection::connect(&addr, ConnectionProperties::default())
            .await
            .expect("connection error");
        debug!("Connected to rabbitMQ");

        let channel = conn.create_channel().await.expect("create_channel");
        debug!("Channel status {:?}", conn.status().state());

        debug!("Declaring queue: \"nienna_backburner\"");
        let queue = channel
            .queue_declare(
                "nienna_backburner",
                QueueDeclareOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("queue_declare");

        debug!("Create consumer for queue: \"nienna_backburner\"");
        let mut consumer = channel
            .basic_consume(
                "nienna_backburner",
                "backburner",
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("basic_consume");

        while let Some(delivery) = consumer.next().await {
            if let Ok(delivery) = delivery {
                debug!("EventSerialization {:#?}", EventSerialization::from(delivery.1.data));
                // delivery
                //     .ack(BasicAckOptions::default())
                //     .await
                //     .expect("basic_ack");
            }
        }
    })
}