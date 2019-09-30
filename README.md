# Golendar

Simple calendar web app written in Go.

## How to use

Use `go run main.go` to start server without installation

## API

### Endpoint /events 

**Get all events**

Method: GET  
Example: `GET /events HTTP/1.1`  

Sample result:   

```json
[
    {
        "UUID": "d4da8d8b-b6fe-438c-a2b5-b34701947660",
        "Title": "Spam",
        "Datetime": "",
        "Duration": "",
        "Desc": "BAZINGA!",
        "User": "",
        "Notify": ""
    },
    {
        "UUID": "183bca1b-3f56-42b0-8cf2-581797849209",
        "Title": "Vookah",
        "Datetime": "",
        "Duration": "",
        "Desc": "You gonna like it",
        "User": "Mack",
        "Notify": ""
    }
]
```

**Create event**

Method: POST  
Example: `POST /events HTTP/1.1`  

Provide event data as JSON:  

```json
{
	"Title": "Wookah",
	"User": "Mack",
	"Desc": "YOU GONNA LIKE THIS"
}
```

**Update event**

Method: PUT  
Params: UUID (required)  
Example: `PUT /events?UUID=3405be29-2852-44fa-b8c1-5c5be22d0c0f HTTP/1.1`  

Provide event data as JSON. Note that all fields will be updated, not only the fields present in provided JSON. If there is no event with the provided UUID, one will be created. 

**Delete event**

Method: DELETE  
Params: UUID (required)  
Example: `DELETE /events?UUID=3405be29-2852-44fa-b8c1-5c5be22d0c0f HTTP/1.1`  

## Using with postgres database

Database connection string must be specified in config file and storage type must be "Postgres".

```yaml
storage: "Postgres"
dburi: "postgres://golendar:golendar@localhost:5432"
```

You can run postgres in docker with the following command:

```bash
docker run --rm --name pg -e POSTGRES_USER=golendar -e POSTGRES_PASSWORD=golendar -d -p 5432:5432 postgres
```

Add -v argument like shown bellow if you want to persist data.

```bash
-v $HOME/docker/volumes/postgres:/var/lib/postgresql/data 
```

## Using notifications sender

Notification sender consists of two daemons which can be found in cmd/daemons folder. It requires RabbitMQ to operate. 

RabbitMQ can be launched in docker:
```bash
docker run -d --hostname rabbit --name rmq -e RABBITMQ_DEFAULT_USER=golendar -e RABBITMQ_DEFAULT_PASS=golendar -d -p 5672:5672 rabbitmq:3
```

As RabbitMQ is ready launch daemons:
```bash
go run cmd/daemons/creator/notificaions_creator.go 
```
```bash
go run cmd/daemons/sender/notifications_sender.go
```

Now if new event will be created (via rest api for example), notification sender will get the message. Notifications are send only once for each event. Time interval for notifications can be set in config.yaml (in minutes). Please mind that discrepancies in docker timezone (for Postgres) and host machine can lead to unwanted behaviour. No timezone control is implemented. 

