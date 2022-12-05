CREATE VIEW USER_LANGUAGE_EXP AS
SELECT u.user_idx, u.use_language, FLOOR(u.amount*l.exp_ratio) AS exp FROM user_language AS u JOIN language_exp AS l ON u.use_language = l.language_name;