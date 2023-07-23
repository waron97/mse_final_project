cd ..
docker compose down
docker compose -f docker-compose.dev.yml up db ranker --build --remove-orphans