Sistema de Acompanhamento Sintomático (USP Saúde)
--------
Devido aos cuidados e às preocupações resultantes dos perigos da pandemia de Covid-19, pensamos que a comunidade universitária precisa de uma ferramenta que ajude a controlar o acesso, acompanhar e orientar as pessoas. Uma aplicação bem sucedida dessa metodologia ocorreu em alguns lugares do mundo, como na Coréia do Sul, onde houve o achatamento da curva de contágio com a ajuda de ferramentas de monitoramento de contato social. \
O objetivo do sistema proposto é orientar e ajudar os alunos, professores e funcionários da comunidade universitária por meio do acompanhamento de quadros sintomáticos, situação vacinal e exames. Além disso, o sistema permitiria enviar notificações para os usuários que tiveram contato com indivíduos que testaram positivo. Para tal, pensamos inicialmente nos seguintes usos para os diferentes bancos de dados:

Links úteis
----
Estamos construíndo os esquemas conceituais das nossas base de dados, que podem ser acessados pelos links abaixo:
- [Banco de dados de documentos (MongoDB)](https://lucid.app/lucidchart/939746f5-38e5-4a29-9721-b0783e184e40/edit?invitationId=inv_85e84339-061e-4765-ac10-c51ca6d9fffa)
- [Banco de dados relacional (PostgreSQL)](https://lucid.app/lucidchart/843e66f8-f029-49cd-8bf7-0d8a65f6b385/edit?invitationId=inv_fc2ca232-6bcf-4a00-862f-46e2be9dc952)
- [Banco de dados de Grafos (Neo4J)](https://lucid.app/lucidchart/422e7106-4ef4-41e3-83c2-07dbb354022c/edit?invitationId=inv_dfef46f2-1964-4f8f-831b-11a692c9020d) 
- [Esclarecimento do funcionamento dos BDs](https://docs.google.com/document/d/1ogJMdn4OgbnGiW7WrHYYl2R29OkIyDyG7kjJ5a6FRAU/edit?usp=sharing)
- [Apresentação](https://docs.google.com/presentation/d/1jRMrtB4VZ2IffKIPk2C2ZEkr3Ck8qOOpAB04Yc3mH2A/edit?usp=sharing)

Conecte-se aos bancos de dados que criamos a partir das URLs:
- PostgreSQL (Microsoft Azure):  
```jdbc:postgresql://sas-mac439.postgres.database.azure.com:5432/postgres?user=mac439&password=Atila_Iamarino&sslmode=require```
- MongoDB (Atlas):  
```mongodb+srv://mac439:Atila_Iamarino@mac439.fsau9s1.mongodb.net/?retryWrites=true&w=majority```
- Neo4J (Aura):  
```neo4j+s://3184388f.databases.neo4j.io```
    - user: neo4j
    - pwd: ArbFLmo6VLVl1qHLUMzQalzXmOwBlsFqpyafiX7a218
- Redis (RedisLabs):  
```redis-10242.c83.us-east-1-2.ec2.cloud.redislabs.com:10242```
    - user: default
    - pwd: GVNLKjX374wY9OL9iJ2qZlEMw3xTsSJA


# Conectando ao servidor

Para subir o servidor precisa instalar a linguagem Go em sua máquina, isso pode ser feito através do [link](https://go.dev/doc/install). A versão utilizada no projeto está descrito no arquivo `go.mod`.
Em seguida, rode o servidor através do comando `go run server.go` que executará o arquivo `server.go`.
Para checar se a aplicação está funcionando corretamente, após o servidor estar em pé, abra em seu navegador o link `http://localhost:1323/healthz`.

Execute com o comando:
  ```CompileDaemon -command="./backend.exe"```