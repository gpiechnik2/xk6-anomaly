import anomaly from 'k6/x/anomaly'


export default function () {
    const trainData = [1.0, 2.0, 3.0, 4.0, 5.0]
    const testData = [
        { value: 6.0, timestamp: "2023-07-17T12:02:00"},
        { value: 3, timestamp: "2023-07-17T12:02:00"},
        { value: 1.5, timestamp: "2023-07-17T12:02:00"}
    ]

    const anomalies = anomaly.oneClassSvm(trainData, testData)

    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly.value}, Timestamp: ${anomaly.timestamp}`)
    })

    // INFO[0000] New anomaly detected. X: 6, Y: 0, Timestamp: 2023-07-17T12:02:00  source=console
}
