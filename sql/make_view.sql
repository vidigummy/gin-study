CREATE VIEW USER_LANGUAGE_EXP AS
SELECT u.user_idx, u.use_language, FLOOR(u.amount*l.exp_ratio) AS exp FROM user_languages AS u JOIN language_exps AS l ON u.use_language = l.language_name;