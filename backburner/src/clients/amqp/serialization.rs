use serde::{Deserialize, Serialize};
use serde_json::error;

use crate::clients::amqp::errors::AmqpError;

#[cfg(test)]
#[path = "./serialization_tests.rs"]
mod serialization_tests;

#[derive(Serialize, Deserialize, Debug, PartialEq)]
pub struct EventSerialization {
    pub event: String,
    pub slug: String,
    pub content: String,
}

impl EventSerialization {
    pub fn new(event: String, slug: String, content: String) -> EventSerialization {
        EventSerialization {
            event,
            slug,
            content,
        }
    }

    pub fn from(raw: Vec<u8>) -> Result<Self, AmqpError> {
        let event: error::Result<EventSerialization> = serde_json::from_slice(raw.as_slice());
        if event.is_err() {
            return Err(AmqpError::FailDeserialization);
        }
        let event = event.unwrap();
        if !EventSerialization::check_event(&event) {
            return Err(AmqpError::UnrecognizedEvent);
        }
        return Ok(event);
    }

    pub fn to_json(&self) -> Result<Vec<u8>, AmqpError> {
        if let Ok(serialize) = serde_json::to_vec(self) {
            return Ok(serialize);
        }
        Err(AmqpError::FailSerialization)
    }

    fn check_event(event: &EventSerialization) -> bool {
        match event.event.as_str() {
            "EventVideoReadyForProcessing" => true,
            _ => false
        }
    }
}
