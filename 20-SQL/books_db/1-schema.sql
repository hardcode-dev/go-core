/*
    Схема учебной БД "Книги".
    Имена таблиц во множественном числе.
*/

/*
    Удаляем таблицы, если они существуют.
    Удаление производится в обратном относительно создания порядке.
*/ 
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS publishers;

/*
    Создаём таблицы БД.
    Сначала создаются зависимые таблицы.
*/

-- authors - писатели
CREATE TABLE authors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

-- publishers - издатели
CREATE TABLE publishers (
    id SERIAL PRIMARY KEY, -- первичный ключ
    name TEXT NOT NULL,
    website TEXT NOT NULL DEFAULT ''
);

-- books - книги
CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    isbn BIGINT UNIQUE, -- ISBN
    title TEXT NOT NULL, -- название
    year INTEGER DEFAULT 0, -- год выпуска (максимум текущий + 10)
    public_domain BOOLEAN DEFAULT FALSE, -- является ли общественным достоянием
    author_id INTEGER NOT NULL REFERENCES authors(id) DEFAULT 0,
    publisher_id INTEGER REFERENCES publishers(id) DEFAULT 0,
    price INTEGER DEFAULT 0 CHECK (price >= 0)
);

-- функция-триггер для проверки года выпуска книги
CREATE OR REPLACE FUNCTION check_book_year()
  RETURNS TRIGGER AS $$
BEGIN
    IF NEW.year < (SELECT (extract(year from current_date) + 10))
        THEN RETURN NEW;
        ELSE RAISE EXCEPTION 'Invalid book year'; --RETURN NULL;
    END IF;
END;
$$ LANGUAGE plpgsql;
-- регистрация тригера для таблицы
CREATE TRIGGER check_book_year BEFORE INSERT OR UPDATE ON books 
FOR EACH ROW EXECUTE PROCEDURE check_book_year();