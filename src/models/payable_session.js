const db = require("../utils/db")

const showDataPayable = (payable_id) => {
    sqlQuery = `
        SELECT * FROM payable_sessions WHERE payable_id = ? ORDER BY date DESC
    `
    return db.execute(sqlQuery, [payable_id])
}

const createPayableSession = (payable_id, body) => {
    sqlQuery = `
        INSERT INTO payable_sessions (payable_id, amount, note, status, date)
        VALUES (?, ?, ?, ?, NOW())
    `
    return db.execute(sqlQuery, [payable_id, body.amount, body.note, body.status])
}

module.exports = {
    showDataPayable,
    createPayableSession
}