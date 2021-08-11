CREATE TABLE IF NOT EXISTS questions_types (
    id serial PRIMARY KEY,
    type VARCHAR(15) UNIQUE
);

INSERT INTO questions_types (type) VALUES ('single_choice');
INSERT INTO questions_types (type) VALUES ('multiple_choice');

CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY,
    type VARCHAR(15) NOT NULL ,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL UNIQUE,
    FOREIGN KEY (type) REFERENCES questions_types(type)
);

CREATE TABLE IF NOT EXISTS answers (
    id serial PRIMARY KEY,
    question_id int NOT NULL,
    created_at timestamp NOT NULL DEFAULT  CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL UNIQUE,
    is_correct BOOL NOT NULL DEFAULT FALSE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);