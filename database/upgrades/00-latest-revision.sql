-- v0 -> v3: Latest revision

CREATE TABLE "user" (
    -- only: postgres
    rowid BIGINT  PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    -- only: sqlite
    rowid INTEGER PRIMARY KEY,

    mxid     TEXT NOT NULL UNIQUE,
    phone_id TEXT UNIQUE,
    session  jsonb,

    self_participant_ids jsonb NOT NULL DEFAULT '[]',

    management_room TEXT,
    space_room      TEXT,

    access_token TEXT
);

CREATE TABLE puppet (
    id               TEXT    NOT NULL,
    receiver         BIGINT  NOT NULL,
    phone            TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    name_set         BOOLEAN NOT NULL DEFAULT false,
    avatar_id        TEXT    NOT NULL,
    avatar_mxc       TEXT    NOT NULL,
    avatar_set       BOOLEAN NOT NULL DEFAULT false,
    contact_info_set BOOLEAN NOT NULL DEFAULT false,

    PRIMARY KEY (id, receiver),
    CONSTRAINT puppet_user_fkey    FOREIGN KEY (receiver) REFERENCES "user"(rowid) ON DELETE CASCADE,
    CONSTRAINT puppet_phone_unique UNIQUE (phone, receiver)
);

CREATE TABLE portal (
    id         TEXT    NOT NULL,
    receiver   BIGINT  NOT NULL,
    self_user  TEXT,
    other_user TEXT,
    type       INTEGER NOT NULL,
    mxid       TEXT    UNIQUE,
    name       TEXT    NOT NULL,
    name_set   BOOLEAN NOT NULL DEFAULT false,
    avatar_id  TEXT    NOT NULL,
    avatar_mxc TEXT    NOT NULL,
    avatar_set BOOLEAN NOT NULL DEFAULT false,
    encrypted  BOOLEAN NOT NULL DEFAULT false,
    in_space   BOOLEAN NOT NULL DEFAULT false,

    PRIMARY KEY (id, receiver),
    CONSTRAINT portal_user_fkey   FOREIGN KEY (receiver) REFERENCES "user"(rowid) ON DELETE CASCADE,
    CONSTRAINT portal_puppet_fkey FOREIGN KEY (other_user, receiver) REFERENCES puppet(id, receiver) ON DELETE CASCADE
);

CREATE TABLE message (
    conv_id       TEXT   NOT NULL,
    conv_receiver BIGINT NOT NULL,
    id            TEXT   NOT NULL,
    mxid          TEXT   NOT NULL,
    mx_room       TEXT   NOT NULL,
    sender        TEXT   NOT NULL,
    timestamp     BIGINT NOT NULL,
    status        jsonb  NOT NULL,

    PRIMARY KEY (conv_receiver, id),
    CONSTRAINT message_portal_fkey FOREIGN KEY (conv_id, conv_receiver) REFERENCES portal(id, receiver) ON DELETE CASCADE,
    CONSTRAINT message_mxid_unique UNIQUE (mxid)
);
CREATE INDEX message_conv_timestamp_idx ON message(conv_id, conv_receiver, timestamp);

CREATE TABLE reaction (
    conv_id       TEXT   NOT NULL,
    conv_receiver BIGINT NOT NULL,
    msg_id        TEXT   NOT NULL,
    sender        TEXT   NOT NULL,
    reaction      TEXT   NOT NULL,
    mxid          TEXT   NOT NULL,

    PRIMARY KEY (conv_receiver, msg_id, sender),
    CONSTRAINT reaction_message_fkey FOREIGN KEY (conv_receiver, msg_id) REFERENCES message(conv_receiver, id) ON DELETE CASCADE,
    CONSTRAINT reaction_mxid_unique  UNIQUE (mxid)
)
