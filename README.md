Generate image
1. cd backend;
2. docker build -t mateusreinert/backend-kafka:1.0 .

Send image to dockerhub
1. docker push mateusreinert/backend-kafka:1.0



# Demonstration

Run docker compose
1. See your server machine IP
2. Open docker-compose.yml
3. Change SERVER_HOST_AND_PORT vairable to have your server machine ip
4. docker-compose up
5. Demo to the class