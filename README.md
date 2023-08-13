# xk6-anomaly

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

Simple example:

```javascript
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
```

*Important!* At the moment when we want to "sensitize" the algorithm to the anomalies being checked, we should decrease the value of the 2nd argument in the called lof function (by default, this value is set to 1). When we want to have more acceptable data, we should increase the third value of the function (in the example, increasing the value from 1.0 to 2.0 resulted in ignoring one of the anomalies).

### One-Class SVM

One-Class SVM is an algorithm used for anomaly detection. It trains on a set of normal data points to define a boundary that encloses the normal class. New data points falling outside this boundary are considered anomalies or novelties.

Important! This algorithm requires a large (se suggested quantity is minimum of 50) amount of data to train in order to work properly. An example of the correct code can be found in the `examples` directory.

Simple example:

```javascript
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
```

Important! When we want to "sensitize" the algorithm to the anomalies being checked, we should increase the value of the 3rd argument in the called lof function (by default, this value is set to 50). When we want to have more acceptable data, we should decrease the third value of the function (in the example, we reduced the value from 50 to 1, as a result, the value 1.5 was not considered an anomaly).

## Visualization

A real-world usage example can be found at the path `examples/usageExample.js`. In short: upon detecting anomalies based on the collected data, they are sent to the influxDB database. They are tagged with the name of the endpoint where they were found and the name of the algorithm. The used dashboard is located in the `dashboards` directory.

![k6 anomaly detector grafana](https://github.com/gpiechnik2/xk6-anomaly/blob/main/images/k6-anomaly-grafana.jpg)


### Future

I am aware that depending on the application and technologies involved, not all algorithms will be suitable for a project. Therefore, it is necessary to consider multiple different algorithms in order to choose the one that fits best.
