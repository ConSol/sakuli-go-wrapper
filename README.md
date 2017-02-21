## Platform-independent Sakuli-Starter in Go
### Compile by your self
#### Requirements:
- Go 1.5

#### Installation
```bash
go get github.com/ConSol/sakuli-go-wrapper
```

- Own OS
```bash
go build github.com/ConSol/sakuli-go-wrapper
```

- Cross-Platform (Example is for compiling on Windows for Linux)
```bash
set GOOS=linux
set GOARCH=386
go build github.com/ConSol/sakuli-go-wrapper
```

- Build windows and linux binary
``bash
make
```
or 
```
mvn install
```

### Usage

```
Generic Sakuli test starter.
2016 - The Sakuli team / Philip Griesbacher.
http://www.sakuli.org
https://github.com/ConSol/sakuli

Usage of sakuli-go-wrapper:
  -D value
    	JVM option to set a property on runtime, overwrites the 'sakuli.properties'
  -browser string
    	(optional) browser for the test execution (default: Firefox)
  -encrypt string
    	encrypt a secret
  -interface string
    	(optional) network interface used for encryption
  -javahome string
    	Java bin dir (overwrites PATH)
  -loop int
    	loop this suite, wait n seconds between executions, 0 means no loops (default: 0)
  -postHook value
    	A programm which will be executed after sakuli (Can be added multiple times)
  -preHook value
    	A programm which will be executed before sakuli (Can be added multiple times)
  -run string
    	run a sakuli test suite
  -sahi_home string
    	(optional) Sahi installation folder
  -sakuli_home string
    	(optional) SAKULI_HOME folder, default: environment variable 'SAKULI_HOME' (default "C:\\XXX")

```
