const payablesModel = require("../models/payables")

const showAllPayables = async(req, res) => {
    const user_id = req.user.id
    const [body] = await payablesModel.showAllPayables(user_id)
    try {
        res.json({
            message: "Get All Payables",
            data: body
        })
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const createNamePayable = async(req, res) =>{
    const user_id = req.user.id
    const { body } = req
    try {
        await payablesModel.createNamePayable(user_id, body)
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

const deletePayable = async(req, res) => {
    const user_id = req.user.id
    const { id_payable } = req.params
    const [data] = await payablesModel.deletePayable(id_payable, user_id)
    try {
        if (data.affectedRows === 0){
            res.status(404).json({
                message: "ID Not Found",
            })
        }
        res.json({
            message: "Delete Transaction Was Successed",
            deletedId: id_payable
        })
    }
    catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

module.exports = {
    showAllPayables,
    createNamePayable,
    deletePayable
}