CREATE TABLE IF NOT EXISTS artist_access_logs(
    user_id BIGINT UNSIGNED,
    ip VARCHAR(15),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_artist_access_logs_users1 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);
