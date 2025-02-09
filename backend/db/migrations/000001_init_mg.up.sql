CREATE TABLE host_statuses (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(100) NOT NULL UNIQUE,
    ping_time_ms INT,
    last_success_date TIMESTAMP
);

CREATE INDEX idx_host_statuses_ip_address ON host_statuses(ip_address);