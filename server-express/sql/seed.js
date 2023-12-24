import { createUserTable, dropUserTable,
	createScheduleTable, dropScheduleTable
} from './sql.js';
import dotenv from 'dotenv';
import sqlite3 from 'sqlite3'
import { open } from 'sqlite'

const db = await open({
    filename: 'db/siabsen.db',
    driver: sqlite3.Database
})

await initialize();


async function initialize() {
	try {
		//clear the existing records
        console.log('Initializing...\n');
        console.log('dropping all tables...');
        await db.exec(dropUserTable);
        console.log('***dropped user table***');
		await db.exec(dropScheduleTable);
		console.log('***dropped schedule table***\n');

		//create the tables
        console.log('creating all tables...');
        await db.exec(createUserTable);
        console.log('***created user table***');
		await db.exec(createScheduleTable);
		console.log('***created schedule status table***');
    
	}catch(err){
		console.error(err);
	}
}
process.exit();
