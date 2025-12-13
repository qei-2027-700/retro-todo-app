-- ユーザーテーブル作成
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    -- 将来のOIDC移行用フィールド
    external_id VARCHAR(255),
    provider VARCHAR(50) DEFAULT 'local',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- インデックス作成
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_external_id ON users(external_id);

-- テスト用ユーザー追加（パスワード: `password123`）
-- bcryptハッシュ: $2a$10$uoZqVUeLQWgDuFO24r5Eo.5v63qs0gq0W03brjMccY9rT8kXTzGS2
INSERT INTO users (username, email, password_hash) VALUES
('testuser', 'test@example.com', '$2a$10$uoZqVUeLQWgDuFO24r5Eo.5v63qs0gq0W03brjMccY9rT8kXTzGS2')
ON CONFLICT (username) DO NOTHING;
