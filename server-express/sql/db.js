import dotenv from 'dotenv';
import sqlite3 from 'sqlite3'
import { open } from 'sqlite'

dotenv.config();
export const db = await open({
    filename: process.env.DATABASE_PATH,
    driver: sqlite3.Database
})
console.log('Connected to Database!')