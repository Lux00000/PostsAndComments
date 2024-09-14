CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id VARCHAR(255) NOT NULL,
    allow_comments BOOLEAN NOT NULL DEFAULT TRUE
    );

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL
    FOREIGN KEY (post) REFERENCES posts(id),
    parent_comment_id INT
    FOREIGN KEY (parent_comment_id) REFERENCES comments(id),
    author_id VARCHAR(255) NOT NULL,
    text TEXT(2000) NOT NULL
    );