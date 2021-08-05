CREATE TABLE IF NOT EXISTS users(
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    artist_genre_id INT UNSIGNED,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(255) NOT NULL,
    email_verified_at datetime,
    password VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    comment VARCHAR(200),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_artist_genres1 FOREIGN KEY (artist_genre_id) REFERENCES artist_genres (id) ON DELETE CASCADE ON UPDATE CASCADE
);
