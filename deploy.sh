sudo apt install docker.io;
sudo apt install docker-compose;

sudo service docker stop;
sudo usermod -a -G docker ubuntu;
sudo service docker start;

mkdir ./data;
mkdir ./data/server;
mkdir ./data/server/cert;

sudo openssl req -out ./data/server/cert/CSR.csr -new -newkey rsa:2048 -nodes -keyout ./data/server/cert/privateKey.key;
sudo openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout privateKey.key -out ./data/server/cert/certificate.crt;

curl https://raw.githubusercontent.com/PolytechLyon/cloud-project-equipe-8/master/docker-compose.yml -o docker-compose.yml;

docker-compose pull;

docker-compose up -d;
