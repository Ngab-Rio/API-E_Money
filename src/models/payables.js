const db = require("../utils/db")

function getDateNow() {
  const now = new Date();

  return now.getFullYear() + "-" +
    String(now.getMonth() + 1).padStart(2, "0") + "-" +
    String(now.getDate()).padStart(2, "0") + " " +
    String(now.getHours()).padStart(2, "0") + ":" +
    String(now.getMinutes()).padStart(2, "0") + ":" +
    String(now.getSeconds()).padStart(2, "0");
}


const showAllPayables = (user_id) =>{
    sqlQuery = `
        SELECT * FROM payables WHERE user_id = ?
    `
    return db.execute(sqlQuery, [user_id])
}

const createNamePayable = (user_id, body) => {
    sqlQuery = `
        INSERT INTO payables (user_id, name, created_at)
        VALUES (?, ?, ?)
    `
    return db.execute(sqlQuery, [user_id, body.name, getDateNow()])
}

const deletePayable = (id_payable, user_id) => {
    sqlQuery = `
        DELETE FROM payables WHERE id = ? AND user_id = ?
    `
    return db.execute(sqlQuery, [id_payable, user_id])
}

module.exports = {
    showAllPayables,
    createNamePayable,
    deletePayable
}