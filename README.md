# Golendar

Simple calendar web app written in Go.

### How to use

Use `go run main.go` to start server without installation

### API

#### Endpoint /events 

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

Provide event data as JSON. Note that all field will be updated, not only the field present in provided JSON. If there is no event with the provided UUID, it will be created. 

**Delete event**

Method: DELETE  
Params: UUID (required)  
Example: `DELETE /events?UUID=3405be29-2852-44fa-b8c1-5c5be22d0c0f HTTP/1.1`  

