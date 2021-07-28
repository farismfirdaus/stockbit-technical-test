/*
 Navicat Premium Data Transfer

 Source Server         : postgres-local
 Source Server Type    : PostgreSQL
 Source Server Version : 130002
 Source Host           : localhost:5432
 Source Catalog        : logging
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130002
 File Encoding         : 65001

 Date: 28/07/2021 23:52:10
*/


-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS "public"."log";
CREATE TABLE "public"."log" (
  "id" int4 NOT NULL GENERATED ALWAYS AS IDENTITY (
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
),
  "timestamp" timestamp(6),
  "method" varchar(10) COLLATE "pg_catalog"."default",
  "request" text COLLATE "pg_catalog"."default",
  "response" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Primary Key structure for table log
-- ----------------------------
ALTER TABLE "public"."log" ADD CONSTRAINT "log_pkey" PRIMARY KEY ("id");
