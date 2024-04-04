CREATE TABLE IF NOT EXISTS drafts
(
    id          SERIAL PRIMARY KEY,
    challengeID INT REFERENCES draft_challenges (id) NOT NULL,
    userID1     INT REFERENCES users (id)            NOT NULL,
    userID2     INT REFERENCES users (id)            NOT NULL,
    status      VARCHAR(20)                          NOT NULL
);

CREATE TABLE IF NOT EXISTS draft_rounds
(
    id           SERIAL PRIMARY KEY,
    draftID      INT REFERENCES drafts (id) NOT NULL,
    roundNumber  INT                        NOT NULL,
    status       VARCHAR(20)                NOT NULL,
    deck         TEXT[]                     NOT NULL,
    winnerUserID INT REFERENCES users (id),
    CONSTRAINT unique_round_per_draft UNIQUE (draftID, roundNumber)
);

CREATE TABLE IF NOT EXISTS draft_deck
(
    id      SERIAL PRIMARY KEY,
    roundID INT REFERENCES draft_rounds (id) NOT NULL,
    userID  INT REFERENCES users (id)        NOT NULL,
    deck    TEXT[]                           NOT NULL
);