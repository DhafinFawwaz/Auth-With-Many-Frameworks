export const dropUserTable = 'DROP TABLE IF EXISTS user';
export const dropScheduleTable = 'DROP TABLE IF EXISTS schedule';

export const createUserTable = `CREATE TABLE "user"(
    "id" INTEGER PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "nim" VARCHAR(255) NOT NULL,
    "cookies" VARCHAR(255) NOT NULL,
    "created_at" DATETIME NOT NULL,
    "modified_at" DATETIME NOT NULL
)`;

export const createScheduleTable = `CREATE TABLE "schedule"(
    "id" INTEGER PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "schedule_at" DATETIME NOT NULL,
    "schedule_status" TINYINT NOT NULL,
    "matkul" VARCHAR(255) NOT NULL,
    "timerange" VARCHAR(255) NOT NULL,
    "absen_status" TINYINT NOT NULL
)`;