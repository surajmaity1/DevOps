const mongoose = require("mongoose");
const express = require('express')

const app = express()

const database = "mongodb://mymongo:27017/testup";

mongoose
    .connect(database)
    .then(() => {
        console.log("Database Connected");
    })
    .catch(() => {
        console.log("Database connection failed ...");
    });

app.get("/", (req, res) => {
    res.json({
        msg: "Hi Everyone ..."
    });
});

app.listen(8000, () => {
    console.log("App is listening -> 8000 ...");
});