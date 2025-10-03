const express = require("express")
const router = express.Router()

const dashboardController = require("../controller/dashboard")
const authMiddleware = require("../middleware/verifyUser")

// SHOW TOTAL INCOME
router.get("/dashboard", authMiddleware.authenticationToken, dashboardController.totalIncome)

module.exports = router