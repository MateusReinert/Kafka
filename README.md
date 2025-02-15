Boa noite professor, segue aqui as configurações necessárias para rodar o projeto.

1. Download docker desktop: https://www.docker.com/products/docker-desktop/
2. Installe o docker

Possível complicação:
caso ocorra erro de WSL update, execute no cmd do windows (wsl --update).

Para rodar:

3. abra um terminal entre na pasta backend (cd backend) e rode: docker-compose up
4. Em seguida, abra um segundo terminal e rode:
docker run --rm --net=host -it -e NGROK_AUTHTOKEN=2t08pG5Crp5YVZ1mj9IpZw9aPBD_6PXHDnR9DgWeSS5WKdtTG -v C:\Users\Mateus\Repositorios\KAFKA\ngrok.yml:/var/lib/ngrok/ngrok.yml ngrok/ngrok:latest start backend frontend

OBSERVAÇÃO: substitua esse bloco: C:\Users\Mateus\Repositorios por onde está esse documento no computador do professor.

5. após rodar, ele irá gerar dois links:

exemplo:

https://6c69-45-160-36-147.ngrok-free.app --> localhost:3000
https://6c69-45-160-36-147.ngrok-free.app --> localhost:4000

copie o link do localhost:4000 (IMPORTANTE QUE SEJA O 4000), neste caso: 6c69-45-160-36-147.ngrok-free.app

entre no arquivo dentro do projeto: kafka/docker-compose.yml

e substitua o url dessas duas variáveis pela que você copiou:

REACT_APP_REST_HOST: https://1055-177-51-86-72.ngrok-free.app 
REACT_APP_WS_HOST: wss://1055-177-51-86-72.ngrok-free.app 

6. Abra um terceiro terminal, acesse o frontend: cd frontend e rode dock-compose up

por fim, acesse o link que foi gerado na sua porta localhost:3000 (dessa vez o 3000 e não o 4000)

Para acessar o consumer: apenas o link que está o localhost:3000
Para acessar o producer: adicione o mesmo link e inclua no final /producer1

Outro comandos, caso o professor necessite:



Generate image
1. cd backend;
2. docker build -t mateusreinert/backend-kafka:1.0 .
3. cd frontend
4. docker build -t mateusreinert/frontend-kafka:1.0 .

Send image to dockerhub
1. docker push mateusreinert/backend-kafka:1.0
2. docker push mateusreinert/frontend-kafka:1.0