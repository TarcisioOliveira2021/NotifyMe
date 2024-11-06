# NotifyMe
Aplicação de notificação de novos lançamentos de álbuns no spotify

- Refinamento de email : https://github.com/emailmonday/Cerberus/tree/main.
- Lembrar de mencionar sobre habilitar o appPassword na conta do google.
- API MAIL: https://api-docs.mailtrap.io/docs/mailtrap-api-docs/67f1d70aeb62c-send-email-including-templates

COMANDOS DOCKER

( Em /NotifyMe)
- docker network create goapp
- docker build -f app/Dockerfile -t notifyme-app . 
- docker build -f notifymepooling/Dockerfile -t notifyme-poolingapp .
- docker run --name notifyme-app --network goapp -p 8080:8080 -it notifyme-app 
- docker run --name notifyme-poolingapp --network goapp -it notifyme-poolingapp