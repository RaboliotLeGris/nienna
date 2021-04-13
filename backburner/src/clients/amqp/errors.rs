#[derive(Debug)]
pub enum AmqpError{
    FailFetchEvent,
    FailDeserialization,
    UnrecognizedEvent,
}

impl std::error::Error for AmqpError {}

impl std::fmt::Display for AmqpError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            AmqpError::FailDeserialization => write!(f, "Failed to parse json"),
            AmqpError::FailFetchEvent => write!(f, "Failed to fetch an event"),
            AmqpError::UnrecognizedEvent => write!(f, "Unable to recognize event"),
        }
    }
}