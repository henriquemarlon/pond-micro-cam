# Backend

## Como Rodar

```bash
make env
```

```bash
make run
```

# Descrição:

Por trás dos panos estão sendo criados dois clusters com 3 replicas para cada serviço ( user e image ) e está sendo utilziado o nginx como proxy reverso e load balancer. Ademais, toda aplicação está virtualizada com Docker e orquestrada por um docker compose. P sagger está disponível na rota: http://localhost/api/v1/docs/index.html#

# Logs:

Os logs estão sendo armazenados em uma tabela no db postgres:
![image](https://github.com/henriquemarlon/pond-micro-cam/assets/89201795/f9297619-f8b5-4e49-893e-1ba163d85c4d)
