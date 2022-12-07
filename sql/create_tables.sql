CREATE TABLE users(
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_name VARCHAR(255) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP
);
CREATE TABLE user_languages(
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_idx BIGINT NOT NULL,
    use_language VARCHAR(50) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_idx) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE language_exps(
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    language_name VARCHAR(255) NOT NULL,
    exp_ratio FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT USER_IDX UNIQUE (user_name);
ALTER TABLE user_languages ADD CONSTRAINT USER_LANGUAGE_IDX UNIQUE (user_idx, use_language);
