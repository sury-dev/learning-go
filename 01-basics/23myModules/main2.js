const https = require("https");

// Reusable function to fetch LTP from Yahoo Finance
function fetchLTP(symbol, callback) {
  const url = `https://query1.finance.yahoo.com/v8/finance/chart/${symbol}?interval=1m&range=1d`;

  const options = {
    headers: {
      "User-Agent":
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/91.0.4472.124 Safari/537.36",
    },
  };

  https
    .get(url, options, (res) => {
      let data = "";

      res.on("data", (chunk) => {
        data += chunk;
      });

      res.on("end", () => {
        try {
          const parsed = JSON.parse(data);
          const result = parsed.chart.result?.[0];
          const price = result?.meta?.regularMarketPrice;
          const time = new Date().toLocaleTimeString();

          if (price) {
            callback(null, { price, time });
          } else {
            callback("Price not found", null);
          }
        } catch (err) {
          callback("Error parsing data", null);
        }
      });
    })
    .on("error", (err) => {
      callback("HTTP error", null);
    });
}

// Wrap LTP logging in a reusable function
function logLTP(symbol) {
  fetchLTP(symbol, (err, data) => {
    if (err) {
      console.error(`[${symbol}] Error:`, err);
    } else {
      console.log(`[${data.time}] ${symbol} - ₹${data.price}`);
    }
  });
}

// === Main Execution Loop ===
const SYMBOL = "RELIANCE.NS"; // Change to your desired symbol
const INTERVAL_MS = 500; // 15 seconds

console.log(`📈 Tracking LTP for ${SYMBOL} every 15 seconds...\n`);
logLTP(SYMBOL); // Initial log
setInterval(() => logLTP(SYMBOL), INTERVAL_MS);
