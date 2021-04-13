use std::sync::{Arc, mpsc, Mutex};

use crate::worker_pool::worker::Worker;
use crate::worker_pool::jobs::job::Job;

#[cfg(test)]
#[path = "./worker_pool_tests.rs"]
mod worker_pool_tests;

pub enum Message {
    NewJob(Job),
    Terminate,
}

pub struct WorkerPool {
    worker_count: usize,
    workers: Vec<Worker>,
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
            worker_count,
            workers,
            sender,
        }
    }

    pub fn submit<F>(&self, f: F)
        where
            F: FnOnce() + Send + 'static
    {
        let job = Box::new(f);

        self.sender.send(Message::NewJob(job)).unwrap();
    }

    pub fn terminate_all(&self) {
        self.sender.send(Message::Terminate);
    }
}

