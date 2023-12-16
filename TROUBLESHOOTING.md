# Config Issues

# Http Issues

# Postgres Issues

## Failed to insert login pq: column "XXXX" of relation "XXXX" does not exist
Usually means the database is out of date, there's no reason way to update for now. Since this is pre-production I'll say wipe the database each run, or manually add those new columns.