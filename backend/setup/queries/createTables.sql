-- #####################################################################################################################
-- YGO Related
-- #####################################################################################################################

-- Table: cards
-- This table contains all information about cards.
CREATE TABLE IF NOT EXISTS cards
(
    id        INTEGER NOT NULL,
    name      TEXT,
    type      TEXT,
    "desc"    TEXT,
    atk       INTEGER,
    def       INTEGER,
    level     INTEGER,
    race      TEXT,
    attribute TEXT,
    sets      TEXT,
    CONSTRAINT cards_pkey PRIMARY KEY (id)
);

-- Table: card_sets
-- This table contains all data for the card sets.
CREATE TABLE IF NOT EXISTS card_sets
(
    set_code        TEXT,
    set_name        TEXT,
    set_rarity      TEXT,
    set_rarity_code TEXT,
    CONSTRAINT card_sets_pkey PRIMARY KEY (set_name)
);

-- #####################################################################################################################
-- User Management
-- #####################################################################################################################

-- Table: users
-- This table contains all data for users.
CREATE TABLE IF NOT EXISTS users
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL,
    password_hash VARCHAR(60)  NOT NULL,
    display_name  VARCHAR(255) NOT NULL,
    is_admin      BOOLEAN   DEFAULT FALSE,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
);

-- Table: friends
-- This table contains the friend relation between two users.
CREATE TABLE IF NOT EXISTS friends
(
    id           SERIAL PRIMARY KEY,
    user_id      INT         NOT NULL,
    friend_id    INT         NOT NULL,
    invite_date  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    relationship VARCHAR(20) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_friend FOREIGN KEY (friend_id) REFERENCES users (id),
    CHECK (user_id <> friend_id),
    unique (user_id, friend_id)
);

CREATE INDEX IF NOT EXISTS idx_friends_user_id ON friends (user_id);
CREATE INDEX IF NOT EXISTS idx_friends_friend_id ON friends (friend_id);
CREATE INDEX IF NOT EXISTS idx_inviter_friend_id_relationship ON friends (friend_id, relationship);

-- #####################################################################################################################
-- Draft Related
-- #####################################################################################################################

-- Table: drafts
-- This table contains all information about drafts.
CREATE TABLE IF NOT EXISTS drafts
(
    id                   SERIAL PRIMARY KEY,
    challenger_id        INT REFERENCES users (id) NOT NULL,
    receiver_id          INT REFERENCES users (id) NOT NULL,
    current_round_number INT                       NOT NULL,
    maximum_round_number INT                       NOT NULL,
    challenge_date       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    winner_user_id       INT                       NOT NULL,
    status               VARCHAR(20)               NOT NULL,
    settings             JSON                      NOT NULL,
    CHECK (challenger_id <> receiver_id)
);

-- Table: draft_rounds
-- This table contains all information about one draft round each row.
CREATE TABLE IF NOT EXISTS draft_rounds
(
    id             SERIAL PRIMARY KEY,
    draft_id       INT REFERENCES drafts (id) NOT NULL,
    round_number   INT                        NOT NULL,
    status         VARCHAR(20)                NOT NULL,
    winner_user_id INT NOT NULL,
    CONSTRAINT unique_round_per_draft UNIQUE (draft_id, round_number)
);

-- Table: draft_deck
-- This table contains the decks drafted in every round for each user.
CREATE TABLE IF NOT EXISTS draft_deck
(
    id       SERIAL PRIMARY KEY,
    round_id INT REFERENCES draft_rounds (id) NOT NULL,
    user_id  INT REFERENCES users (id)        NOT NULL,
    deck     TEXT                             NOT NULL
);