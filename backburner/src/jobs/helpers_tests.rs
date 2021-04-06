#[cfg(test)]
mod helpers_tests {
    use crate::jobs::helpers::go_to_working_directory;

    // FIXME: Broke others tests
    #[test]
    #[ignore]
    fn should_go_and_return_a_tmp_folder() {
        let res = go_to_working_directory();
        assert!(res.is_ok());
        let res = res.unwrap();
        let path_as_stf =  res.to_str().unwrap();
        assert!(path_as_stf.contains("nienna"));
        let current = std::env::current_dir().unwrap();
        assert_eq!(current.to_str().unwrap(), path_as_stf);
    }
}