INSERT INTO authors (id, first_name) VALUES (0, 'автор не указан');
 -- поскольку id установлен принудительно, то необходимо изменить начало послеовательности
ALTER SEQUENCE authors_id_seq RESTART WITH 100;

INSERT INTO publishers (id, name) VALUES (0, 'издательство не указано');
ALTER SEQUENCE publishers_id_seq RESTART WITH 100;

-- insert into books (title, year) VALUES ('Book', 2050); - что будет в результате выполнения?