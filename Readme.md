**Disclaimer**

Work in progress. Not ready for use.

---

# Logging

A simple wrapper around the [logrus library](https://github.com/sirupsen/logrus). This library defines a Logger interface which simplifies the logging framework.

## Motivation

Logrus provides lots of features, which can be useful while also a bit confusing or overwhelming. This library allows the use of all those features, while providing a limited interface to keep the logging portion as simple as possible.

The interfaces supports only 3 levels of logging:

* Debug
* Info
* Error

This was inspired by [Dave Cheney's thoughts on logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging). Although Dave Cheney recommends the use of only 2 levels, Debug and Info, this library allows the use of Error level as well to keep things a little flexible. Sometimes we do need to log Error messages.

The logging levels also have their corresponding `f` versions, i.e., Debugf, Infof, Errorf.

In addition to this, there are 2 more methods available, `WithField` and `WithFields`. These work exactly like they do in logrus. 

In fact, all methods under the Logger interface pass the calls directly to the underlying Logrus methods transparently. They don't modify the log entries at all. The only modification that it does carry out is in the logging level. Since it supports only 3 logging levels, the logging level is reset to the closest one if a different one was set. This means that a log level higher than Debug is reset to Debug and a level lower than Error is reset to Error.

## Usage

```
func main(){

    // create a new logrus logger
    logrusLogger := logrus.New()

    // setup logrus
    // - add all your favourite logrus hooks and other setup
    ...

    // create a new logger
    logger := logging.NewLogger(logrus.Logger)

    // use it as you would use logrus
    logger.Info("Hello")
    logger.Infof("%d Worlds", 3)

    logger.WithField("animal", "walrus").Info("found animal")
}
```

