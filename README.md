um judger ainda em desenvolvimento

roda (por enquanto) código somente em python com docker

tem fila em memória pras requests de execução e retorno por webhook

### rota POST /worker
body:
```json
{
  "code": string,
  "input": string
  "webhook_url": string
}
```

response:
```json
{
  "token": string,
	"message": string
}
```

### rota GET /worker?token={token}

response:
```json
{
  "ID": string
	"status": string 
	"stdout": string 
	"stderr": string
	"error": string 
}
```

o post pro webhook retorna o mesmo que o GET por token

**próximas alterações?**
- [ ] múltiplas linguagens
- [ ] cache da fila permanente
- [ ] timeout na requisição de execução
