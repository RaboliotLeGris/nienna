#[cfg(test)]
mod jobs_tests {
    use crate::jobs::jobs::job_process_video;
    use crate::amqp::serialization::EventSerialization;
    use crate::s3::client_stub::S3ClientStub;
    use std::sync::Arc;

    #[test]
    fn should_job_video_to_dash() {
        // given
        let event = EventSerialization::new(String::from(""), String::from(""), String::from(""));
        let s3_client_stub = S3ClientStub::new();
        // when
        job_process_video(event, Arc::new(Box::new(s3_client_stub)))()
        // then
        // TODO a check of producted resources
    }
}