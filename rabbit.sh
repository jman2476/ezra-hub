#! /bin/bash

start_or_run () {
    docker inspect ezra_rabbitmq > /dev/null 2>&1

    if [ $? -eq 0 ]; then
        echo "Starting Ezra Hub RabbitMQ container..."
        docker start ezra_rabbitmq
    else
        echo "Ezra Hub RabbitMQ container not found, creating a new one..."
        docker run -d --name ezra_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.3-management
    fi
}

case "$1" in
    start)
        start_or_run
        ;;
    stop)
        echo "Stopping Ezra Hub RabbitMQ container..."
        docker stop ezra_rabbitmq
        ;;
    logs)
        echo "Fetching logs for Ezra Hub RabbitMQ container..."
        docker logs -f ezra_rabbitmq
        ;;
    *)
        echo "Usage: $0 {start|stop|logs}"
        exit 1
esac