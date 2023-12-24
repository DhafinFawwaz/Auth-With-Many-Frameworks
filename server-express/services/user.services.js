import bcrypt from "bcryptjs";
import { generateAccessToken, authenticateToken } from "../middlewares/auth.middlewares.js";
import { db } from "../sql/db.js";

export async function loginUserService({email, password}, callback) {
    const userQuery = await db.get(
        "SELECT * FROM user WHERE email = ?",
        [email],
        (err, data) => {
            if (err) {
                return callback(err);
            }
        }
    );

    if (userQuery === undefined) {
        return callback("User not found");
    }

    const validPassword = await bcrypt.compare(password, userQuery.password);
    if (!validPassword) {
        return callback("Password is not correct");
    }

    console.log(userQuery);
    const token = generateAccessToken({
        id: userQuery.id,
        email: userQuery.email,
    });
    return callback(null, { ...userQuery, token });
    
}

export async function registerUserService(params, callback) {
    if(params.username === undefined) {
        return callback({ message: "Username is required"});
    }
    if(params.password === undefined) {
        return callback({ message: "Password is required"});
    }
    if(params.email === undefined) {
        return callback({ message: "Email is required"});
    }

    const userQuery = await db.get(
        "SELECT * FROM user WHERE email = ?",
        [params.email],
        (err, data) => {
            if (err) {
                return callback(err);
            }
        }
    );

    if (userQuery) {
        return callback({ message: "email is already taken"});
    }

    const salt = await bcrypt.genSalt(10);
    const hashedPassword = await bcrypt.hash(params.password, salt);

    const newUser = await db.run(
        "INSERT INTO user (username, password, email, nim, cookies, created_at, modified_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
        [params.username, hashedPassword, params.email, "", "", new Date(), new Date()],
        (err, data) => {
            console.log(data);
            if (err) {
                return callback(err);
            }
        }
    );
    

    return callback(null, {
        id: newUser.lastID,
        username: params.username,
        password: params.password,
        email: params.email,
    });

}