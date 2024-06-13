#!/bin/sh

cleanup() {
	echo "Container stopped. Removing nginx configuration."
	rm /etc/nginx/sites-enabled/naarad.metaploy.conf
}

trap 'cleanup' SIGQUIT SIGTERM SIGHUP

"${@}" &

cp /naarad.metaploy.conf /etc/nginx/sites-enabled

wait $!

echo "lmao"