#[derive(Debug)]
pub enum AmqpError{
    FailFetchEvent,
    FailDeserialization,
    FailSerialization,
    FailPublishEvent,
    FailAckEvent,
    FailNAckEvent,
    UnrecognizedEvent,
}

impl std::error::Error for AmqpError {}

impl std::fmt::Display for AmqpError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            AmqpError::FailDeserialization => write!(f, "Failed to parse json"),
            AmqpError::FailSerialization => write!(f, "Failed to serialize struct to json"),
            AmqpError::FailFetchEvent => write!(f, "Failed to fetch an event"),
            AmqpError::FailPublishEvent => write!(f, "Failed publish an event"),
            AmqpError::UnrecognizedEvent => write!(f, "Unable to recognize event"),
            AmqpError::FailAckEvent => write!(f, "Failed to ack an event"),
            AmqpError::FailNAckEvent => write!(f, "Failed to nack an event")
        }
    }
}