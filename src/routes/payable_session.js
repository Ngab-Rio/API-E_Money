const express = require("express")
const router = express.Router()

const payableSessionController = require("../controller/payable_session")
const authMiddleware = require("../middleware/verifyUser")

// SHOW ALL PAYABLE SESSION
router.get("/payable/:payable_id/session", authMiddleware.authenticationToken, payableSessionController.showDataPayable)

// CREATE NEW PAYABLE SESSION
router.post("/payable/:payable_id/session", authMiddleware.authenticationToken, payableSessionController.createPayableSession)



module.exports = router