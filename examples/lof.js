import anomaly from 'k6/x/anomaly'


export default function () {
    const testData = [
        { value: 12, timestamp: "2023-07-17T12:02:00"},
        { value: 10, timestamp: "2023-07-17T12:02:00"},
        { value: 11, timestamp: "2023-07-17T12:02:00"},
        { value: 7, timestamp: "2023-07-17T12:02:00"},
        { value: 12, timestamp: "2023-07-17T12:02:00"},

        // anomalies
        { value: 323, timestamp: "2023-07-17T12:02:00"},
        { value: 150, timestamp: "2023-07-17T12:02:00"}
    ]

    const anomalies = anomaly.lof(testData, 1)
    
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly.value}, Timestamp: ${anomaly.timestamp}, lofScore: ${anomaly.lof_score}`)
    })

    // INFO[0000] New anomaly detected. Value: 6, Timestamp: 2023-07-17T12:02:00  source=console
    // INFO[0000] New anomaly detected. Value: 1.5, Timestamp: 2023-07-17T12:02:00  source=console
}

