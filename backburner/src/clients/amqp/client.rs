use futures_lite::StreamExt;
use lapin::{Channel, Connection, ConnectionProperties, Consumer, options::*, types::FieldTable, BasicProperties};

use crate::clients::amqp::errors::AmqpError;
use crate::clients::amqp::serialization::EventSerialization;
use std::time::Duration;

pub struct AMQP {
    uri: String,
    conn: Connection,
    channel: Channel,
    consumer: Option<Consumer>,
    queue: String,
}

impl AMQP {
    pub async fn new(uri: String, queue: String, enable_consumer: bool) -> Self {
        let (conn, channel, consumer) = AMQP::connect(&uri, &queue, enable_consumer).await.expect("Must establish connection with AMQP server");
        AMQP {
            uri,
            conn,
            channel,
            consumer,
            queue,
        }
    }

    pub async fn next(&mut self) -> Result<EventSerialization, AmqpError> {
        if self.consumer.is_none() {
            return Err(AmqpError::ConsumerIsNone);
        }
        if self.conn.status().state() != lapin::ConnectionState::Connected {
            self.reconnect().await?;
        }
        if let Some(Ok(delivery)) = self.consumer.as_mut().unwrap().next().await {
            let _ = delivery.0.basic_ack(delivery.1.delivery_tag, BasicAckOptions::default()).await;
            return EventSerialization::from(delivery.1.data);
        }
        Err(AmqpError::FailFetchEvent)
    }

    pub async fn publish(&mut self, payload: Vec<u8>) -> Result<(), AmqpError> {
        if self.conn.status().state() != lapin::ConnectionState::Connected {
            self.reconnect().await?;
        }
        if self.channel.basic_publish("", self.queue.as_str(), BasicPublishOptions::default(), payload, BasicProperties::default()).await.is_ok() {
            return Ok(());
        }
        Err(AmqpError::FailPublishEvent)
    }

    async fn connect(uri: &String, queue: &String, enable_consumer: bool) -> Result<(Connection, Channel, Option<Consumer>), AmqpError> {
        let (conn, channel) = AMQP::create_conn(uri).await?;
        let _ = AMQP::create_queue(&channel, &queue).await?;

        let mut consumer = None;
        if enable_consumer {
            debug!("Create consumer for queue: \"{}\"", queue);
            consumer = Some(AMQP::create_consumer(&channel, &queue).await?);
        }

        Ok((conn, channel, consumer))
    }

    async fn reconnect(&mut self) -> Result<(), AmqpError> {
        let mut retry_count = 10;
        let mut connected = false;
        while !connected && retry_count > 0 {
            if let Ok((conn, channel, consumer)) = AMQP::connect(&self.uri, &self.queue, self.consumer.is_some()).await {
                self.conn = conn;
                self.channel = channel;
                self.consumer = consumer;
                connected = true;
                break;
            }
            retry_count = retry_count - 1;
            tokio::time::sleep(Duration::from_secs(5)).await;
        }
        if !connected {
            return Err(AmqpError::ConnectionFailed);
        }
        Ok(())
    }

    async fn create_conn(uri: &String) -> Result<(Connection, Channel), AmqpError> {
        info!("Connecting to rabbitMQ");
        let conn = Connection::connect(uri.as_str(), ConnectionProperties::default())
            .await?;

        info!("Creating amqp channel");
        let channel = conn.create_channel().await?;
        debug!("Channel status {:?}", conn.status().state());

        Ok((conn, channel))
    }

    async fn create_queue(channel: &Channel, queue: &String) -> Result<(), AmqpError> {
        debug!("Declaring queue: \"{}\"", queue);
        let _queue = channel
            .queue_declare(
                queue.as_str(),
                QueueDeclareOptions::default(),
                FieldTable::default(),
            )
            .await?;
        Ok(())
    }

    async fn create_consumer(channel: &Channel, queue: &String) -> Result<Consumer, AmqpError> {
        let consumer = channel
            .basic_consume(
                queue.as_str(),
                "backburner",
                BasicConsumeOptions::default(),
                FieldTable::default(),
            )
            .await?;
        Ok(consumer)
    }
}