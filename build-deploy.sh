cd /opt/game
git clone https://github.com/petros-an/server-test.git ./be
cd ./be
docker build . -t game-be:latest
docker compose down
docker compose up -d 


