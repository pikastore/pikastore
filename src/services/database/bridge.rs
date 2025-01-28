use rusqlite::{params, Connection, Result};

fn insert(conn: &Connection, key: &[u8], value: &[u8]) -> Result<()> {
    conn.execute(
        "INSERT INTO my_table (key, value) VALUES (?1, ?2)",
        params![key, value],
    )?;
    Ok(())
}
