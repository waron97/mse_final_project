cd ..
docker compose down
docker compose -f docker-compose.yml up db crawler frontier-acceptor logs --build --remove-orphans -d