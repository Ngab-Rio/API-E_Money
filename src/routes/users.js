const express = require("express")
const router = express.Router()

const userController = require('../controller/users')
const authMiddleware = require('../middleware/verifyUser')

router.post("/register", userController.registerUser);


router.post("/login", userController.loginUser)

router.get("/profile", authMiddleware.authenticationToken, (req, res) => {
    res.json({
        message: `Welcome ${req.user.username}`,
        // user: req.user
    })
})

module.exports = router