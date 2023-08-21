CREATE TABLE user_book (
                           user_id integer references "user"(id),
                           book_id integer references book(id)
);