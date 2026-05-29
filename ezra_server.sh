#! /bin/bash

start_or_run () {
    docker inspect ezra_server > /dev/null 2>&1

    if [ $? -eq 0 ]; then
        echo "Starting Ezra Hub Server container..."
        docker start ezra_server
    else
        echo "Ezra Hub Server container not found, creating a new one..."
        docker run --name ezra_server --env-file ./app/server/.env --add-host=host.docker.internal:host-gateway  -p 3294:3294 ezrahub_server:latest 
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping Ezra Hub Server container..."
        docker stop ezra_server
        ;;
    logs)
        echo "Fetching logs for Ezra Hub Server container..."
        docker logs -f ezra_server
        ;;
    build)
        echo "Building Ezra Hub Server container..."
        go build -C ./app/server -o ezra_server
        docker build ./app/server -t ezrahub_server:latest
        ;;
    *)
        echo "Usage: $0 {start|stop|logs|build}"
        exit 1
esac