* -version 
    ```
    --- Environment variables ---
    SAKULI_HOME:                  /home/tschneck/sakuli/sakuli-v0.9.3-SNAPSHOT
    MOZ_DISABLE_OOP_PLUGINS:      1
    MOZ_DISABLE_AUTO_SAFE_MODE:   1
    MOZ_DISABLE_SAFE_MODE_KEY:    1
    ```

* -help
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
* remove `-sakuli_home`


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