const payableSessionModel = require("../models/payable_session")

const createPayableSession = async(req, res) => {
    const payable_id = req.params.payable_id
    const { body } = req
    await payableSessionModel.createPayableSession(payable_id, body)
    try {
        res.json({
            message: "Create Payable Success",
            data: body
        })
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const showDataPayable = async(req, res) => {
    const payable_id = req.params.payable_id
    const [rows] = await payableSessionModel.showDataPayable(payable_id)
    try {
        res.json({
            message: "Show Data",
            data: rows
        })
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

module.exports = {
    createPayableSession,
    showDataPayable
}