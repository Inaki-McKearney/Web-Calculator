// index.js

const express = require("express");
const path = require("path");
const request = require("request");
const { createLogger, transports ,format} = require('winston');
const lineByLine = require('n-readlines');


const app = express();
const port = process.env.PORT || "80";
const proxyURL = 'http://proxy/';
const operators = ['add','modulo','multiply','quotient','power', 'subtract'];

app.set("views", path.join(__dirname, "views"));
app.set("view engine", "pug");
app.use(express.static(path.join(__dirname, "public")));

var allResponses = {};

// Create instance of Winston logger
const logger = createLogger({
    level: 'info',
    format: format.combine(
        format.json(),
        format.timestamp()
    ),
    transports: [
        new transports.File({ filename: 'responseTimes.log', timestamp: true })
    ]
});

function loadPreviousResults() {
    operators.forEach(function(operator) {
        allResponses[operator] =  {
            numberResponses: 0,
            average: 0,
            max: 0,
            min: 100,
            totalTime: 0,
            status: 0,
            correct: 0
        };
    })
    
    var liner = new lineByLine('./responseTimes.log');

    // var line;

    while (line = liner.next()) {
        updateTimeStats(JSON.parse(line));
    }
        
    startRequestLoop();
    return allResponses; 
}

function updateTimeStats(stats) {
    op = stats.operator;
    allResponses[op].numberResponses = allResponses[op].numberResponses + 1;
    allResponses[op].totalTime = allResponses[op].totalTime + stats.responseTime;
    allResponses[op].average = allResponses[op].totalTime / allResponses[op].numberResponses;
    allResponses[op].max = Math.max(allResponses[op].max,stats.responseTime);
    allResponses[op].min = Math.min(allResponses[op].min,stats.responseTime);
    allResponses[op].status = stats.status;
    allResponses[op].correct = stats.correct ? allResponses[op].correct + 1 : allResponses[op].correct;           
}

app.get("/", (req, res) => {
    res.render("index", {title: "Home", stats: allResponses});
});


app.listen(port, () => {
    console.log(`Listening to requests on http://localhost:${port}`);
    loadPreviousResults();
});

//Generate random numbers for API query parameters
function genRand(min, max, excludeZero) {
    var num = Math.floor(Math.random() * (max - min + 1)) + min;
    return (excludeZero && num === 0) ? genRand(min, max) : num;
}

//Execute getResponse every x milliseconds
function startRequestLoop() {
    var responseLoop = setInterval(getResponse,1000);
}

function generateRequest() {
    var operator = operators[Math.floor(Math.random()*operators.length)];
    var url = proxyURL+operator;
    var x = genRand(-18,18,false) // -18 <= y <= 15
    var y = genRand(-15,15,true); // -15 <= y < 0 & 0 < y <= 15
    answer = 0

    switch (operator) {
        case 'add':
            answer = x + y;    
            break;
        case 'modulo':
            answer = ((x % y) + y) % y;    
            break;
        case 'multiply':
            answer = x * y;    
            break;
        case 'quotient':
            answer = ~~(x / y);    
            break;
        case 'power':
            answer = Math.round(Math.pow(x, y));
            break;
        case 'subtract':
            answer = x - y;    
            break;
    }

    return {
        url: `${url}?x=${x}&y=${y}`,
        operator: operator,
        answer: answer,
    }
}

function getResponse(url) {
    var req = generateRequest();
    var url = req.url;
    var operator = req.operator;
    var expectedAnswer = req.answer;

    request.get({
        url : url,
        time : true,
        json: true
    },function(err, response){
        if (err != null) {
            return console.log(`error: ${err}`);
        }
        if (response.statusCode == 200) {
            answer = response.body.answer;
        } else {
            answer = Infinity;
        }
        var result = {
            operator: operator,
            correct: Math.abs(answer - expectedAnswer) <= 0.00001,
            status: response.statusCode,
            responseTime: response.elapsedTime,
            timestamp: new Date().toUTCString()
        }
        logger.log('info',result);
        updateTimeStats(result);
    });
}