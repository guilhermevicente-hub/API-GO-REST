# Projeto de API REST em Go

Este projeto é parte do meu estudo contínuo da linguagem Go. Aqui, estou desenvolvendo uma API REST utilizando Go, conectando-a a um banco de dados hospedado em Docker. Estou utilizando o ORM (Object-Relational Mapping) do Go para facilitar a conexão com o banco de dados e manipular as requisições de forma mais eficiente, sem a necessidade de escrever SQL manualmente.

# Objetivo

O objetivo deste projeto é demonstrar minhas habilidades em desenvolvimento de APIs REST com Go e meu conhecimento em integração com bancos de dados usando Docker para ambientes de desenvolvimento.

# Tecnologias Utilizadas

- Linguagem: Go
- Banco de Dados: PostgreSQL (Docker)
- ORM: [Nome do ORM utilizado, se aplicável]
- Outras Dependências: [Liste outras dependências importantes]


# Pré-requisitos

Antes de começar, certifique-se de ter instalado em sua máquina:

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/) (opcional, se você já possui um banco de dados PostgreSQL instalado localmente)
# Como executar

1. Clone este repositório:


```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
```

2. Navegue até o diretório do projeto:

```bash
   cd seu-repositorio
```

3. Execute o banco de dados PostgreSQL usando Docker:

```bash
   docker-compose up -d
```

4. Execute a aplicação Go:

```bash
  go run main.go
```



Agora você pode acessar a API em **http://localhost:PORTA**, onde **PORTA** é a porta configurada para sua aplicação Go.


## Capitulo 1 - Json, Request, Response e Go

- Entendenmos o que é uma API
- Realizamos uma requisição GET retornando um Json.

## Capitulo 2 - Roteador, recursos por ID e Docker

- Adicionamos o [Gorilla Mux](https://github.com/alura-cursos/api-go-rest/blob/aula_2/routes/routes.go) como novo roteador da nossa aplicação;
- etornamos um [único recurso](https://github.com/alura-cursos/api-go-rest/blob/aula_2/controllers/controllers.go) através do id;
- Criamos uma imagem do banco de dados Postgres com Docker e executamos um [script SQL que adicionava alguns registros em nosso banco de dados](https://github.com/alura-cursos/api-go-rest/blob/aula_2/migration/docker-database-initial.sql).

## Capitulo 3 - Conexão com banco e exibindo os dados

- Instalamos o Gorm
```bash
go get -u gorm.io/gorm
```
- Realizamos a conexão da API com banco de dados
    >  criei uma pasta chamada “database” e dentro dela um arquivo chamado “db.go”. Esse arquivo “db.go” vai ser responsável por conectar com o banco de dados.
- Alteramos as funções do controller para exibir as informações do banco de dados.

> Usando o "Find()" buscamos todas as personalidades passando o endereço de memória da estrutura que temos, que queremos exibir e para retornarmos apenas uma utilizamos o "DB.First" para ele voltar passando o endereço de memória da personalidade, da pessoa que estamos buscando e o Id.

## Capitulo 4 - Criando, deletando e editando com Gorm

- Adicionamos um endpoint com método Post para criar uma nova personalidade e salvá-la no banco de dados;
- Adicionamos um endpoint com método Delete para deletar uma personalidade e removê-la do banco de dados;
- Adicionamos um endpoint com método Put para atualizar uma personalidade e alterá-la no banco de dados.


# Contribuições
Contribuições são bem-vindas! Se você encontrar um problema, tiver uma ideia ou quiser contribuir de qualquer forma, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Licença
Este projeto está licenciado sob a Licença MIT.
