/*
 Navicat Premium Data Transfer

 Source Server         : psql-127.0.0.1
 Source Server Type    : PostgreSQL
 Source Server Version : 130006
 Source Host           : 127.0.0.1:5432
 Source Catalog        : icsCloud.wx
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130006
 File Encoding         : 65001

 Date: 05/10/2022 00:15:56
*/


-- ----------------------------
-- DisEnable UUID EXTEND
-- ----------------------------
-- DROP EXTENSION "uuid-ossp";

-- ----------------------------
-- Table structure for detail
-- ----------------------------
DROP TABLE IF EXISTS "public"."detail";
CREATE TABLE "public"."detail" (
  "entid" uuid NOT NULL,
  "openid" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "id" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "detail" jsonb,
  "timesp" timestamp(0) DEFAULT CURRENT_TIMESTAMP,
  "is_valid" bool DEFAULT true,
  "sid" varchar COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of detail
-- ----------------------------
INSERT INTO "public"."detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '2c5e4923-3c5b-4ed5-a6ed-a0e4de08f4a4', '{"detail": "断电重启", "description": "k8s 启动失败"}', '2022-06-08 23:15:36', 't', '');
INSERT INTO "public"."detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '1706c4a3-cb71-4422-9e0f-fe96231d8781', '{"detail": "版本不兼容", "description": "dockerd"}', '2022-06-08 23:16:02', 't', '');
INSERT INTO "public"."detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'c0777bb3-dc30-487c-83db-f12c18e16ecc', '{"detail": "...", "description": "test"}', '2022-06-08 23:17:31', 't', '');
INSERT INTO "public"."detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '8d3df1e1-9891-4d23-8435-a62c3e3bf6a1', '{"detail": "订单数据", "description": "订单"}', '2022-12-31 23:59:59', 't', NULL);
INSERT INTO "public"."detail" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'a0895a01-bd40-407f-ba72-5aba6afdbce8', '{"detail": "过期订单", "description": "订单"}', '2022-06-07 23:59:59', 't', NULL);

-- ----------------------------
-- Table structure for ent
-- ----------------------------
DROP TABLE IF EXISTS "public"."ent";
CREATE TABLE "public"."ent" (
  "entid" uuid NOT NULL,
  "openid" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar COLLATE "pg_catalog"."default",
  "num" varchar COLLATE "pg_catalog"."default",
  "oth" varchar COLLATE "pg_catalog"."default",
  "is_valid" bool DEFAULT true,
  "sid" varchar COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of ent
-- ----------------------------
INSERT INTO "public"."ent" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '西安予兮信息技术服务有限公司', '654321', '', 't', '予兮信息');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS "public"."order";
CREATE TABLE "public"."order" (
  "entid" uuid NOT NULL,
  "openid" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "ordid" varchar COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "is_valid" bool DEFAULT true,
  "sid" varchar COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO "public"."order" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '8d3df1e1-9891-4d23-8435-a62c3e3bf6a1', 't', NULL);

-- ----------------------------
-- Table structure for repair
-- ----------------------------
DROP TABLE IF EXISTS "public"."repair";
CREATE TABLE "public"."repair" (
  "entid" uuid NOT NULL,
  "openid" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "repid" varchar COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "is_valid" bool DEFAULT true,
  "sid" varchar COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of repair
-- ----------------------------
INSERT INTO "public"."repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'c0777bb3-dc30-487c-83db-f12c18e16ecc', 't', '');
INSERT INTO "public"."repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '2c5e4923-3c5b-4ed5-a6ed-a0e4de08f4a4', 't', '');
INSERT INTO "public"."repair" VALUES ('001d6fdb-79b5-41ff-af29-99a0e8618c37', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxx', '1706c4a3-cb71-4422-9e0f-fe96231d8781', 't', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "openid" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "session_key" varchar COLLATE "pg_catalog"."default",
  "unionid" varchar COLLATE "pg_catalog"."default",
  "avatar_url" varchar COLLATE "pg_catalog"."default",
  "is_authorization" bool DEFAULT false,
  "entid" varchar COLLATE "pg_catalog"."default" DEFAULT uuid_generate_v4(),
  "is_valid" bool DEFAULT true
)
;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO "public"."user" VALUES ('xxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'eJ5kLzxvyX4ha98Sqc1Wiw==', '', 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL4Yf2awot0oXO6wgOmZSfTg0ybqXqWAypuhrh7d3shGwd2B1zDnVVVP0aWmFdX0icO99ssMNGtfdw/132', 't', '001d6fdb-79b5-41ff-af29-99a0e8618c37', 't');

-- ----------------------------
-- Primary Key structure for table detail
-- ----------------------------
ALTER TABLE "public"."detail" ADD CONSTRAINT "detail_pkey" PRIMARY KEY ("entid", "openid", "id");

-- ----------------------------
-- Primary Key structure for table ent
-- ----------------------------
ALTER TABLE "public"."ent" ADD CONSTRAINT "ent_pkey" PRIMARY KEY ("entid", "openid", "sid");

-- ----------------------------
-- Primary Key structure for table order
-- ----------------------------
ALTER TABLE "public"."order" ADD CONSTRAINT "order_pkey" PRIMARY KEY ("entid", "openid", "ordid");

-- ----------------------------
-- Primary Key structure for table repair
-- ----------------------------
ALTER TABLE "public"."repair" ADD CONSTRAINT "repair_pkey" PRIMARY KEY ("entid", "openid", "repid");

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("openid");

-- ----------------------------
-- Enable UUID EXTEND
-- ----------------------------
CREATE EXTENSION "uuid-ossp";


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