CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    author_name TEXT
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    author_id INT,
    title TEXT,
    FOREIGN KEY (author_id) REFERENCES author(id)
);

CREATE TABLE genre (
    id SERIAL PRIMARY KEY,
    genre_name TEXT
);

CREATE TABLE book_genre (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL,
    genre_id INT NOT NULL,
    FOREIGN KEY (book_id) REFERENCES book (id),
    FOREIGN KEY (genre_id) REFERENCES genre (id)
);

INSERT INTO author (author_name) VALUES ('Yann Martel');
INSERT INTO author (author_name) VALUES ('Alexandre Dumas');
INSERT INTO author (author_name) VALUES ('George Orwell');

INSERT INTO genre (genre_name) VALUES ('Action');
INSERT INTO genre (genre_name) VALUES ('Science Fiction');

INSERT INTO book (title, author_id) SELECT 'Life Of Pi', id from author where author_name = 'Yann Martel';
INSERT INTO book (title, author_id) SELECT 'The Three Musketeers', id from author where author_name = 'Alexandre Dumas';
INSERT INTO book (title, author_id) SELECT '1984', id from author where author_name = 'George Orwell';

INSERT INTO book_genre (book_id, genre_id) VALUES (
    (SELECT id from book where title = 'Life Of Pi'),
    (SELECT id from genre where genre_name = 'Action')
);
INSERT INTO book_genre (book_id, genre_id) VALUES (
    (SELECT id from book where title = 'The Three Musketeers'),
    (SELECT id from genre where genre_name = 'Action')
);
INSERT INTO book_genre (book_id, genre_id) VALUES (
    (SELECT id from book where title = '1984'),
    (SELECT id from genre where genre_name = 'Science Fiction')
);
