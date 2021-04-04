use std::path::PathBuf;

use nanoid::nanoid;

use crate::jobs::errors::JobsError;

#[cfg(test)]
#[path = "./helpers_tests.rs"]
mod helpers_tests;

pub fn go_to_working_directory() -> Result<PathBuf, JobsError> {
    let mut working_dir = std::env::temp_dir();
    working_dir.push("nienna");
    working_dir.push("backburner");
    working_dir.push(nanoid!(10));
    debug!("Creating working folder: {:?}", working_dir.as_path());

    match std::fs::create_dir_all(&working_dir) {
        Err(ref e) if e.kind() == std::io::ErrorKind::AlreadyExists => {}
        Err(e) => {
            error!("failed to create working dir {:?} - err:{}", working_dir.as_path(), e);
            return Err(JobsError::FailCreateWorkingFolder);
        }
        _ => {}
    };
    std::env::set_current_dir(&working_dir)?;
    Ok(working_dir)
}
