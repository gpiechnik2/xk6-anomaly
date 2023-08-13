import anomaly from 'k6/x/anomaly'

export default function () {
    let anomalies
    const testData = [ 
        12, 10, 11, 7, 12,
        323, 150 // anomalies
    ]

    anomalies = anomaly.lof(testData)
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly.value} lofScore: ${anomaly.lof_score}`)

        // INFO[0000] New anomaly detected. Value: 323 lofScore: 0.004032258064516129  source=console
        // INFO[0000] New anomaly detected. Value: 150 lofScore: 0.008036739380022962  source=console
    })

    // we increased threshold by 1 (by default it is 1.0)
    anomalies = anomaly.lof(testData, 2)
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly.value} lofScore: ${anomaly.lof_score}`)

        // INFO[0000] New anomaly detected. Value: 323 lofScore: 0.004032258064516129  source=console
    })
}
