const express = require("express")
const router = express.Router()

const transactionController = require("../controller/transactions")
const authMiddleware = require("../middleware/verifyUser")

// GET ALL TRANSACTIONS
router.get("/list", authMiddleware.authenticationToken, transactionController.getAllTransactions)

// GET TRANSACTION BY ID
router.get("/list/:idTransaction", authMiddleware.authenticationToken, transactionController.getTransactionByID)

// ADD TRANSACTION
router.post("/add", authMiddleware.authenticationToken, transactionController.addTransaction)

// UPDATE TRANSACTION
router.put("/update/:idTransaction", authMiddleware.authenticationToken, transactionController.updateTransaction)

// DELETE TRANSACTION
router.delete("/delete/:idTransaction",authMiddleware.authenticationToken, transactionController.deleteTransaction)

module.exports = router