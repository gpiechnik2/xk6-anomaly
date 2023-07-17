import anomaly from 'k6/x/anomaly'


export default function () {
    const data = [
        { x: 12, y: 10, timestamp: "2023-07-17T12:02:00"},
        { x: 10, y: 14, timestamp: "2023-07-17T12:02:00"},
        { x: 11, y: 9, timestamp: "2023-07-17T12:02:00"},
        { x: 7, y: 10, timestamp: "2023-07-17T12:02:00"},
        { x: 323, y: 14, timestamp: "2023-07-17T12:02:00"},
        { x: 150, y: 9, timestamp: "2023-07-17T12:02:00"},
        { x: 12, y: 14, timestamp: "2023-07-17T12:02:00"}
    ]

    const anomalies = anomaly.lof(data, 1)

    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. X: ${anomaly.x}, Y: ${anomaly.y}, Timestamp: ${anomaly.timestamp}, lofScore: ${anomaly.lof_score}`)
    })

    // INFO[0000] New anomaly detected. X: 323, Y: 14, Timestamp: 2023-07-17T12:02:00, lofScore: 0.004031878708778318  source=console
    // INFO[0000] New anomaly detected. X: 150, Y: 9, Timestamp: 2023-07-17T12:02:00, lofScore: 0.00803434875892064  source=console
}

