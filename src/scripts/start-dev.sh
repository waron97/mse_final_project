cd ..
docker compose down
docker compose -f docker-compose.dev.yml up --build --remove-orphans -d