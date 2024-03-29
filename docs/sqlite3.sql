
/*
 Navicat Premium Data Transfer

 Source Server         : icsCloud.wx
 Source Server Type    : SQLite
 Source Server Version : 3035005
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3035005
 File Encoding         : 65001

 Date: 07/08/2023 13:32:17
*/


PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for detail
-- ----------------------------
DROP TABLE IF EXISTS "detail";
CREATE TABLE "detail" (
  ---
  "entid" text NOT NULL DEFAULT (uuid()),
  "openid" text NOT NULL,
  "id" text NOT NULL,
  "detail" text,
  "timesp" timestamp DEFAULT (DATETIME('now','localtime')),
  "is_valid" integer DEFAULT 1,
  "sid" text,
  PRIMARY KEY (
  "entid", "openid", "id"
  )
)
;

-- ----------------------------
-- Records of detail
-- ----------------------------
INSERT INTO "detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '2c5e4923-3c5b-4ed5-a6ed-a0e4de08f4a4', '{"detail": "断电重启", "description": "k8s 启动失败"}', '2022-06-08 23:15:36', 1, '');
INSERT INTO "detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '1706c4a3-cb71-4422-9e0f-fe96231d8781', '{"detail": "版本不兼容", "description": "dockerd"}', '2022-06-08 23:16:02', 1, '');
INSERT INTO "detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'c0777bb3-dc30-487c-83db-f12c18e16ecc', '{"detail": "...", "description": "test"}', '2022-06-08 23:17:31', 1, '');
INSERT INTO "detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '8d3df1e1-9891-4d23-8435-a62c3e3bf6a1', '{"detail": "订单数据", "description": "订单"}', '2022-12-31 23:59:59', 1, NULL);
INSERT INTO "detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'a0895a01-bd40-407f-ba72-5aba6afdbce8', '{"detail": "过期订单", "description": "订单"}', '2022-06-07 23:59:59', 1, NULL);

-- ----------------------------
-- Table structure for ent
-- ----------------------------
DROP TABLE IF EXISTS "ent";
CREATE TABLE "ent" (
  ---
  "entid" text NOT NULL DEFAULT (uuid()),
  "openid" text NOT NULL,
  "name" text,
  "num" text,
  "oth" text,
  "is_valid" integer DEFAULT 1,
  "sid" text NOT NULL,
  PRIMARY KEY (
  "entid", "openid", "sid"
  )
)
;

-- ----------------------------
-- Records of ent
-- ----------------------------
INSERT INTO "ent" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '西安予兮信息技术服务有限公司', '654321', '', 1, '予兮信息');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS "order";
CREATE TABLE "order" (
  ---
  "entid" text NOT NULL DEFAULT (uuid()),
  "openid" text NOT NULL,
  "ordid" text NOT NULL DEFAULT (uuid()),
  "is_valid" integer DEFAULT 1,
  "sid" text,
  PRIMARY KEY (
  "entid", "openid", "ordid"
  )
)
;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO "order" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '8d3df1e1-9891-4d23-8435-a62c3e3bf6a1', 1, NULL);

-- ----------------------------
-- Table structure for repair
-- ----------------------------
DROP TABLE IF EXISTS "repair";
CREATE TABLE "repair" (
  ---
  "entid" text NOT NULL DEFAULT (uuid()),
  "openid" text NOT NULL,
  "repid" text NOT NULL DEFAULT (uuid()),
  "is_valid" integer DEFAULT 1,
  "sid" text,
  PRIMARY KEY (
  "entid", "openid", "repid"
  )
)
;

-- ----------------------------
-- Records of repair
-- ----------------------------
INSERT INTO "repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'c0777bb3-dc30-487c-83db-f12c18e16ecc', 1, '');
INSERT INTO "repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '2c5e4923-3c5b-4ed5-a6ed-a0e4de08f4a4', 1, '');
INSERT INTO "repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '1706c4a3-cb71-4422-9e0f-fe96231d8781', 1, '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  ---
  "openid" text NOT NULL,
  "session_key" text,
  "unionid" text,
  "avatar_url" text,
  "is_authorization" integer DEFAULT false,
  "entid" text DEFAULT (uuid()),
  "is_valid" integer DEFAULT 1,
  PRIMARY KEY (
  "openid"
  )
)
;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "user" VALUES ('xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'eJ5kLzxvyX4ha98Sqc1Wiw==', '', 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL4Yf2awot0oXO6wgOmZSfTg0ybqXqWAypuhrh7d3shGwd2B1zDnVVVP0aWmFdX0icO99ssMNGtfd./132', 1, '001d6fdb-79b5-41ff-af29-99a0e8618c37', 1);

PRAGMA foreign_keys = true;


/* ----------------------------
-- 
-- ----------------------------
-- _________           .___       --
-- \_   ___ \ __ __  __| _/ ___   --
-- /    \  \/|  |  \/ __ |/  _ \  --
-- \     \___|  |  / / / (  <_> ) --
--  \______  /____/\_____|\____/  --
--         \/           \/        --
-- ----------------------------
 ---------------------------- */

-- *****

