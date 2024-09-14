CREATE DATABASE postsandcomments;

\c postsandcomments;

CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       author_id VARCHAR(255) NOT NULL,
                       allow_comments BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          post_id INT NOT NULL REFERENCES posts(id),
                          parent_comment_id INT REFERENCES comments(id),
                          author_id VARCHAR(255) NOT NULL,
                          text TEXT NOT NULL
);