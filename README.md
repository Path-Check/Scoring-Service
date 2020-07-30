# Scoring Service

## Link to Trello Board

- <https://trello.com/b/dA1ZZcHp/scoring-service>

## Project Layout

    - Note: this will change over time
    .
    ├── LICENSE
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── log
    │   ├── Dockerfile
    │   └── main.go <- Entry point for Log sidecar service
    ├── pb <- Directory for the proto files
    │   ├── linasnotificationdraft.proto
    │   ├── log.pb.go
    │   ├── log.proto
    │   ├── notification.pb.go
    │   └── notification.proto
    ├── protocompile.sh
    └── server
        ├── Dockerfile
        └── main.go <- Main entry point for Notification Service

    3 directories, 14 files

## TODO

    - Finalize log.proto

        - The following are Lina's feedback and will be incorporated over time:

            - specific structure of requests/responses/logs will change as per docs
            - log file name: How about opening a new file with the current date/time to ms resolution as part of the file name? That way, we can easily write from more than one instance of the service without worrying about file inconsistencies. With this we’d only open as writeonly, never append.
            - Open & close new files periodically… no need to open anew every time a new request comes in, but do close and open a new one at least every day or so, so that files will be finalized and done.
            - log.fatal: I believe this would exit the service? Unless this is a permanent error that it’s impossible to recover from, I suggest trying to recover gracefully and logging an error so that we can detect it, but don’t completely halt the service. In this case, I would say if we can’t open a file on startup of the service, definitely log.fatal. But if we’re having some kind of write error in the middle of operation, and we can still return a response to the client, we want to make sure we find out there’s a problem but we don’t want to kill the service thus rendering the app not able to notify.

        - Logs should output to storage bucket (s3, google storage buckets, etc)

            - Daily batch jobs:
                
                - New logfiles will be created on a daily basis with a timer to rotate logs according to date with ms resolution
                    - will be write only

            - Alternative:

                - ship logs to Mongo with a capped collection (this ensures that logs are ephemeral with respect to storage limits in order to keep costs low)
                
                    - Postgres can also be used here but a separate cleanup service will be needed (unless postgres has an equivalent of a capped collection)

                - daily batch jobs will export a backup copy to a storage bucket as well as the big data tool of choice (bigQuery, redshift, etc)
                
                    - This ensures modularity, scalability, and high availability

        - Use Context / Cancellation for saveToFile()

        - Update log.fatal() accordingly
    
    - Finalize notification.proto
        
        - Get a greenlight for the datastructures (depends on diagram)

    - Add a diagram
        
        - Finalize Dataflow

## Notes

    - Everything is subject to change
        
        - proto3 can be refactored to proto2

    - Once the project seed is finished, Dev/Staging/Prod branches will be created
