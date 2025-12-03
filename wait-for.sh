#!/bin/sh
# wait-for.sh - a POSIX-compliant script to wait for services to be ready.

set -e

# Loop until we find the '--' argument
while [ "$1" != "--" ]; do
    SERVICE=$1
    # Use cut to split the service into host and port, which is POSIX compliant
    HOST=$(echo "$SERVICE" | cut -d: -f1)
    PORT=$(echo "$SERVICE" | cut -d: -f2)
    
    echo "Waiting for $SERVICE..."
    
    # Use a loop with nc to check if the port is open
    until nc -z "$HOST" "$PORT"; do
      echo "Service $SERVICE is unavailable - sleeping"
      sleep 1
    done
    
    echo "Service $SERVICE is up."
    shift
done

# Shift past the '--' argument
shift

# Execute the rest of the command
echo "All services are up - executing command: $@"
exec "$@"