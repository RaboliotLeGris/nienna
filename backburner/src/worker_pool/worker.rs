use std::sync::{Arc, mpsc, Mutex};
use std::thread;
use std::time::Instant;

use crate::worker_pool::pool::Message;

pub struct Worker {
    _id: usize,
    _thread: thread::JoinHandle<()>,
}

impl Worker {
    pub fn new(id: usize, receiver: Arc<Mutex<mpsc::Receiver<Message>>>) -> Worker {
        let thread = thread::spawn(move || loop {
            let job = receiver.lock().expect("Failed to lock receiver").recv();
            match job {
                Ok(Message::Terminate) => break,
                Ok(Message::NewJob(job)) => {
                    debug!("[{:?}][Worker {}] received a jobs; executing.", Instant::now(), id);
                    job();
                    debug!("[{:?}][Worker {}]: jobs finished", Instant::now(), id)
                }
                Err(e) => {
                    error!("unable to get a jobs with error: {}", e);
                }
            };
        });

        Worker {
            _id: id,
            _thread: thread,
        }
    }
}