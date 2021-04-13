#[cfg(test)]
mod job_helpers_tests {
    use crate::jobs::helpers::go_to_working_directory;

    #[test]
    #[serial]
    fn should_go_and_return_a_tmp_folder() {
        let original = std::env::current_dir().unwrap();
        let res = go_to_working_directory();
        assert!(res.is_ok());

        let res = res.unwrap();
        let path_as_str = res.to_str().unwrap();
        assert!(path_as_str.contains("nienna"));

        let current = std::env::current_dir().unwrap();
        assert_eq!(current.to_str().unwrap(), path_as_str);

        std::env::set_current_dir(original).unwrap();
    }
}