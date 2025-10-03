const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');
const dotenv = require("dotenv");
const path = require("path");
const { access } = require('fs');
dotenv.config({ path: path.resolve(__dirname, "../.env") });

const token_jwt = process.env.ACCESS_JWT_SECRET

const registerUser = (req, res) => {
    const username = req.body.username
    const email = req.body.email
    const user = {
        username: username,
        email: email,
    }

    const token = jwt.sign(user, token_jwt)
    res.json({accessToken: token})
}


const authenticationToken = (req, res, next) => {
    const headerAuth = req.headers['authorization']
    const token = headerAuth && headerAuth.split(' ')[1]
    if(token == null)return res.sendStatus(401)

    jwt.verify(token, token_jwt, (err, user) => {
        if (err) return res.sendStatus(403)
        req.user = user
        next()
    })
}


module.exports = {
    registerUser,
    authenticationToken
}