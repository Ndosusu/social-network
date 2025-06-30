PRAGMA foreign_keys = OFF;

DROP TABLE IF EXISTS "event_members_rel";
DROP TABLE IF EXISTS "privacy_post_rel";
DROP TABLE IF EXISTS "group_member_rel";
DROP TABLE IF EXISTS "follow_rel";
DROP TABLE IF EXISTS "notifications";
DROP TABLE IF EXISTS "chat_log";
DROP TABLE IF EXISTS "chats";
DROP TABLE IF EXISTS "likes";
DROP TABLE IF EXISTS "comments";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "events";
DROP TABLE IF EXISTS "groups";
DROP TABLE IF EXISTS "users";

PRAGMA foreign_keys = ON;