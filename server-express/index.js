import express from "express";
import auth from "./routes/auth.routes.js";
import setting from "./routes/setting.routes.js";
import { authenticateToken }from "./middlewares/auth.middlewares.js";
import { errorHandler } from "./middlewares/error.middlewares.js";
import dotenv from "dotenv";
import { expressjwt } from "express-jwt";
import { db } from "./sql/db.js";

const app = express();
dotenv.config();

app.use(expressjwt({
    secret: process.env.JWT_SECRET,
    algorithms: ['HS256'],
}).unless({
    path: [
        '/auth/login',
        '/auth/register',
        '/oauth2callback'
    ]
}));

app.use(express.json());

app.use("/auth", auth);
app.use("/setting", setting)

app.use(errorHandler);


app.listen(process.env.PORT, () => {
    console.log(`Server is running on port ${process.env.PORT}`)
});
