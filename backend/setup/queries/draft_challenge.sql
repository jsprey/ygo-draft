CREATE TABLE IF NOT EXISTS draft_challenge
(
    id             SERIAL PRIMARY KEY,
    challenger_id  INT         NOT NULL,
    receiver_id    INT         NOT NULL,
    challenge_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status         VARCHAR(20) NOT NULL,
    settings       JSON        NOT NULL,
    CONSTRAINT fk_challenger FOREIGN KEY (challenger_id) REFERENCES users (id),
    CONSTRAINT fk_receiver FOREIGN KEY (receiver_id) REFERENCES users (id),
    CHECK (challenger_id <> receiver_id)
);

CREATE INDEX IF NOT EXISTS idx_draft_challenge_challenger_id ON draft_challenge (challenger_id);
CREATE INDEX IF NOT EXISTS idx_draft_challenge_receiver_id ON draft_challenge (receiver_id);