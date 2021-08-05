CREATE TABLE IF NOT EXISTS product_access_logs(
    user_id BIGINT UNSIGNED,
    product_id BIGINT UNSIGNED NOT NULL,
    ip VARCHAR(15),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_access_logs_users1 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_product_access_logs_products1 FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE ON UPDATE CASCADE
);
