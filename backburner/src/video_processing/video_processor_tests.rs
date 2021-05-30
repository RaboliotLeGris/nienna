#[cfg(test)]
mod video_processor_tests {
    use crate::video_processing::video_processor::{VideoProcessor, Resolution};

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
        let expected: Vec<String> = vec![
            String::from("../.dev/samples/SampleVideo_1280x720_30mb.mp4"),
            String::from("../.dev/samples/sample_960x400_ocean_with_audio.avi"),
            String::from("../.dev/samples/sample_960x400_ocean_with_audio.mkv"),
            String::from("../.dev/samples/sample_960x400_ocean_with_audio.flv"),
        ];

        for filepath in expected {
            let _ = std::fs::create_dir("workdir");
            let _ = std::env::set_current_dir("workdir");

            VideoProcessor::process(&String::from(filepath)).unwrap();

            assert!(std::path::Path::new("master.m3u8").exists());
            assert!(std::path::Path::new("v0/part_index.m3u8").exists());
            assert!(std::path::Path::new("v0/part0.ts").exists());

            let _ = std::env::set_current_dir("..");
            let _ = std::fs::remove_dir_all("workdir");
        }
    }

    #[test]
    fn ffmpeg_command_generation() {
        let expected: Vec<(String, Resolution, bool, Vec<&str>)> = vec![
            (String::from("small_640x480.mkv"), (640, 480), false, vec!["-y", "-i", "small_640x480.mkv", "-pix_fmt", "yuv420p", "-vcodec", "libx264", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-map", "0:0", "-map", "0:1", "-s:v:0", "640x480", "-c:v:0", "libx264", "-b:v:0", "1000k", "-c:a", "aac", "-b:a", "128k", "-ac", "2", "-var_stream_map", "v:0,a:0", "-master_pl_name", "master.m3u8", "-f", "hls", "-hls_time", "6", "-hls_list_size", "0", "-hls_segment_filename", "v%v/part%d.ts", "v%v/part_index.m3u8"]),
            (String::from("SampleVideo_1280x720_30mb.mp4"), (1280, 720), false, vec!["-y", "-i", "SampleVideo_1280x720_30mb.mp4", "-pix_fmt", "yuv420p", "-vcodec", "libx264", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-s:v:0", "640x480", "-c:v:0", "libx264", "-b:v:0", "1000k", "-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k", "-c:a", "aac", "-b:a", "128k", "-ac", "2", "-var_stream_map", "v:0,a:0 v:1,a:1", "-master_pl_name", "master.m3u8", "-f", "hls", "-hls_time", "6", "-hls_list_size", "0", "-hls_segment_filename", "v%v/part%d.ts", "v%v/part_index.m3u8"]),
            (String::from("HD_1920x1080.mkv"), (1920, 1080), false, vec!["-y", "-i", "HD_1920x1080.mkv", "-pix_fmt", "yuv420p", "-vcodec", "libx264", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-s:v:0", "640x480", "-c:v:0", "libx264", "-b:v:0", "1000k", "-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k", "-s:v:2", "1920x1080", "-c:v:2", "libx264", "-b:v:2", "4000k", "-c:a", "aac", "-b:a", "128k", "-ac", "2", "-var_stream_map", "v:0,a:0 v:1,a:1 v:2,a:2", "-master_pl_name", "master.m3u8", "-f", "hls", "-hls_time", "6", "-hls_list_size", "0", "-hls_segment_filename", "v%v/part%d.ts", "v%v/part_index.m3u8"]),
            (String::from("4K-10bit.mkv"), (3840, 2160), false, vec!["-y", "-i", "4K-10bit.mkv", "-pix_fmt", "yuv420p", "-vcodec", "libx264", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-s:v:0", "640x480", "-c:v:0", "libx264", "-b:v:0", "1000k", "-s:v:1", "1280x720", "-c:v:1", "libx264", "-b:v:1", "2000k", "-s:v:2", "1920x1080", "-c:v:2", "libx264", "-b:v:2", "4000k", "-s:v:3", "3840x2160", "-c:v:3", "libx264", "-b:v:3", "8000k", "-c:a", "aac", "-b:a", "128k", "-ac", "2", "-var_stream_map", "v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3", "-master_pl_name", "master.m3u8", "-f", "hls", "-hls_time", "6", "-hls_list_size", "0", "-hls_segment_filename", "v%v/part%d.ts", "v%v/part_index.m3u8"]),
            (String::from("some-video"), (10, 10), true, vec![""]),
        ];

        for (filename, resolution, expect_error, expected) in expected.iter() {
            let res = VideoProcessor::generate_command(filename, resolution);
            if *expect_error {
                assert!(res.is_err())
            } else {
                assert!(res.is_ok());
                assert_eq!(*expected, res.unwrap());
            }
        }
    }

    #[test]
    #[serial]
    fn extract_resolution_from_source() {
        let _ = std::fs::create_dir("workdir");
        let _ = std::env::set_current_dir("workdir");

        let expected: Vec<(String, Resolution)> = vec![
            (String::from("../.dev/samples/sample_960x400_ocean_with_audio.avi"), (960u32, 400u32)),
            (String::from("../.dev/samples/SampleVideo_1280x720_30mb.mp4"), (1280u32, 720u32)),
            (String::from("../.dev/samples/4K-10bit.mkv"), (3840, 2160)),
        ];

        for (filename, resolution) in expected.iter() {
            let res = VideoProcessor::extract_resolution(&filename).unwrap();
            assert_eq!(*resolution, res)
        }

        let _ = std::env::set_current_dir("..");
        let _ = std::fs::remove_dir_all("workdir");
    }

    #[test]
    #[serial]
    fn extract_miniature_from_source() {
        let _ = std::fs::create_dir("workdir");
        let _ = std::env::set_current_dir("workdir");
        VideoProcessor::extract_miniature(&String::from("../.dev/samples/sample_960x400_ocean_with_audio.avi")).unwrap();
        let _ = std::env::set_current_dir("..");
        let _ = std::fs::remove_dir_all("workdir");
    }
}