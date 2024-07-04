# Logs beautifier fo Go apps

Make logs looks readable

![image](https://github.com/fruitgum/logger/assets/31319804/e5426ed2-efa1-470b-92b8-ac98c0a7b2ab)



## Usage:

```
package main

import "github.com/fruitgum/logger"

func main() {

    logger.System("This is System message") //white
    logger.Debug("This is Debug message") //magenta
    logger.Info("This is Info message") //Cyan
    logger.Success("This is Success message") //Green
    logger.Warn("This is Warn message") //Yellow
    logger.Error("This is Error message") //Red
    

}
```

#### You can also switch log levels
```
package main

import "github.com/fruitgum/logger"

func main() {

    logLevel := flag.String("loglevel", "info", logger.HelpUsage())
    setLogLevel := logger.SetLogLevel(*logLevel)
    logger.System("Log level set to %s", setLogLevel)
    
}

```
#### or
```
package main

import "github.com/fruitgum/logger"

func main() {

    logLevel := debug
    setLogLevel := logger.SetLogLevel(logLevel)
    logger.System("Log level set to %s", setLogLevel)
    
}

```
### Available levels you can set:
* `debug` - will suppress nothing
* `info` - will suppress debug messages
* `warn` - will suppress info and debug messages
* `fatal` or `error` - will suppress info, warn and debug messages
* `minimal` - will suppress all messages except System and Success 
* `none` - will suppress everything 

