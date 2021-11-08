#[cfg(test)]
mod worker_pool_tests {
    use std::{thread, time};
    use std::sync::{Arc, Mutex};
    use std::time::Instant;

    use crate::worker_pool::pool::WorkerPool;

    #[test]
    fn should_create_5_worker_and_execute_a_simple_job() {
        let worker_count = 5;
        let worker_pool = WorkerPool::new(worker_count);
        let w = Arc::new(Mutex::new(5));

        for _ in 0..worker_count {
            let w = Arc::clone(&w);
            worker_pool.submit(move || {
                let mut count = w.lock().unwrap();
                *count -= 1;
            });
        }

        thread::sleep(time::Duration::from_secs(2));
        let w = w.lock().expect("failed to lock value mutex");
        assert_eq!(0, *w, "{} should be executed by WorkerPool", worker_count);
        worker_pool.terminate_all();
        thread::sleep(time::Duration::from_secs(1));
    }

    #[test]
    fn should_run_long_lasting_jobs_concurrently_within_constrained_time() {
        // A test a bit wacky but it allows to verify that the jobs are run concurrently and are not locked by a mutex
        let worker_count = 2;
        let worker_pool = WorkerPool::new(worker_count);
        let w = Arc::new(Mutex::new(2));
        let start = Instant::now();
        for _ in 0..worker_count {
            let w = Arc::clone(&w);
            worker_pool.submit(move || {
                // really long running jobs
                thread::sleep(time::Duration::from_secs(5));
                let mut count = w.lock().unwrap();
                *count -= 1;
            });
        }
        thread::sleep(time::Duration::from_secs(6));
        worker_pool.terminate_all();
        let w = w.lock().expect("failed to lock value mutex");
        assert_eq!(0, *w, "Jobs should have ended");
        assert!(start.elapsed().as_secs() < 7, "jobs took to long to run");
    }
}