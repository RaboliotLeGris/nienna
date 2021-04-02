mod amqp;

use futures_lite::StreamExt;
use lapin::{
    options::*, publisher_confirm::Confirmation, types::FieldTable, BasicProperties, Connection,
    ConnectionProperties, Result,
};
use crate::amqp::serialization::EventSerialization;

#[macro_use]
extern crate log;

fn main() {
    if std::env::var("RUST_LOG").is_err() {
        std::env::set_var("RUST_LOG", "INFO");
    }
    env_logger::init();

    info!("Starting service");

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