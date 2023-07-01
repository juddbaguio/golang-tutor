ALTER TABLE users ADD COLUMN
    created_at timestamptz default now();

ALTER TABLE orders add column receiving_date date default now();