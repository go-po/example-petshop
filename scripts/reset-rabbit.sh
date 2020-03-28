#!/bin/sh

IFS=$'\n'
for q in $(docker exec petstore_mq rabbitmqctl -qs list_queues name); do
  docker exec petstore_mq rabbitmqctl delete_queue "${q}"
done;
