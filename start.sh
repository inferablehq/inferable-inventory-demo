#!/usr/bin/env bash

# Function to prefix output with a label
prefix_output() {
    sed "s/^/$1 /"
}

# Start docker compose (pgsql)
(docker compose up -d) 2>&1 | prefix_output "[Docker]" &

# Start static server for inventory assets
(cd inventory/assets && npm i && npx -y node-static -p 5556) 2>&1 | prefix_output "[Static Server]" &

# Start inventory dev server
(cd inventory && npm i &&  npm run dev) 2>&1 | prefix_output "[Inventory Server]" &

# Start customer orders service
(cd customers && go run .) 2>&1 | prefix_output "[Customer Orders]" &

# Wait for all background processes to finish
wait
