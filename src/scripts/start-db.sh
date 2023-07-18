cd ..
docker compose down
docker compose -f docker-compose.yml up db --build --remove-orphans -d