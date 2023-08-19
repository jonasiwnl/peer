fn connect() -> Result<Connection, Box<dyn Error>> {
    let mut conn = Connection::open("data.db")?;
    conn.execute(
        "CREATE TABLE IF NOT EXISTS data (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            value TEXT NOT NULL
        )",
        [],
    )?;
    Ok(conn)
}
