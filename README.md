Generate image
1. cd backend;
2. docker build -t mateusreinert/backend-kafka:1.0 .
3. cd frontend
4. docker build -t mateusreinert/frontend-kafka:1.0 .

Send image to dockerhub
1. docker push mateusreinert/backend-kafka:1.0
2. docker push mateusreinert/frontend-kafka:1.0

# Demonstration

Run backend
1. cd backend;
2. docker-compose up

Run ngrok
1. docker run --rm --net=host -it -e NGROK_AUTHTOKEN=2t08pG5Crp5YVZ1mj9IpZw9aPBD_6PXHDnR9DgWeSS5WKdtTG -v C:\Users\edson\Repositorios\KAFKA\ngrok.yml:/var/lib/ngrok/ngrok.yml ngrok/ngrok:latest start backend frontend

Run frontend
1. Look at the ngrok URL created for the backend port (4000) (i.e.: https://6c69-45-160-36-147.ngrok-free.app)
2. Open docker-compose.yml
3. Change REACT_APP_REST_HOST variable with that URL (i.e.: REACT_APP_REST_HOST: https://6c69-45-160-36-147.ngrok-free.app)
4. Change REACT_APP_WS_HOST variable with that URL but replace https by wss (i.e.: REACT_APP_WS_HOST: wss://6c69-45-160-36-147.ngrok-free.app)
5. docker-compose up
6. Demo to the class