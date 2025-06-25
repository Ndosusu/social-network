CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER NOT NULL UNIQUE,
    "uuid" VARCHAR(36) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(60) NOT NULL UNIQUE,
    "first_name" VARCHAR(64) NOT NULL,
    "last_name" VARCHAR(64) NOT NULL,
    "date_birth" VARCHAR(10) NOT NULL,
    "avatar" VARCHAR(128) NOT NULL,
    "nick_name" VARCHAR(64) NOT NULL,
    "about" VARCHAR(512) NOT NULL,
    "date_creation" VARCHAR(25) NOT NULL,
    "private_mode" BOOLEAN NOT NULL,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "groups" (
    "id" INTEGER NOT NULL UNIQUE,
    "admin_id" INTEGER NOT NULL,
    "title" VARCHAR(64) NOT NULL,
    "about" VARCHAR(512) NOT NULL,
    PRIMARY KEY("id"),
    FOREIGN KEY ("admin_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "events" (
    "id" INTEGER NOT NULL UNIQUE,
    "group_id" INTEGER NOT NULL,
    "title" VARCHAR(64) NOT NULL,
    "about" VARCHAR(512) NOT NULL,
    "date_schedule" VARCHAR(25) NOT NULL,
    PRIMARY KEY("id"),
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "posts" (
    "id" INTEGER NOT NULL UNIQUE,
    "author_id" INTEGER NOT NULL,
    "message" VARCHAR(1024) NOT NULL,
    "image" VARCHAR(128),
    "privacy_mode" INTEGER NOT NULL,
    "group_id" INTEGER,
    PRIMARY KEY("id"),
    FOREIGN KEY ("author_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "comments" (
    "id" INTEGER NOT NULL UNIQUE,
    "post_id" INTEGER NOT NULL,
    "author_id" INTEGER NOT NULL,
    "message" VARCHAR(1024) NOT NULL,
    "image" VARCHAR(128),
    "group_id" INTEGER,
    PRIMARY KEY("id"),
    FOREIGN KEY ("author_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("post_id") REFERENCES "posts"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "likes" (
    "id" INTEGER NOT NULL UNIQUE,
    "user_id" INTEGER NOT NULL,
    "post_id" INTEGER,
    "comment_id" INTEGER,
    PRIMARY KEY("id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("post_id") REFERENCES "posts"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("comments_id") REFERENCES "comments"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "chats" (
    "id" INTEGER NOT NULL UNIQUE,
    "group_id" INTEGER UNIQUE,
    "user1_id" INTEGER,
    "user2_id" INTEGER,
    PRIMARY KEY("id"),
    FOREIGN KEY ("user1_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user2_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "chat_log" (
    "id" INTEGER NOT NULL UNIQUE,
    "chat_id" INTEGER NOT NULL,
    "author_id" INTEGER NOT NULL,
    "log" VARCHAR(1024) NOT NULL,
    "date" VARCHAR(25) NOT NULL,
    PRIMARY KEY("id"),
    FOREIGN KEY ("author_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("chat_id") REFERENCES "chats"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "notifications" (
    "id" INTEGER NOT NULL UNIQUE,
    "type" INTEGER NOT NULL,
    "user_to" INTEGER NOT NULL,
    "user_from" INTEGER NOT NULL,
    "group_id" INTEGER,
    "event_id" INTEGER,
    PRIMARY KEY("id"),
    FOREIGN KEY ("user_to") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user_from") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("event_id") REFERENCES "events"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "follow_rel" (
    "user_from" INTEGER NOT NULL,
    "user_to" INTEGER NOT NULL,
    FOREIGN KEY ("user_from") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user_to") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "group_member_rel" (
    "group_id" INTEGER NOT NULL,
    "member_id" INTEGER NOT NULL,
    FOREIGN KEY ("member_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("group_id") REFERENCES "groups"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "privacy_post_rel" (
    "post_id" INTEGER NOT NULL,
    "follower_id" INTEGER NOT NULL,
    FOREIGN KEY ("follower_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("post_id") REFERENCES "posts"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS "event_members_rel" (
    "event_id" INTEGER NOT NULL,
    "member_id" INTEGER NOT NULL,
    FOREIGN KEY ("member_id") REFERENCES "users"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("event_id") REFERENCES "events"("id")
        ON UPDATE NO ACTION ON DELETE NO ACTION
);
