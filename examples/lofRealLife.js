import http from 'k6/http'
import { Counter } from 'k6/metrics'
import anomaly from 'k6/x/anomaly'

const anomaliesMetric = new Counter('anomalies')
const data = []

function pushAndFindAnomaly(responseTime, name) {
    data.push(responseTime)

    // need minimum of 40 training data
    if (data.length >= 40) {
        const anomalies = anomaly.lof(data, 1)
        const lastAnomaly = anomalies[anomalies.length - 1]

        // current response time is qualified as an anomaly
        if (lastAnomaly.value == responseTime) {
            anomaliesMetric.add(lastAnomaly.value, {
                name: name
            })
        }
    }
}

export default function () {
    let response
    let name
    let responseTime

    group('First API call', function () {
        name = "GET /"
        response = http.get('http://test.k6.io')
        responseTime = response.timings.duration
        pushAndFindAnomaly(responseTime, name)
    })

    group('Second API call', function () {
        name = "GET / 2"
        response = http.get('http://test.k6.io')
        responseTime = response.timings.duration
        pushAndFindAnomaly(responseTime, name)
    })

    group('Third API call', function () {
        name = "GET / 3"
        response = http.get('http://test.k6.io')
        responseTime = response.timings.duration
        pushAndFindAnomaly(responseTime, name)
    })

    const testData = [ 
        12, 10, 11, 7, 12, 
        323, 150 // anomalies
    ]
        

    const anomalies = anomaly.lof(testData, 1)
    
    anomalies.forEach(anomaly => {
        console.log(`New anomaly detected. Value: ${anomaly.value}, Timestamp: ${anomaly.timestamp}, lofScore: ${anomaly.lof_score}`)
    })

    // INFO[0000] New anomaly detected. Value: 6, Timestamp: 2023-07-17T12:02:00  source=console
    // INFO[0000] New anomaly detected. Value: 1.5, Timestamp: 2023-07-17T12:02:00  source=console
}
