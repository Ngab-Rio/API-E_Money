const db = require("../utils/db");

const getDashboardSummry = (id_user) => {
    const sqlQuery = `
        SELECT
          SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) AS total_income,
          SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) AS total_expense
        FROM transactions WHERE user_id = ?
    `;
    return db.execute(sqlQuery, [id_user]);
}

module.exports = {
    getDashboardSummry,
}