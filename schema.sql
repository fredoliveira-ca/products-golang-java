
CREATE TABLE product (
    product_id text PRIMARY KEY,
    price_in_cents integer,
    title text,
    description text
);

CREATE TABLE "user" (
    user_id text PRIMARY KEY,
    first_name text,
    last_name text,
    date_of_birth date NOT NULL
);
