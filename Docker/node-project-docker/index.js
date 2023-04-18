const express = require("express");
const app = express();

app.get("/", (req, res) => {
    res.json({
        msg: "This is from root ..."
    })
});
app.get("/greet", (req, res) => {
    res.json({
        msg: "This is from greet. Hello ...."
    })
});
app.get("/about", (req, res) => {
    res.json({
        msg: "This is from about"
    })
});
app.get("/premium", (req, res) => {
    res.json({
        msg: "Great Day! It's time to grab premium ...."
    })
});
app.listen(8000, () => {
    console.log("App is running at 8000")
})