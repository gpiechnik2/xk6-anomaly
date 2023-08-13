import anomaly from 'k6/x/anomaly'

export default function () {
    let anomalies
    const trainData = [1.0, 2.0, 3.0, 4.0, 5.0]
    const testData = [
        3.0, 4,1,
        6.0, 1.5 // anomalies
    ]

    anomalies = anomaly.oneClassSvm(trainData, testData)
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly}`)
        // INFO[0000] New anomaly detected. Value: 6                source=console
        // INFO[0000] New anomaly detected. Value: 1.5              source=console
    })

    // we decreased threshold by 49 (by default it is 50.0)
    anomalies = anomaly.oneClassSvm(trainData, testData, 1)
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly}`)
        // INFO[0000] New anomaly detected. Value: 6                source=console
    })
}
