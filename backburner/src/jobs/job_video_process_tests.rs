#[cfg(test)]
mod job_video_process_tests {
    use crate::jobs::job_video_process::job_process_video;
    use crate::amqp::serialization::EventSerialization;
    use crate::s3::client_stub::S3ClientStub;
    use std::sync::Arc;

    #[test]
    fn should_job_video_to_dash() {
        // given
        let event = EventSerialization::new(String::from(""), String::from(std::env::current_dir().unwrap().to_str().unwrap()), String::from("..\\tests\\samples\\SampleVideo_1280x720_2mb.mp4"));
        let s3_client_stub = S3ClientStub::new();
        // when
        job_process_video(event, Arc::new(Box::new(s3_client_stub)))()
        // then
        // TODO a check of producted resources
    }
}