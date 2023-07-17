import anomaly from 'k6/x/anomaly'


export default function () {
    const trainingData = [
		{x: 1.0, y: 2.0},
		{x: 2.0, y: 3.0},
		{x: 3.0, y: 4.0},
		{x: 4.0, y: 5.0},
		{x: 5.0, y: 6.0}
    ]

    const testData = [
        { x: 6.0, y: 7, timestamp: "2023-07-17T12:02:00"},
        { x: 3, y: 4, timestamp: "2023-07-17T12:02:00"},
        { x: 1.5, y: 2.5, timestamp: "2023-07-17T12:02:00"},
    ]

    const anomalies = anomaly.oneClassSvm(trainingData, testData, 1)

    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. X: ${anomaly.x}, Y: ${anomaly.Y}, Timestamp: ${anomaly.timestamp}`)
    });
}
