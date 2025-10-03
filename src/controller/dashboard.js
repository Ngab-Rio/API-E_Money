const dashboardModel = require("../models/dashboard");

const totalIncome = async (req, res) => {
  try {
    const id_user  = req.user.id
    const username = req.user.username
    const [rows] = await dashboardModel.getDashboardSummry(id_user);
    const { total_income, total_expense } = rows[0];
    res.json({
      message: "Get Dashboard Summry",
      data: {
        id_user,
        username,
        total_income: total_income || 0,
        total_expense: total_expense || 0,
        balance: total_income - total_expense
      }
    });
  } catch (error) {
    console.error("Error in Dashboard Summry:", error);
    res.status(500).json({
      message: "Server Error",
      error: error.message, 
    });
  }
};


module.exports = {
    totalIncome
}