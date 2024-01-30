import jwt from 'jsonwebtoken';

export function authenticateToken(req, res, next) {
    const authHeader = req.headers['authorization'];
    
    const token = authHeader && authHeader.split(' ')[1];
    if (token == null) return res.sendStatus(401);
    jwt.verify(token, process.env.JWT_SECRET, (err, data) => {
        if (err) return res.sendStatus(403);
        req.body.email = data.email;
        req.body.id = data.id;
        next();
    });
}

export function generateAccessToken(data) {
    return jwt.sign(data, process.env.JWT_SECRET);
}
