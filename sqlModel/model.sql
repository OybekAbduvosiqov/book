CREATE TABLE books (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE categorys (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE book_category (
    id UUID PRIMARY KEY,
    categorys_id UUID references categorys(id), 
    books_id UUID references books(id)
);

Insert Into book_category (id, categorys_id, books_id) values 
('eab45685-7e78-4e37-ab15-d977e45284b9', 'fbcc1bf3-4c54-4bab-956e-2615ba163d0f','2609cc1c-bd5d-4fa3-abd4-65a9bdc15cf8'),
('b14162e2-073e-4e69-b17d-15a90da0d65f', 'fbcc1bf3-4c54-4bab-956e-2615ba163d0f','347d2bf4-c1ba-4916-916c-460894cc5b80');




