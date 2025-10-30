-- 001_init.sql

CREATE TABLE IF NOT EXISTS services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    method TEXT DEFAULT 'GET',
    expected_status INTEGER DEFAULT 200,
    expected_substring TEXT,
    timeout_ms INTEGER DEFAULT 5000,
    interval_sec INTEGER DEFAULT 120,
    retries INTEGER DEFAULT 2,
    enabled INTEGER DEFAULT 1,
    status  TEXT DEFAULT 'UNKNOWN',
    latency_ms INTEGER DEFAULT 0,
    last_checked_at DATETIME,
    error TEXT
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_services_enabled ON services (enabled);


CREATE TABLE IF NOT EXISTS check_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    service_id INTEGER NOT NULL,
    status_code INTEGER,
    ok BOOLEAN,
    response_time_ms INTEGER,
    checked_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(service_id) REFERENCES services(id)
);
