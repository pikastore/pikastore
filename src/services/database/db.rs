use rusqlite::{Connection, Result};
pub fn open_db(path: &str) -> Result<()> {
    let conn = Connection::open(path)?;
    Ok(())
}