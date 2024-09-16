CREATE TABLE IF NOT EXISTS Posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id VARCHAR(255) NOT NULL,
    allow_comments BOOLEAN NOT NULL DEFAULT TRUE
    );

CREATE TABLE IF NOT EXISTS Comments (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    parent_comment_id INT,
    author_id VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES Posts(id),
    FOREIGN KEY (parent_comment_id) REFERENCES Comments(id)
    );
