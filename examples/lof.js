import anomaly from 'k6/x/anomaly'


export default function () {
    const data = [
        { x: 12, y: 0, timestamp: "My timestamp"},
        { x: 10, y: 0, timestamp: "dsadsa"},
        { x: 11, y: 0, timestamp: "dsadsa"},
        { x: 7, y: 0, timestamp: "dsadsa"},
        { x: 323, y: 0, timestamp: "dsadsa"},
        { x: 150, y: 0, timestamp: "dsadsa"},
        { x: 12, y: 0, timestamp: "dsadsa"}
    ]

    const anomalies = anomaly.lof(data)
    console.log(anomalies)
}

