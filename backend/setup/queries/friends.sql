CREATE TABLE IF NOT EXISTS friends
(
    id              SERIAL PRIMARY KEY,
    user_id         INT NOT NULL,
    friend_id       INT NOT NULL,
    invite_date     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    relationship    VARCHAR(20) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_friend FOREIGN KEY (friend_id) REFERENCES users(id),
    CHECK (user_id <> friend_id),
    unique (user_id, friend_id)
);

CREATE INDEX IF NOT EXISTS idx_friends_user_id ON friends(user_id);
CREATE INDEX IF NOT EXISTS idx_friends_friend_id ON friends(friend_id);
CREATE INDEX IF NOT EXISTS idx_inviter_friend_id_relationship ON friends(friend_id, relationship);