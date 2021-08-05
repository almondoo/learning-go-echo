CREATE TABLE IF NOT EXISTS product_images(
    product_id BIGINT UNSIGNED NOT NULL,
    image VARCHAR(255) NOT NULL,
    CONSTRAINT fk_product_images_products1 FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE ON UPDATE CASCADE
);
