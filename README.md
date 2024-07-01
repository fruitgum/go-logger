### Logs beautifier fo Go apps

Makes logs looks readable

![image](https://github.com/fruitgum/logger/assets/31319804/4515612f-55a7-4a5d-afbc-6f8fcc8b8a52)

Usage:

```
package main

import "github.com/fruitgum/logger"

func main() {

	logger.Debug("This is Debug message")
	logger.Info("This is Info message")
	logger.Success("This is Success message")
	logger.Warn("This is Warn message")
	logger.Error("This is Error message") // This func will stop further running of code 

}
```
Thanks Fatih for their [color lib](https://github.com/fatih/color)
