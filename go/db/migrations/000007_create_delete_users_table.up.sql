CREATE TABLE IF NOT EXISTS delete_users(
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    artist_genre_id INT UNSIGNED NOT NULL,
    name VARCHAR(30) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_delete_users_artist_genres1 FOREIGN KEY (artist_genre_id) REFERENCES artist_genres (id) ON DELETE CASCADE ON UPDATE CASCADE
);
