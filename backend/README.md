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