const db = require("../utils/db");

function getDateNow() {
  const now = new Date();

  return now.getFullYear() + "-" +
    String(now.getMonth() + 1).padStart(2, "0") + "-" +
    String(now.getDate()).padStart(2, "0") + " " +
    String(now.getHours()).padStart(2, "0") + ":" +
    String(now.getMinutes()).padStart(2, "0") + ":" +
    String(now.getSeconds()).padStart(2, "0");
}

function idTrx(){
    const now = new Date();

    return `TRX-${now.getFullYear()}${String(now.getMonth() + 1).padStart(2, "0")}${String(now.getDate()).padStart(2, "0")}${String(now.getHours())}${String(now.getMinutes())}`
}

const getAllTransactions = (user_id) => {
    const sqlQuery = `SELECT * FROM transactions WHERE user_id = ?`
    return db.execute(sqlQuery, [user_id])
}

const getTransactionByID = (idTransaction, id_user) => {
    const sqlQuery = `SELECT * FROM transactions WHERE id = ? AND user_id = ?`
    return db.execute(sqlQuery, [idTransaction, id_user])
}

const addTransaction = (body, user_id) => {
    const sqlQuery = `
        INSERT INTO transactions (id_trx, category, amount, type, note, date, user_id) 
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `;
    return db.execute(sqlQuery, [idTrx(), body.category, body.amount, body.type, body.note, getDateNow(), user_id]);
}

const updateTransaction = (body, idTransaction, user_id) => {
    const sqlQuery = 
    `   UPDATE transactions 
        SET category=?, 
            amount=?, 
            type=?, 
            note=?, 
            date=?
        WHERE id=? AND user_id=?
    `;
    return db.execute(sqlQuery, [body.category, body.amount, body.type, body.note, getDateNow(), idTransaction, user_id]);
}

const deleteTransaction = (idTransaction, user_id) => {
  const sqlQuery = "DELETE FROM transactions WHERE id = ? AND user_id = ?";
  return db.execute(sqlQuery, [idTransaction, user_id]);
};

module.exports = {
    getAllTransactions,
    getTransactionByID,
    addTransaction,
    updateTransaction,
    deleteTransaction
}