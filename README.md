# Scoring Service

## TODO

    - Finalize log.proto
        - The following are Lina's feedback and will be incorporated over time:
            - specific structure of requests/responses/logs will change as per docs
            - log file name: How about opening a new file with the current date/time to ms resolution as part of the file name? That way, we can easily write from more than one instance of the service without worrying about file inconsistencies. With this we’d only open as writeonly, never append.
            - Open & close new files periodically… no need to open anew every time a new request comes in, but do close and open a new one at least every day or so, so that files will be finalized and done.
            - log.fatal: I believe this would exit the service? Unless this is a permanent error that it’s impossible to recover from, I suggest trying to recover gracefully and logging an error so that we can detect it, but don’t completely halt the service. In this case, I would say if we can’t open a file on startup of the service, definitely log.fatal. But if we’re having some kind of write error in the middle of operation, and we can still return a response to the client, we want to make sure we find out there’s a problem but we don’t want to kill the service thus rendering the app not able to notify.
    
    - Finalize notification.proto
        - Get a greenlight for the datastructures

    - Add a diagram
        - Finalize Dataflow

## Notes

    - Everything is subject to change
        - proto3 can be refactored to proto2
