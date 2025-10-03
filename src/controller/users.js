const usersModel = require('../models/users')
const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');
const dotenv = require("dotenv");
const path = require("path");
const { access } = require('fs');
dotenv.config({ path: path.resolve(__dirname, "../.env") });

const token_jwt = process.env.ACCESS_JWT_SECRET

const registerUser = async(req, res) => {
    try {
        const {username, email, password} = req.body
        const hashPass = await bcrypt.hash(password, 10)

        const existingUser = await usersModel.findUserbyEmail(email);
        if (existingUser) {
            return res.status(400).json({
                message: "Email sudah terdaftar, silakan gunakan email lain."
            });
        }
        
        const result = await usersModel.registerUser(username, email, hashPass);

        const token = jwt.sign(
            { id: result[0].insertId, username, email }, // ambil insertId dari result
            token_jwt,
            { expiresIn: "1h" }
        );

        res.json(
            {
                message: "User Telah Berhasil Registrasi",
                accessToken: token
            }
        )
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}


const loginUser = async(req, res) => {
    try {
        const {email, password} = req.body
        const user = await usersModel.findUserbyEmail(email)
        if (!user){
            return res.status(404).json({ message: "User tidak ditemukan" });
        }
    
        const validPass = await bcrypt.compare(password, user.password)
        if(!validPass){
            return res.status(401).json({ message: "Password salah" })
        }
    
        const token = jwt.sign(
            {id: user.id, username: user.username, email: user.email},
            token_jwt,
            {expiresIn: "1h"}
        )
    
        res.json({
            message: "Login Berhasil",
            accessToken: token
        })
    } catch (error) {
        res.status(500).json({ message: "Server Error", error: error.message });
    }
}

module.exports = {
    registerUser,
    loginUser
}