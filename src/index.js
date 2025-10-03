const express = require("express");
const app = express();

const cors = require("cors");

app.use(cors()); 

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const dotenv = require("dotenv")
const path = require("path")
dotenv.config({ path: path.resolve(__dirname, ".env") });

const host = process.env.HOST
const port = process.env.PORT


const dashboardRoute = require("./routes/dashboard")
app.use("/", dashboardRoute)

const transactionRoute = require("./routes/transactions")
app.use("/", transactionRoute)

const userRoute = require("./routes/users")
app.use("/auth", userRoute)

const payablesRoute = require("./routes/payables")
app.use("/", payablesRoute)

const payableSessionRoute = require("./routes/payable_session")
app.use("/", payableSessionRoute)

app.listen(port, ()=>{
    console.log(`[+] Server Running in http://${host}:${port}`)
})