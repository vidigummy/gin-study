DROP VIEW USER_LANGUAGE_EXP
DROP TABLE language_exps;
DROP TABLE user_languages;
DROP TABLE users;

CREATE TABLE users(
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_name VARCHAR(255) NOT NULL ,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP
);
CREATE TABLE user_languages(
    idx BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_idx BIGINT NOT NULL,
    use_language VARCHAR(50) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_idx) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE language_exps(
    idx BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    language_name VARCHAR(255) NOT NULL,
    exp_ratio FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    updated_at TIMESTAMP not null default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT USER_IDX UNIQUE (email);
ALTER TABLE user_languages ADD CONSTRAINT USER_LANGUAGE_IDX UNIQUE (user_idx, use_language);


INSERT INTO RPG.users (user_name, email) VALUES ('fakeVidi', 'fakeVidi@gmail.com');
INSERT INTO RPG.user_languages (user_idx, use_language, amount) VALUES (1, 'Java', 10000000);
INSERT INTO RPG.user_languages (user_idx, use_language, amount) VALUES (1, 'Go', 400000);
INSERT INTO RPG.user_languages (user_idx, use_language, amount) VALUES (1, 'Python', 1000000);

INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('C', 0.5);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('Java', 0.6);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('Python', 1.0);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('Go', 0.8);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('C++', 0.5);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('CSS', 0.1);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('HTML', 0.1);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('Dockerfile', 2.0);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('BatchFile', 0.3);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('JavaScript', 0.9);
INSERT INTO RPG.language_exps (language_name, exp_ratio) VALUES ('TypeScript', 0.9);

CREATE VIEW USER_LANGUAGE_EXP AS
SELECT u.user_idx, u.use_language, FLOOR(u.amount*l.exp_ratio) AS exp FROM user_languages AS u JOIN language_exps AS l ON u.use_language = l.language_name;