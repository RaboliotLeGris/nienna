#[cfg(test)]
mod video_processor_tests {
    use std::path::Path;

    use crate::video_processing::video_processor::VideoProcessor;

    #[test]
    fn should_extract_video_mimetype() {
        // give
        struct TestVideosPath<'a> {
            path: &'a str,
            mimetype: &'a str,
        }

        let videos_paths = vec![
            TestVideosPath{ path: ".dev/samples/SampleVideo_1280x720_2mb.mp4", mimetype: "video/mp4"},
            TestVideosPath{ path: ".dev/samples/SampleVideo_1280x720_30mb.mp4", mimetype: "video/mp4"},
            TestVideosPath{ path: ".dev/samples/sample_960x400_ocean_with_audio.flv", mimetype: "video/x-flv"},
            TestVideosPath{ path: ".dev/samples/sample_960x400_ocean_with_audio.avi", mimetype: "video/x-msvideo"},
        ];

        for v in videos_paths {
            let video_path = String::from(v.path);
            let video_processor = VideoProcessor::new();

            // when
            let mimetype = video_processor.get_mimetype(video_path).unwrap();

            // then
            let expected_mimetype = String::from(v.mimetype);
            assert_eq!(expected_mimetype, mimetype);
        }
    }

    #[test]
    fn should_fail_to_extract_video_mimetype() {
        let filepath = String::from("some/wrong/path.mp4");
        let video_processor = VideoProcessor::new();
        let mimetype = video_processor.get_mimetype(filepath).unwrap();
        assert_eq!("cannot", mimetype);
    }
}