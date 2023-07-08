cd ..
docker compose down
docker compose -f docker-compose.dev.yml up db crawler frontier-acceptor logs --build --remove-orphans -d