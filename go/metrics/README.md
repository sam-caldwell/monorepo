metrics
=======

## Objective:
This project creates a generic metrics collector/abstraction tool.  This tool will (a) abstract data collectors from
any third-party from our actual instrumentation to ensure we can pivot quickly, and (b) this tool will allow us to
pursue performance and accuracy (hint: we should know when and where aggregations are occurring rather than making
assumptions).

## Usage (simple metrics emitter)
### Assumptions
* This tool will emit metrics to a log file, stdout or stderr as structured metrics to avoid overhead of network
  connections and other resource-consuming alternatives.
* This tool assumes the operating system or parent process executing the instrumented code will consume the emitted
  metrics data and perform any storage/analysis.

## Architecture
The following diagram defines the end state of this project when completed.
```text
                Emitter ---------> Collector
                    ^                  |
                    |                  V
                MetricSet             TSDB
                 /     \                |
                /       \               V
               /         \           Analyzer
           Scalar      State
           Metric      Metric
```

### Emitter
* The `Emitter` uses a timer to emit metrics from a `MetricSet` periodically and write the same as a structured record.
* The `Emitter` will ensure that all `MetricSet` emissions are properly structured, written to the intended target.

### Collector
* The `Collector` will consume information from a file or stdout/stderr and process it into a time-series database
  (TSDB)

### MetricSet
* A `MetricSet` is a collection of `Scalar`, `BigScalar` or `State` Metric objects.  The MetricSet is a map of metric
  names and corresponding metric objects which can be scanned by the Emitter to serialize the metric records and
  emit them to the target.
* A `MetricSet` is a logically related set of metrics.  A program could have one or many `MetricSet` objects.

### Scalar Metrics
* A Scalar metric is some numeric value which can be represented on a number line.  This metric contains the current
  value, a bucket of past values spanning some window (number of values) and facilities to perform standard aggregations

#### Scalar Aggregations
* The scalar supports the following standard aggregations:
  * MovMin - the moving minimum value of the metric window
  * MovMax - the moving maximum value of the metric window
  * MovAvg - moving average value (an average of the current window)
  * Min - the overall minimum value from the first collected metric
  * Max - the overall maximum value from the first collected metric
  * Avg - the overall average value from the first collected metric
  * Count - the overall count of values collected since first metric

### State Metric
* A State metric is an object which collects some blob state (text, bytes, etc.) representing a current value (last
  recorded state) and a window of recent past states.

#### State Aggregations
* The State Metric supports capture of the last captured state.
