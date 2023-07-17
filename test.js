import anomaly from 'k6/x/anomaly';
import http from "k6/http";


export default function () {
    // const response = http.get("https://test-api.k6.io/public/crocodiles/");

    const data = [
        { x: 12, y: 0, timestamp: "dsadsa"},
        { x: 10, y: 0, timestamp: "dsadsa"},
        { x: 11, y: 0, timestamp: "dsadsa"},
        { x: 7, y: 0, timestamp: "dsadsa"},
        { x: 323, y: 0, timestamp: "dsadsa"},
        { x: 150, y: 0, timestamp: "dsadsa"},
        { x: 12, y: 0, timestamp: "dsadsa"},
    ]

    anomaly.lof(data)
    // console.log(anomalies)
}

