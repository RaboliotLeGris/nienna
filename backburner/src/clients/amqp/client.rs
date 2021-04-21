use futures_lite::StreamExt;
use lapin::{Channel, Connection, ConnectionProperties, Consumer, options::*, types::FieldTable, BasicProperties};

use crate::clients::amqp::errors::AmqpError;
use crate::clients::amqp::serialization::EventSerialization;

pub struct AMQP {
    conn: Connection,
    channel: Channel,
    consumer: Consumer,
    queue: String,
}

impl AMQP {
    pub async fn new(uri: String, queue: String) -> Self {
        let conn = Connection::connect(uri.as_str(), ConnectionProperties::default())
            .await
            .expect("connection error");
        info!("Connected to rabbitMQ");

        let channel = conn.create_channel().await.expect("create_channel");
        debug!("Channel status {:?}", conn.status().state());

        debug!("Declaring queue: \"{}\"", queue);
        let _queue = channel
            .queue_declare(
                queue.as_str(),
                QueueDeclareOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("queue_declare");

        debug!("Create consumer for queue: \"{}\"", queue);
        let consumer = channel
            .basic_consume(
                queue.as_str(),
                "backburner",
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("basic_consume");

        AMQP {
            conn,
            channel,
            consumer,
            queue,
        }
    }

    pub async fn next(&mut self) -> Result<EventSerialization, AmqpError> {
        if let Some(delivery) = self.consumer.next().await {
            if let Ok(delivery) = delivery {
                delivery.0.basic_ack(delivery.1.delivery_tag, BasicAckOptions::default()).await;
                return EventSerialization::from(delivery.1.data);
            }
        }
        Err(AmqpError::FailFetchEvent)
    }

    pub async fn publish(&mut self, payload: Vec<u8>) -> Result<(), AmqpError> {
        if self.channel.basic_publish("", self.queue.as_str(), BasicPublishOptions::default(), payload, BasicProperties::default()).await.is_ok() {
            return Ok(())
        }
        Err(AmqpError::FailPublishEvent)
    }

    pub async fn ack(&mut self, delivery_tag: u64) -> Result<(), AmqpError> {
        if self.channel.basic_ack(delivery_tag, BasicAckOptions::default()).await.is_ok() {
            return Ok(())
        }
        Err(AmqpError::FailPublishEvent)
    }
}