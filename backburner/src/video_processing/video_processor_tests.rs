#[cfg(test)]
mod video_processor_tests {
    use crate::video_processing::video_processor::VideoProcessor;

    #[test]
    #[serial]
    fn should_extract_video_mimetype() {
        // give
        struct TestVideosPath<'a> {
            path: &'a str,
            mimetype: &'a str,
        }

        let videos_paths = vec![
            TestVideosPath { path: ".dev/samples/SampleVideo_1280x720_2mb.mp4", mimetype: "video/mp4" },
            TestVideosPath { path: ".dev/samples/SampleVideo_1280x720_30mb.mp4", mimetype: "video/mp4" },
            TestVideosPath { path: ".dev/samples/sample_960x400_ocean_with_audio.flv", mimetype: "video/x-flv" },
            TestVideosPath { path: ".dev/samples/sample_960x400_ocean_with_audio.avi", mimetype: "video/x-msvideo" },
            TestVideosPath { path: ".dev/samples/sample_960x400_ocean_with_audio.mkv", mimetype: "video/x-matroska" },
        ];

        for v in videos_paths {
            let video_path = String::from(v.path);

            // when
            let mimetype = VideoProcessor::extract_mimetype(video_path).unwrap();

            // then
            let expected_mimetype = String::from(v.mimetype);
            assert_eq!(expected_mimetype, mimetype);
        }
    }

    #[test]
    #[serial]
    fn should_fail_to_extract_video_mimetype() {
        let filepath = String::from("some/wrong/path.mp4");
        let result = VideoProcessor::extract_mimetype(filepath);
        assert!(result.is_err());
    }

    #[test]
    #[serial]
    fn process_video_to_hls() {
        std::fs::create_dir("workdir");
        std::env::set_current_dir("workdir");
        VideoProcessor::process(&String::from("../.dev/samples/SampleVideo_1280x720_30mb.mp4")).unwrap();
        std::env::set_current_dir("..");
        std::fs::remove_dir_all("workdir");
    }

    #[test]
    #[serial]
    fn extract_miniature_from_source() {
        std::fs::create_dir("workdir");
        std::env::set_current_dir("workdir");
        VideoProcessor::extract_miniature(&String::from("../.dev/samples/SampleVideo_1280x720_30mb.mp4")).unwrap();
        std::env::set_current_dir("..");
        std::fs::remove_dir_all("workdir");
    }
}