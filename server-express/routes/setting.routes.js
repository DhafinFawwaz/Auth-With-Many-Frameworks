import express from "express";
import { editUser } from "../controllers/setting.controllers.js";
import { authenticateToken }from "../middlewares/auth.middlewares.js";

const router = express.Router();

router.post("/edit", authenticateToken, editUser);

export default router;