# xk6-anomaly (experimental)

An xk6 extension for finding anomalies in an automated way from large data sets. The goal of the extension is to be able to detect anomalies easily without the need for third-party tools.

## Build

```shell
xk6 build --with github.com/gpiechnik2/xk6-anomaly@latest
```

## Alghoritms

There are two algorithms in the current version. The first is Local Outlier Factor (lof). The second, on the other hand, is One-Class SVM. Their simplified principle of operation can be found below.

![Alghoritms](https://github.com/gpiechnik2/xk6-anomaly/blob/main/images/alghoritms.png)

Image from [there](https://towardsdatascience.com/5-anomaly-detection-algorithms-every-data-scientist-should-know-b36c3605ea16).

## Example

### Local Outlier Factor

The Local Outlier Factor (LOF) is an algorithm used for outlier detection in data. It is an unsupervised method that evaluates the degree of atypicality of data points relative to their local neighborhood. The LOF algorithm compares the density of a data point to the density of its neighbors, identifying outlier objects that have a lower density compared to their neighbors. As a result, data points with high LOF values are considered potential anomalies or deviations from the norm in the data.

An example:

```javascript
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
```

Important! The moment we want to increase the threshold (that is, decrease the threshold of acceptable data), the second argument will increase (in example: 1).

```javascript
const anomalies = anomaly.lof(data, 6.4)
```

If too much data is considered anomalous, reduce the threshold (for example to `0.04`).

### One-Class SVM

One-Class SVM is an algorithm used for anomaly detection. It trains on a set of normal data points to define a boundary that encloses the normal class. New data points falling outside this boundary are considered anomalies or novelties.

Important! This algorithm requires a large (se suggested quantity is minimum of 50) amount of data to train in order to work properly. An example of the correct code can be found in the `examples` directory.

An example:

```javascript
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
```

Important! The moment we want to increase the threshold (that is, decrease the threshold of acceptable data), the second argument will increase (50 by default).

### Future

I am aware that depending on the application and technologies involved, not all algorithms will be suitable for a project. Therefore, it is necessary to consider multiple different algorithms in order to choose the one that fits best.
