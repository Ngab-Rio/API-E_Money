const transactionsModel = require("../models/transactions")

const getAllTransactions = async(req, res) => {
    try {
        const user_id = req.user.id
        const [rows] = await transactionsModel.getAllTransactions(user_id)
        res.json({
            message: "Get All transactions Was Successed",
            data: rows
        })
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const getTransactionByID = async(req, res) => {
    try {
        const {idTransaction} = req.params
        const user_id = req.user.id
        const [rows] = await transactionsModel.getTransactionByID(idTransaction, user_id)
        
        if (rows.length === 0){
            return res.status(404).json({
                message: "ID Was Not Found",
                error: console.error()
            })
        }else{
            res.json({
                message: "Get All transactions Was Successed",
                data: rows
            })
        }
        
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const addTransaction = async(req, res) => {
    const { body } = req
    const user_id = req.user.id
    try {
        await transactionsModel.addTransaction(body, user_id)
        res.json({
            message: "Add transaction Was Successed",
            data: body
        })
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const updateTransaction = async(req, res) => {
    const { idTransaction } = req.params
    const { body } = req
    const user_id = req.user.id

    try {
        await transactionsModel.updateTransaction(body, idTransaction, user_id)
        res.json({
            message: "Update transaction Was Successed",
            data: {
                id: idTransaction,
                ...body
            }
        }) 
    } catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

const deleteTransaction = async(req, res) => {
    try {
        const { idTransaction } = req.params
        const user_id = req.user.id
        const [result] = await transactionsModel.deleteTransaction(idTransaction, user_id)

        if (result.affectedRows === 0){
            res.status(404).json({
                message: "ID Not Found",
            })
        }
        res.json({
            message: "Delete Transaction Was Successed",
            deletedId: idTransaction
        }) 
    }catch (error) {
        res.status(500).json({
            message: "Server Error",
            error: error.message
        })
    }
}

module.exports = {
    getAllTransactions,
    addTransaction,
    updateTransaction,
    deleteTransaction,
    getTransactionByID
}