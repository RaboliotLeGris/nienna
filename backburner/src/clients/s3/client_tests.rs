#[cfg(test)]
mod client_tests {
    use crate::clients::s3::client::S3Client;

    #[test]
    #[ignore]
    fn should_get_given_file() {
        // Need to have a test env to run this test
        let endpoint = String::from("http://s3:9000");
        let bucket_name = String::from("nienna-1");
        let slug = String::from("5Dg7VlmWQ0");
        let path = String::from("source.mp4");

        let s3_client = S3Client::new(endpoint, bucket_name.clone(), "minio".into(), "minio123".into());
        s3_client.get(&slug, &path).unwrap();
    }

    #[test]
    #[ignore]
    fn should_put_given_file() {
        // Need to have a test env to run this test
        let endpoint = String::from("http://s3:9000");
        let bucket_name = String::from("nienna-1");
        let slug = String::from("SomeSlug");
        let path = String::from("source.mp4");

        let s3_client = S3Client::new(endpoint, bucket_name.clone(), "minio".into(), "minio123".into());
        s3_client.put(&String::from(".dev/samples/SampleVideo_1280x720_2mb.mp4"), &slug,  &format!("HLS/{}", path)).unwrap();
    }
}