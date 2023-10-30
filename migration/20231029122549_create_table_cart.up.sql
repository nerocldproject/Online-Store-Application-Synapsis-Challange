CREATE TABLE IF NOT EXISTS cart (
    cart_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    quantity INT NOT NULL,
    invoice_id VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES product(product_id)
    ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES customer(user_id)
    ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (invoice_id) REFERENCES invoice(invoice_id)
    ON DELETE RESTRICT ON UPDATE CASCADE
);