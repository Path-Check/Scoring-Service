# Scoring Service

## Link to Trello Board

- <https://trello.com/b/dA1ZZcHp/scoring-service>

## Important Links and References

- <https://www.google.com/covid19/exposurenotifications/>
- <https://developers.google.com/android/exposure-notifications/exposure-notifications-api>
- <https://developer.apple.com/documentation/exposurenotification/enexposureconfiguration>
- <https://github.com/google/exposure-notifications-server>
- <https://github.com/google/exposure-notifications-verification-server>
- <https://www.aphlblog.org/bringing-covid-19-exposure-notification-to-the-public-health-community>
- <https://www.reuters.com/article/us-health-coronavirus-britain-tracing-ap/britains-covid-19-app-the-game-changer-that-wasnt-idUSKBN2400YM>
- < https://github.com/google/exposure-notifications-server/tree/main/examples/export>
- <https://github.com/google/exposure-notifications-server/blob/main/examples/export/testExport-2-records-1-of-1.zip>

## Project Layout

    - Note: this will change over time
    .
    ├── go.mod
    ├── go.sum
    ├── LICENSE
    ├── log
    │   ├── Dockerfile
    │   └── main.go     <- Entry point for Logger
    ├── pb
    │   ├── notification.pb.go
    │   └── notification.proto
    ├── protocompile.sh
    ├── README.md
    └── server
        ├── Dockerfile
        └── main.go     <- Entry point for Notification Server

    3 directories, 11 files

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
        
        - Greenlight for the datastructures (depends on diagram)

    - Add a diagram
        
        - Finalize Dataflow

## Notes

    - Everything is subject to change
        
        - proto3 can be refactored to proto2

    - Once the project seed is finished, Dev/Staging/Prod branches will be created
