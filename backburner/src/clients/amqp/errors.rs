use lapin::Error;

#[derive(Debug)]
pub enum AmqpError {
    FailFetchEvent,
    FailDeserialization,
    FailSerialization,
    FailPublishEvent,
    ConsumerIsNone,
    UnrecognizedEvent,
    ConnectionFailed,
    QueueCreationFailed
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
            AmqpError::ConsumerIsNone => write!(f, "No consumer created"),
            AmqpError::ConnectionFailed => write!(f, "Connection failed"),
            AmqpError::QueueCreationFailed => write!(f, "Failed to create a queue"),
        }
    }
}

impl From<lapin::Error> for AmqpError {
    fn from(e: lapin::Error) -> Self {
        match e {
            Error::ChannelsLimitReached => AmqpError::ConnectionFailed,
            Error::InvalidProtocolVersion(_) =>  AmqpError::UnrecognizedEvent,
            Error::InvalidChannel(_) => AmqpError::ConnectionFailed,
            Error::InvalidChannelState(_) =>  AmqpError::UnrecognizedEvent,
            Error::InvalidConnectionState(_) => AmqpError::ConnectionFailed,
            Error::IOError(_) =>  AmqpError::UnrecognizedEvent,
            Error::ParsingError(_) =>  AmqpError::UnrecognizedEvent,
            Error::ProtocolError(_) =>  AmqpError::UnrecognizedEvent,
            Error::SerialisationError(_) => AmqpError::FailSerialization,
            _ => AmqpError::UnrecognizedEvent
        }
    }
}