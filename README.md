
# Genesis Exchange

Essa API foi construida durante o teste da Genesis Bank.

# Requirements
<ul>
    <li>Go 1.20+</li>
    <li>MySQL</li>
    <li>Docker (optional)</li>
</ul>


## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar criar um arquivo chamado .env e adicionar as seguintes variáveis de ambiente.

| Nome| Descrição|
|-----------|-------------------------------------------------------------------------------------------------------------------------------------|
| MYSQL_ROOT_PASSWORD | A Senha para o usuario root do MySQL|
| DATABASE_DIALECT | Se caso você quiser mudar para postgres ou outra tipo de Database SQL|
| MYSQL_DATABASE  | O nome da Database em que o app irá interagir |
                                                                                        
## Rodando Localmente

**Entre no diretório root**

```bash
  cd <caminho para a pasta com o docker-compose>
```

*Crie um arquivo chamado ".env" com os valores citados acima ou renomeie o arquivo ".env.example" já existente para ".env"*

**Inicie a Aplicação**

```bash
  docker compose up -d
```

### Por Padrão a aplicação já está configurada para criar e popular as tabelas

# Uso

#### Faz a conversão

```vbnet
  POST /exchange/{amount}/{from}/{to}/{rate}
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `amount` | `float64` | O valor a ser convertido |
| `from` | `string` | A moeda original |
| `To` | `string` | A moeda que o valor foi convertido para |
| `rate` | `float64` | A taxa de conversão da moeda |

**Exemplo de Retorno**

```json
{
    "valorConvertido": 63.384, //Resultado da Conversão
	"simboloMoeda": "R$" //Simbolo que representa a moeda que o valor acima representa
}
```

#### Retorna o histórico de todas as exchanges feitas

```vbnet
  GET /exchange/list
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `null` | `null` | Não aceita parametros |

**Exemplo de Retorno**

```json
[
	{
		"id": 1,
		"amount": 11.12,
		"from": "USD",
		"to": "BRL",
		"rate": 5.7,
		"result": 63.38
	}
]
```
## Etiquetas

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)



## Autores

- [@Obdurat](https://www.github.com/Obdurat)

