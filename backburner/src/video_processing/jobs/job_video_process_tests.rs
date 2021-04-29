#[cfg(test)]
mod job_video_process_tests {
    use std::sync::{Arc, mpsc};

    use crate::clients::amqp::serialization::EventSerialization;
    use crate::clients::s3::client::S3Client;
    use crate::video_processing::jobs::job_video_process::job_process_video;
    use crate::event_publisher::JobEventResult;

    #[test]
    #[ignore]
    fn should_job_video_to_dash() {
        // TODO FIXME
        // given
        let event = EventSerialization::new(String::from(""), String::from("someslug"), String::from("SampleVideo_1280x720_30mb.mp4"));
        let s3_client = S3Client::new(String::from("http://s3:9000"), String::from("nienna-1"), String::from("minio"), String::from("minio123"));
        let (sender, receiver) = mpsc::channel();
        // when
        job_process_video(event, Arc::new(Box::new(s3_client)), sender)();
        // then
        let res = receiver.recv();
        println!("RES {:?}", res);
        assert_eq!(res.unwrap(), JobEventResult::Success(String::from("someslug")));
    }
}