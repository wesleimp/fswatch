const express = require("express")

const app = express()

app.get("/yaya", (_, res) => {
    console.log("receive a request")
    res.json({status: "ok"})
})

app.listen(3000, () => console.log("server running"))
