import anomaly from 'k6/x/anomaly'


export default function () {
    const trainingData = [
        // sample data from 0 to 8
		{ x: 3.0 },
        { x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
		{ x: 3.0 },
        { x: 6.0 },
		{ x: 6.0 },
    ]

    const testData = [
        { x: 3.0, timestamp: "2023-07-17T12:02:00"},
        { x: 39.0, timestamp: "2023-07-17T12:02:00"},
        { x: 14.0, timestamp: "2023-07-17T12:02:00"}
    ]

    const anomalies = anomaly.oneClassSvm(trainingData, testData)

    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. X: ${anomaly.x}, Y: ${anomaly.y}, Timestamp: ${anomaly.timestamp}`)
    })

    // INFO[0000] New anomaly detected. X: 3, Y: 0, Timestamp: 2023-07-17T12:02:00  source=console
    // INFO[0000] New anomaly detected. X: 1.5, Y: 0, Timestamp: 2023-07-17T12:02:00  source=console
}
