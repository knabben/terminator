#!/bin/bash

unset DEBUG

function runService() {
    SERVICE=$1
    NAME=$1
    CONTAINER=$2
    PARAMS=$3

    if ! helm list ${SERVICE} | grep -q ${SERVICE} ; then
        helm install --name ${NAME} ${PARAMS} ${CONTAINER}

    else
        echo "Service ${SERVICE} already up."
    fi
}

# --- RabbitMQ

RABBITMQ_USER="admin"
RABBITMQ_PASS="admin"

runService "rabbitmq" "stable/rabbitmq" "--set rabbitmqUsername=${RABBITMQ_USER}
                                         --set rabbitmqPassword=${RABBITMQ_PASS}"


# --- PostgreSQL 10

PSQL_PASS="postgres"
runService "postgis" "stable/postgresql" "--set image=knabben/postgis-sniffer
                                          --set imageTag=latest
                                          --set postgresPassword=${PSQL_PASS}"
