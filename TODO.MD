* -version
    ```
    --- Environment variables ---
    SAKULI_HOME:                  /home/tschneck/sakuli/sakuli-v0.9.3-SNAPSHOT
    MOZ_DISABLE_OOP_PLUGINS:      1
    MOZ_DISABLE_AUTO_SAFE_MODE:   1
    MOZ_DISABLE_SAFE_MODE_KEY:    1
    ```
    ->
    ```
    C:\Program Files\sakuli\sakuli-v0.9.3-SNAPSHOT\bin>set SAKULI_HOME=C:\Program Files\sakuli\sakuli-v0.9.3-SNAPSHOT
  C:\Program Files\sakuli\sakuli-v0.9.3-SNAPSHOT\bin>sakuli-go-wrapper.exe -version
  Sakuli version:     0.9.3-SNAPSHOT
  Built time:         19.02.2016 @ 16:47:57 CET
  GIT commit ID:      f857134239080b85c63c43f0c686259eaf2f5dc0

  --- Libraries ---
  Sahi OS version:    5.0
  SikuliX version:    1.1.997

  --- Environment variables ---
  SAKULI_HOME:                  C:\Program Files\sakuli\sakuli-v0.9.3-SNAPSHOT
  MOZ_DISABLE_OOP_PLUGINS:      1
  MOZ_DISABLE_AUTO_SAFE_MODE:   1
  MOZ_DISABLE_SAFE_MODE_KEY:    1

    ```
* -help
->
 ```
 D:\go\src\github.com\ConSol\sakuli-go-wrapper>go run main.go -help
Generic Sakuli test starter.
2016 - The Sakuli team / Philip Griesbacher.
http://www.sakuli.org
https://github.com/ConSol/sakuli
Usage:   sakuli[.exe} COMMAND [OPTIONS]
         sakuli -help
         sakuli -run <sakuli suite> [OPTIONS]
         sakuli -encrypt <secret> [OPTIONS]
         ...
Commands:
         run
         encrypt
Options:
         -loop <minutes>           Loop this suite, wait n seconds between executions, 0 means no loops (default: 0)
         -javaHome <folder>        Java bin dir (overrides PATH)
         -preHooks <programpath>   A programm which will be executed before sakuli (Can be added multiple times)
         -postHooks  <programpath> A programm which will be executed after sakuli (Can be added multiple times)
         -D <JVM options>          JVM option to set a property on runtime, overrides the 'sakuli.properties'
         -browser <browser>        Browser for the test execution (default: Firefox)
         -interface <interface>    Network interface used for encryption
         -sahiHome <folder>        Sahi installation folder
         -version                  Version info
 ```
* Options:
    * loop
    * browser <browser>
    * preHook   => file   (or text von sakul.sh -h)
    * postHook  => file
    * javahome  => directory
    * ....
* help:   
    run => directory
    sahi_home => directory
    `sakuli COMMAND VALUE [options]`
    COMMAND: `run` and `encrypt`
        => `sakuli run test_suite -browser chrome`
        => `sakuli encrypt test-secrect [-interface eth0]`

* (SM or TS) usage: example aufrufe hinzufügen
* sakuli -encrypt "test" => `run` or `encrypt` have to contain at least one value
* Fallback if sakuli_home is not set, use the top-folder of the binary, with logging
  -> using folder which is one level above.
* remove `-sakuli_home`
  -> done


help style like docker:
```
 [tschneck:sakuli]# docker
 Usage: docker [OPTIONS] COMMAND [arg...]
        docker daemon [ --help | ... ]
        docker [ --help | -v | --version ]

 A self-sufficient runtime for containers.

 Options:

   --config=~/.docker              Location of client config files
   -D, --debug                     Enable debug mode
   -H, --host=[]                   Daemon socket(s) to connect to
   -h, --help                      Print usage
   -l, --log-level=info            Set the logging level
   --tls                           Use TLS; implied by --tlsverify
   --tlscacert=~/.docker/ca.pem    Trust certs signed only by this CA
   --tlscert=~/.docker/cert.pem    Path to TLS certificate file
   --tlskey=~/.docker/key.pem      Path to TLS key file
   --tlsverify                     Use TLS and verify the remote
   -v, --version                   Print version information and quit

 Commands:
     attach    Attach to a running container
     build     Build an image from a Dockerfile
     commit    Create a new image from a container's changes
     cp        Copy files/folders between a container and the local filesystem
```
