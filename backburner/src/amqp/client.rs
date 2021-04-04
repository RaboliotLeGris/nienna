use lapin::{Connection, ConnectionProperties, options::*, types::FieldTable, Channel, Consumer};
use crate::amqp::serialization::EventSerialization;
use crate::amqp::errors::AmqpError;
use futures_lite::StreamExt;

pub struct AMQP {
    conn: Connection,
    channel: Channel,
    consumer: Consumer,
}

impl AMQP {
    pub async fn new(uri: String) -> Self {
        let conn = Connection::connect(uri.as_str(), ConnectionProperties::default())
            .await
            .expect("connection error");
        debug!("Connected to rabbitMQ");

        let channel = conn.create_channel().await.expect("create_channel");
        debug!("Channel status {:?}", conn.status().state());

        debug!("Declaring queue: \"nienna_backburner\"");
        let _queue = channel
            .queue_declare(
                "nienna_backburner",
                QueueDeclareOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("queue_declare");

        debug!("Create consumer for queue: \"nienna_backburner\"");
        let consumer = channel
            .basic_consume(
                "nienna_backburner",
                "backburner",
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await
            .expect("basic_consume");

        AMQP {
            conn,
            channel,
            consumer
        }
    }

    pub async fn next(&mut self) -> Result<EventSerialization, AmqpError> {
        if let Some(delivery) = self.consumer.next().await {
            if let Ok(delivery) = delivery {
                return EventSerialization::from(delivery.1.data);
            }
        }
        Err(AmqpError::FailFetchEvent)
    }

    
}