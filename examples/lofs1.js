import anomaly from 'k6/x/anomaly'


export default function () {
    const data = [
        { x: 12, timestamp: "2023-07-17T12:02:00"},
        { x: 10, timestamp: "2023-07-17T12:02:00"},
        { x: 11, timestamp: "2023-07-17T12:02:00"},
        { x: 7, timestamp: "2023-07-17T12:02:00"},
        { x: 323, timestamp: "2023-07-17T12:02:00"},
        { x: 150, timestamp: "2023-07-17T12:02:00"},
        { x: 12, timestamp: "2023-07-17T12:02:00"}
    ]

    const anomalies = anomaly.lof(data, 1)
    
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. X: ${anomaly.x}, Y: ${anomaly.y}, Timestamp: ${anomaly.timestamp}, lofScore: ${anomaly.lof_score}`)
    })

    // INFO[0000] New anomaly detected. X: 323, Y: 0, Timestamp: 2023-07-17T12:02:00, lofScore: 0.004032258064516129  source=console
    // INFO[0000] New anomaly detected. X: 150, Y: 0, Timestamp: 2023-07-17T12:02:00, lofScore: 0.008036739380022962  source=console
}

