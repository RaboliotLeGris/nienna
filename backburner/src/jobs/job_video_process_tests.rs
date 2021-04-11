#[cfg(test)]
mod job_video_process_tests {
    use crate::jobs::job_video_process::job_process_video;
    use crate::amqp::serialization::EventSerialization;
    use std::sync::Arc;
    use crate::s3::client::S3Client;

    #[test]
    fn should_job_video_to_dash() {
        // given
        let event = EventSerialization::new(String::from(""), String::from("someslug"), String::from("SampleVideo_1280x720_30mb.mp4"));
        let s3_client = S3Client::new(String::from("http://s3:9000"), String::from("nienna-1"), String::from("minio"), String::from("minio123"));
        // when
        job_process_video(event, Arc::new(Box::new(s3_client)))()
        // then
        // TODO a check of producted resources
    }
}