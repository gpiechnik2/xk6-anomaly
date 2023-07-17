# xk6-anomaly

An xk6 extension for finding anomalies in an automated way from large data sets. The goal of the extension is to be able to detect anomalies easily without the need for third-party tools.

## Build

```shell
xk6 build --with github.com/gpiechnik2/xk6-anomaly@latest
```

## Alghoritms

There are two algorithms in the current version. The first is Local Outlier Factor (lof). The second, on the other hand, is One-Class SVM. Their simplified principle of operation can be found below.

TODO: screen

## Example

### Local Outlier Factor

The Local Outlier Factor (LOF) is an algorithm used for outlier detection in data. It is an unsupervised method that evaluates the degree of atypicality of data points relative to their local neighborhood. The LOF algorithm compares the density of a data point to the density of its neighbors, identifying outlier objects that have a lower density compared to their neighbors. As a result, data points with high LOF values are considered potential anomalies or deviations from the norm in the data. The LOF algorithm is popular in fields such as fraud detection, network monitoring, image analysis, and many others, where detecting unusual observations in data is crucial.

```javascript
TODO
```

### One-Class SVM

One-Class SVM is an algorithm used for anomaly detection. It trains on a set of normal data points to define a boundary that encloses the normal class. New data points falling outside this boundary are considered anomalies or novelties. It is commonly used in applications like fraud detection and intrusion detection.

```javascript
TODO
```

### TODO

I am aware that depending on the application and technologies involved, not all algorithms will be suitable for a project. Therefore, it is necessary to consider multiple different algorithms in order to choose the one that fits best.

In my plans, I aim to cover the following algorithms.

- [x] Local Outlier Factor
- [x] One-Class SVM
- [ ] Isolation Forest
- [ ] Robust Covariance

### License

TODO
