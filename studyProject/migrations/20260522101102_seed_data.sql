-- +goose Up

INSERT INTO courses(name, description)
VALUES
    ('Golang Developer', 'Learn Go'),
    ('Java Developer', 'Learn Java');

INSERT INTO chapters(name, description, "order", course_id)
VALUES
    ('Control structures', 'Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic. Unlike many other languages, Go lacks a dedicated while keyword, instead utilizing a versatile for loop for all iteration needs', 1, 1),
    ('Variables', 'In Go (Golang), variables are explicitly declared containers used by the compiler to store data and check for type correctness. Go is statically typed, meaning once a variable is declared with a specific type, it cannot hold values of a different type. ', 2, 1);

INSERT INTO lessons(name, description, content, "order", chapter_id)
VALUES
    ('If-else Statement in Golang', 'In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.', 'Go language', 1, 1);
-- +goose Down

DELETE FROM lessons;
DELETE FROM chapters;
DELETE FROM courses;
