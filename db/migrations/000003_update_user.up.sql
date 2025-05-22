-- Avvaldan jadval bo'lsa:
ALTER TABLE user_tokens ADD CONSTRAINT unique_user_id UNIQUE (user_id);
