use std::time;

fn main() {
    let ten_seconds = time::Duration::from_secs(10);
    let start = time::Instant::now();
    let mut step = time::Instant::now();

    println!("backburner - starting");
    loop {
        if step.elapsed() > ten_seconds {
            println!("backburner - Looping - Elapsed {:?}", start.elapsed());
            step = time::Instant::now();
        }
    }
}
