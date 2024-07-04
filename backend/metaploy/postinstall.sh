#!/bin/sh

cleanup() {
    echo "Container stopped. Removing nginx configuration."
    rm /etc/nginx/sites-enabled/naarad-api.metaploy.conf
}

trap 'cleanup' SIGQUIT SIGTERM SIGHUP

"${@}" &

cp ./naarad-api.metaploy.conf /etc/nginx/sites-enabled

wait $!
