cd ..
docker compose down
docker compose -f docker-compose.dev.yml up db ranker bert logs --build --remove-orphans