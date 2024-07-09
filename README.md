# Logs beautifier fo Go apps

Make logs looks readable

![image](https://github.com/fruitgum/go-logger/assets/31319804/cbe8ff0e-8380-44bd-af52-24a181886a98)



## Usage:

```
package main

import "github.com/fruitgum/go-logger"

func main() {

    	logs := goLogger.New()

        logs.Success("Success")
        logs.System("System")
        logs.Debug("Debug")
        logs.Info("Info")
        logs.Warn("Warn")
        logs.Error("Error")
        logs.Fatal("Fatal")
    

}

```

### Setting log level

> [!INFO]
> Default log level - info
> If you will not set log level or pass incorrect log level - default log level will be set

```
package main

import "github.com/fruitgum/go-logger"

func main(){

    logs := goLogger.New()
    logs.SetLogLevel("debug")
    
}
```

### Redirecting output to file

> [!IMPORTANT]
> ToFile accepts path as argument
> 
> If path is not exist - ToFile will try to create it
> 
> If path set as "" - ToFile will create directory "logs" next to application
> 
> If ToFile won't be able to create path and/or log file - logs will be redirected to stdout as default


> [!NOTE]
> For example you want to store your logs in /tmp/myApp
> 
> logs.ToFile("/tmp/myApp")
> 
> ToFile will create /tmp/myApp/*current_year*/*current_month*/*current_date*.log and redirect the logs there


```
package main

import "github.com/fruitgum/go-logger"

func main(){

    logs := goLogger.New()
    logs.ToFile("logs")
    
}
```

### Available levels you can set:
* `debug` - will suppress nothing
* `info` - will suppress debug messages
* `warn` - will suppress info and debug messages
* `error` - will suppress info, warn and debug messages
* `fatal` - will suppress all messages except System, Success and Fatal
* `none` - will suppress everything

> [!IMPORTANT]
> `logger.Fatal` level will terminate process even if its output is suppressed  