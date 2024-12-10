# Planejja API

A **Planejja API** é a solução para o backend desenvolvida para o projeto de TCC utilizando Golang (GO) e necessária para obtenção da nota na matéria de Backend.

## Índice

 - [Tecnologias](#tecnologias)
 - [Instalação](#instalação)
 - [Estrutura do projeto](#estrutura-do-projeto)
 - [Informações extras](#informações-extras)

## Tecnologias

- **Docker**: Para rodar o projeto em um container.
- **Docker Compose**: Para facilitar a orquestração dos containers.
- **Make**: Para automação de tarefas.
- **Git**: Para controle de versão.
- **MySQL**: Para o banco de dados (v8.0).
- **Tecnologias em Golang**
    - **Go**: Linguagem de programação utilizada (v1.23).
    - **Gin Web Framework**: Utilizada para criação da API RESTful (v1.10).
    - **Gorm**: Biblioteca de ORM para integrar a API ao Banco de dados (v1.25.12).

## Instalação

### 1. Clonar o Repositório

Primeiramente, clone o repositório para o seu ambiente local:

```bash
git clone https://github.com/Jonathanmoreiraa/planejja_api.git
cd planejja_api
```

### 2. Configuração das Variáveis de Ambiente

Crie um arquivo .env na raiz do projeto com as variáveis de ambiente necessárias, baseando-se no arquivo `.env.example`:

```bash
cp .env.example .env
```

Será necessário preencher os espaços em branco desse arquivo .env, visto que nele estarão contidas as credenciais do banco de dados, chaves, informações de JWT, etc.

### 3. Instalar as Dependências

Com o Go instalado (ou utilizando o definido no **docker**), basta rodar o seguinte comando no terminal para baixar as dependências do projeto:

```bash
go mod tidy
```

### 4. Rodar a Aplicação

**Usando o Docker (recomendado)**

Para rodar a aplicação em um container Docker, você pode usar:

```bash
make up
```

Isso irá construir a imagem e rodar os containers de acordo com a configuração definida no arquivo **docker-compose.yaml**.

**Rodando Localmente**

Se preferir rodar localmente sem Docker, basta executar o seguinte comando:

```bash
go run cmd/api/main.go
```

### 5. Acessar a API

Por padrão, a API estará rodando em **http://localhost:8080**.

## Estrutura do Projeto

Realizei a estrutura do projeto da seguinte forma:

```
planejja_api/
├── cmd/
│   ├── api/
│   │   ├── main.go
├── pkg/
│   ├── api/
│   │   ├── handler/
│   │   │   ├── despesa.go
│   │   │   ├── receita.go
│   │   │   ├── reserva.go
│   │   │   ├── user.go
│   │   ├── middleware.go
│   │   │   ├── auth.go
│   ├── config/
│   │   ├── config.go
│   ├── database/
│   │   ├── database.go
│   ├── di/
│   │   ├── wire_gen.go
│   │   ├── wire.go
│   ├── domain/
│   │   ├── categorias_despesas.go
│   │   ├── categorias_receitas.go
│   │   ├── categorias.go
│   │   ├── despesas_parcelas.go
│   │   ├── despesas.go
│   │   ├── receitas.go
│   │   ├── receitas_meses.go
│   │   ├── reservas.go
│   │   ├── user.go
│   ├── repository/
│   │   ├── interface
│   │   │   ├── despesa.go
│   │   │   ├──  receita.go
│   │   │   ├──  reserva.go
│   │   │   ├──  user.go
│   │   ├── despesa.go
│   │   ├── receita.go
│   │   ├── reserva.go
│   │   ├── user.go
│   ├── routes/
│   │   ├── routes.go
│   ├── test/
├── .env.example
├── .gitignore
├── Dockerfile
├── Makefile
├── README.md
├── docker-compose.yaml
├── go.mod
└── go.sum
```

### Entendo a estrutura criada por pastas

Conforme o modelo de estrutura mostrado acima, vou explicar a finalidade das pastas abaixo:

#### **1. /cmd/api**

Contém os arquivos de inicialização e configuração da principal da aplicação, isto é, o arquivo **main.go**.

#### **2. /pkg**
Contém pacotes e módulos de funcionalidade da API.

- **api/**: Responsável pelos ***handlers*** e ***middleware*** (autenticação com JWT)
- **config/**: Contém o arquivo que carrega as variáveis de configuração do arquivo .env.
- **database/**: Contém o arquivo que fará a conexão com o banco de dados e criará as tabelas automaticamente (caso não existam).
- **di/**: Contém os arquivos que auxiliarão na criação da injeção de dependência
- **domain/**: Contém os arquivos que serão utilizados como base para colunas na criação das tabelas do banco de dados.
- **repository/**: Contém os arquivos responsáveis por lidar com a persistência de dados, abstraindo a interação com o banco de dados ou qualquer outra fonte de dados. Além disso, possui a pasta **interfaces/** que contém os arquivos com a interfaces desses outros arquivos.
- **routes/**: Contém o arquivo que cria e configura as rotas.
- **repository/**: Contém os arquivos responsáveis por coordenar a execução das ações com base nas entradas do usuário, é onde fica a lógica de negócio principal. Além disso, possui a pasta **interfaces/** que contém os arquivos com a interfaces desses outros arquivos.
- **test/**: Contém os testes (tanto de integração quanto unitários).

#### **3. Demais arquivos**
Os arquivos localizados fora das pastas são utilizados para criação de outros arquivos (``.env.example`` e ``.env``), arquivos para funcionamento do contêiners (``Dockerfile`` e ``docker-compose.yaml``) e arquivos para buildar o projeto (``go.sum`` e ``go.mod``).

## Endpoints

#### **Usuários**:

- **POST** /login
- **GET** /api/users
- **GET** /api/users/:id
- **POST** /api/users
- **DELETE** /api/users/:id

#### **Receitas**:

- **GET** /api/receitas
- **GET** /api/receitas/:id
- **POST** /api/receitas
- **DELETE** /api/receitas/:id

#### **Despesas**:

- **GET** /api/despesas
- **GET** /api/despesas/:id
- **POST** /api/despesas
- **DELETE** /api/despesas/:id

#### **Reservas**:

- **GET** /api/reservas
- **GET** /api/reservas/:id
- **POST** /api/reservas
- **DELETE** /api/reservas/:id

## Informações extras
Caso ocorra o seguinte erro na conexão com o banco de dados ``Public Key Retrieval is not allowed`` será preciso alterar a propriedade ``allowPublicKeyRetrieval`` para **true** na aba **Propriedades do driver** no DBeaver ou equivalente.