-- +goose Up

CREATE TABLE courses (
                         id BIGSERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT,
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE chapters (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          "order" INTEGER NOT NULL,

                          course_id BIGINT NOT NULL,

                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW(),

                          CONSTRAINT fk_course
                              FOREIGN KEY(course_id)
                                  REFERENCES courses(id)
                                  ON DELETE CASCADE
);

CREATE TABLE lessons (
                         id BIGSERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         description TEXT,
                         content TEXT,
                         "order" INTEGER NOT NULL,

                         chapter_id BIGINT NOT NULL,

                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW(),

                         CONSTRAINT fk_chapter
                             FOREIGN KEY(chapter_id)
                                 REFERENCES chapters(id)
                                 ON DELETE CASCADE
);

-- +goose Down

DROP TABLE lessons;
DROP TABLE chapters;
DROP TABLE courses;
