import { db } from "../sql/db.js";
import jwt from 'jsonwebtoken';

export async function editUser(req, res, next) {

    console.log(`Editing ${req.body}`);

    console.log(JSON.stringify(req.body));

    const editQuery = await db.run(
        `
        UPDATE user
        SET username=?, password=?, nim=?, cookies=?, modified_at=?
        WHERE email=?
        `,
        [req.body.username, req.body.password, req.body.nim, req.body.cookies, new Date(), req.body.email]
    );

    res.status(200).json({
        message: "success",
        data: req.body
    });
}

