# Backend

O código contido neste recorte do projeto contém o servidor web responsável por cumprir todas as user stories relacionadas às solicitações de remédios, existe também uma interface de mensageria responsável por garantir que os pedidos serão tratados de forma ordenada e assíncrona, além de uma serviço que serve de consumer desa fila kafka e armazena no banco de dados ( postgres ) as informações dos pedidos. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1]. Ademais, também foi implementada, dado requisito de alta escalabilidade definido para o projeto, a arquitetura [hexagonal](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) [^2]

## Dependências e Serviços

Antes de continuar, é necessário instalar as dependências e criar os serviços listados para a execução dos comandos posteriores. Para isso siga as seguintes instruções:

- Docker engine - [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
- Build Essential - [What is Build Essential Package in Ubuntu?](https://itsfoss.com/build-essential-ubuntu/)

## Como rodar o sistema

Siga as intruções abaixo para rodar o sistema junto a todos os seus recortes, simulação, mensageria, banco de dados e vicualização com o Metabase.

### Definir as variáveis de ambiente:
Rode o comando abaixo e preencha com as respectivas variáveis de ambiente o arquivo `.env` criado.

#### Comando:
```shell
make env
```

#### Output:
```shell
================================================= START OF LOG ===================================================
Environment file created at ./.env
================================================== END OF LOG ====================================================
```

> [!NOTE]
> Antes de preencher o arquivo `.env` é necessário criar os serviços de cloud presentes nas seção [#Dependências e Serviços](https://github.com/henriquemarlon/pond-micro-cam/backend/tree/main/backend#depend%C3%AAncias-e-servi%C3%A7os)

### Criar infraestrutura local
Rode o comando abaixo para criar a infraestrutura necessária para o sistema localmente, os comandos para orquestrar os containers respectivos estão sendo abstraídos por um arquivo Makefile, para saber mais detalhes do comando abaixo acesse o [link](https://github.com/henriquemarlon/pond-micro-cam/backend/blob/main/backend/Makefile#L11).

#### Comando:
```shell
make infra
```

#### Output:
```shell
================================================= START OF LOG ===================================================
[+] Running 6/6
 ✔ Network 2024-1b-t02-ec10-g04_default  Created                                                              0.0s 
 ✔ Container zookeeper                   Started                                                              0.1s 
 ✔ Container redis                       Started                                                              0.1s 
 ✔ Container postgres                    Started                                                              0.1s 
 ✔ Container kafka                       Started                                                              0.1s 
 ✔ Container control-center              Started                                                              0.0s 
Creating kafka topics...
Created topic orders.
================================================== END OF LOG ====================================================
```

> [!NOTE]
> O arquivo Docker Compose que é chamado pelo comando acima define um ambiente composto por diversos serviços. Inclui o Zookeeper para coordenação, Kafka para mensagens, Control Center para gerenciamento, PostgreSQL como banco de dados relacional e Redis como armazenamento em cache. Cada serviço é configurado com suas respectivas imagens, variáveis de ambiente e opções de rede. O Kafka, por exemplo, é configurado com detalhes como ID do broker, conexão com o Zookeeper e listeners para comunicação interna e externa. O PostgreSQL e o Redis são configurados com volumes para persistência de dados. Este arquivo proporciona um ambiente completo para desenvolvimento e execução de aplicativos que exigem sistemas de mensageria, banco de dados e armazenamento em cache.

### Rodar o sistema:

Mais uma vez, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/henriquemarlon/pond-micro-cam/backend/blob/main/backend/Makefile#L11).

#### Comando:

```bash
make run
```

#### Output:

```shell
================================================= START OF LOG ===================================================
[+] Running 8/8
 ✔ Network deployments_backend        Created                                                                 0.1s 
 ✔ Container deployments-server-3     Started                                                                 0.1s 
 ✔ Container deployments-conductor-1  Started                                                                 0.1s 
 ✔ Container deployments-server-1     Started                                                                 0.1s 
 ✔ Container deployments-server-2     Started                                                                 0.1s 
 ✔ Container deployments-conductor-2  Started                                                                 0.1s 
 ✔ Container deployments-conductor-3  Started                                                                 0.1s 
 ✔ Container nginx                    Started                                                                 0.0s 
================================================== END OF LOG ====================================================
```

> [!NOTE]
> O arquivo Docker Compose que é chamado pelo comando acima configura três serviços: nginx, server e conductor. O serviço nginx utiliza a imagem mais recente do Nginx, mapeia a porta 80 do host para o contêiner, substitui a configuração padrão do Nginx com um arquivo personalizado e depende dos serviços "server" e "conductor". Os serviços "server" e "conductor" carregam variáveis de ambiente de um arquivo .env, reiniciam automaticamente, são construídos a partir de Dockerfiles específicos e são implantados com três réplicas cada. Todos os serviços estão conectados à rede "backend", que facilita a comunicação entre eles.

## Demonstração do Sistema

A demonstração foi feita testando as rotas através do [Swagger UI](https://swagger.io/tools/swagger-ui/) servido pelo serviço "server" na rota `http://localhost/api/v1/docs/index.html#`.

[Demonstração](https://drive.google.com/file/d/1R9fG24_uBr8LlNc11TM6hkP8Jo1NESSZ/view?usp=sharing)

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.
[^2]: As entidades, repositórios e use cases estão de acordo com os padrões previstos para a arquitetura hexagonal.
