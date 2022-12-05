CREATE TABLE user(
    idx BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT ,
    user_name VARCHAR(255) NOT NULL ,
    email VARCHAR(255) NOT NULL
);
CREATE TABLE user_language(
    idx BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT ,
    user_idx BIGINT NOT NULL,
    use_language VARCHAR(50) NOT NULL,
    amount BIGINT NOT NULL,
    FOREIGN KEY (user_idx) REFERENCES user(idx) ON DELETE CASCADE
);

CREATE TABLE language_exp(
    idx BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    language_name VARCHAR(255) NOT NULL,
    exp_ratio FLOAT NOT NULL
);

ALTER TABLE user ADD CONSTRAINT USER_IDX UNIQUE (email);
ALTER TABLE user_language ADD CONSTRAINT USER_LANGUAGE_IDX UNIQUE (user_idx, use_language);
