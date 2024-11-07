# NotifyMe
Aplicação que envia notificações por email sobre o lançamento de novos álbuns de um artista no Spotify, com uma verificação diária.

### Observações
1- Caso tenha escolhido o servidor do google, precisa habilitar o appPassword na conta do google. (Outros serviços podem ter sua própria necessidade) essas infos vão ser necessárias para preencehr o `EMAIL_USERNAME` e `EMAIL_PASSWORD`.

2- Caso queira testar o envio do email em um ambiente de testes, pode usar o serviço do Mailtrap (https://mailtrap.io/)
3- É necessário criar uma conta no Spotify dev (https://developer.spotify.com/) para usar a api, é adinicionar no `.env` o `CLIENT_ID` e `CLIENT_SECRET` informados no Dashboard.
4- Para descobrir o ID do artista basta abrir a página dele no spotify e pegar o código no path da url após o `/artist/`, por exemplo: Nessa URL: https://open.spotify.com/intl-pt/artist/0Riv2KnFcLZA3JSVryRg4y, o código do artista é: `0Riv2KnFcLZA3JSVryRg4y` 

### Para rodar o projeto
    - Basta clonar o projeto
        - Para rodar local: (necessário ter o Go instalado na máquina)
            - Executar o comando `go mod tidy` em `NotifyMe/` (Vai baixar as dependências necessárias) 
            - É necessário rodar primeiro o pacote `./app` via `go run main.go`, depois executar o pacote `notifymepooling` usando o `go run main.go`

        - Caso tenha o docker basta rodar `docker-compose up` em `NotifyMe/'

### Estrutura do arquivo .env
- Deve ser colocado na raiz do projeto.
    `CLIENT_ID=XPTO`

    `CLIENT_SECRET=XPTO`

    `ARTIST_ID=XPTO`

    `API_URL=http://localhost:8080/notification`

    `EMAIL_TO=`

    `EMAIL_FROM=`

    `EMAIL_HOST=`

    `EMAIL_PORT=`

    `EMAIL_USERNAME=`                               

    `EMAIL_PASSWORD=`


### Roadmap para futuras implementações
- Refinamento do emailBody usando os templates: https://github.com/emailmonday/Cerberus/tree/main.
- Permitir monitorar múltiplos artistas
