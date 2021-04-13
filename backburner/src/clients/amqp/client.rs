use lapin::{Connection, ConnectionProperties, options::*, types::FieldTable, Channel, Consumer};
use futures_lite::StreamExt;
use crate::clients::amqp::serialization::EventSerialization;
use crate::clients::amqp::errors::AmqpError;

pub struct AMQP {
    conn: Connection,
    channel: Channel,
    consumer: Consumer,
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
            consumer
        }
    }

    pub async fn next(&mut self) -> Result<EventSerialization, AmqpError> {
        if let Some(delivery) = self.consumer.next().await {
            if let Ok(delivery) = delivery {
                delivery.1.ack(BasicAckOptions::default());
                return EventSerialization::from(delivery.1.data);
            }
        }
        Err(AmqpError::FailFetchEvent)
    }

    // pub fn publish(&mut self) -> Result<(), AmqpError> {
    //     self.channel.basic_publish()
    //
    //     Ok(())
    // }

    
}