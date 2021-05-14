use std::sync::{Arc, mpsc, Mutex};

use crate::worker_pool::worker::Worker;
use crate::worker_pool::jobs::Job;

#[cfg(test)]
#[path = "./pool_tests.rs"]
mod pool_tests;

#[allow(dead_code)]
pub enum Message {
    NewJob(Job),
    Terminate,
}

pub struct WorkerPool {
    _worker_count: usize,
    _workers: Vec<Worker>,
    sender: mpsc::Sender<Message>,
}

impl WorkerPool {
    pub fn new(worker_count: usize) -> WorkerPool {
        assert!(worker_count > 0);

        let (sender, receiver) = mpsc::channel();
        let receiver = Arc::new(Mutex::new(receiver));

        let mut workers = Vec::with_capacity(worker_count);
        for id in 0..worker_count {
            workers.push(Worker::new(id, Arc::clone(&receiver)))
        }

        WorkerPool {
            _worker_count: worker_count,
            _workers: workers,
            sender,
        }
    }

    pub fn submit<F>(&self, f: F)
        where
            F: FnOnce() + Send + 'static
    {
        let job = Box::new(f);

        let _ = self.sender.send(Message::NewJob(job)).unwrap();
    }

    #[allow(dead_code)]
    pub fn terminate_all(&self) {
        let _ = self.sender.send(Message::Terminate);
    }
}

