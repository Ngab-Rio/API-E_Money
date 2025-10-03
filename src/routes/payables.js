const express = require("express")
const router = express.Router()

const payablesController = require("../controller/payables")
const authMiddleware = require("../middleware/verifyUser")

// GET ALL PAYABLES
router.get("/payables", authMiddleware.authenticationToken, payablesController.showAllPayables)

// ADD PAYABLE
router.post("/payable/create", authMiddleware.authenticationToken, payablesController.createNamePayable)

// DELETE PAYABLE
router.delete("/payable/delete/:id_payable", authMiddleware.authenticationToken, payablesController.deletePayable)

module.exports = router