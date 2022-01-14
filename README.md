# golang-diary

*This is a project mostly meant for my own learning. I decided to write something that would make use of various technologies, even though in this particular instance they may not make the most sense (such as a standalone postgres db instead of embedded sqlite), they accomplish their goal in allowing me to delve into these concepts.*

golang-diary is a self-hosted diary web application that facilitates tracking thoughts, feelings, and whatever events might have occurred in your day. It sports a simple and intuitive web interface that does not clutter nor overwhelm.

The backend API of the application is written in Go (hence the name), while the frontend is written utilizing the [Svelte JS](https://svelte.dev/) javascript framework/compiler. Additionally I made use of the [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc) utility to facilitate code generation for DB operations. Finally, I am utilizing [postgreSQL](https://www.postgresql.org/) for data storage.



## API Reference

#### Get all entries

```http
  GET /api/entry/list
```
#### Get entry

```http
  GET /api/entry/get?id=${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of entry to fetch |

#### Add new entry
```http
  POST /api/entry/add
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Title of diary entry |
| `message`      | `string` | **Required**. The body or content of diary entry |
| `mood`      | `[]string` | Mood or moods (feelings) at the time of entry (zero, one or many) |

#### Update existing entry
```http
  POST /api/entry/update
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. ID of entry to be modified |
| `title`      | `string` | **Required**. Title of diary entry |
| `message`      | `string` | **Required**. The body or content of diary entry |
| `mood`      | `[]string` | Mood or moods (feelings) at the time of entry (zero, one or many) |

```http
  POST /api/entry/delete
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of entry to delete |

# License
This project is licensed (or unlicensed) by [The Unlicense](https://github.com/davidedpg10/golang-diary/blob/master/LICENSE), as it's mostly a learning vehicle for me. You are free to copy, modify, distribute this content at will, with or without attribution.