const db = require("../utils/db");

const registerUser = (username, email, hashPass) =>{
    const sqlQuery = `
        INSERT INTO users (username, email, password)
        VALUES (?, ?, ?)
    `
    return db.execute(sqlQuery, [username, email, hashPass])
}


const findUserbyEmail = async(email) => {
    const [rows] = await db.execute("SELECT * FROM users WHERE email = ?", [email]);
    return rows[0];
}


module.exports = {
    registerUser,
    findUserbyEmail
}