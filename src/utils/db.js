const mysql = require("mysql2");
const dotenv = require("dotenv");
const path = require("path");
const { create } = require("domain");
dotenv.config({ path: path.resolve(__dirname, "../.env") });

const db = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    database: process.env.DB_NAME,
    password: process.env.DB_PASS,
    port: process.env.DB_PORT,
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0
});

module.exports = db.promise();