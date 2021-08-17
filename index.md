# Welcome to Summit API Day

Here you will find all information to get started with building and deploying your web service on heroku.

## Pre-Requisites

* Postman installed on your workstation

## Lab Docs

### Lab 1

Use the provided API endpoint to create a movie entry. ​Use your name for the 'Name' criteria.

**Hint:** *Utilize the movies API documentation below.*

### Lab 2

* Use the provided API endpoint to ensure your movie shows up in the movie catalog.

**Hint:** *Utilize the movies API documentation below.*

### Lab 3

* Using the data from Lab 1, update your entry and change the movie name to <Name_API_Day>. Once done, again, ensure your movie shows up in the movie catalog.

**Hint:** *Utilize the movies API documentation below.*

### Lab 4

* Nice work! Turns out the movie you added isn't an actual movie! Use the provided API endpoint to *delete* your movie!

**Hint:** *Utilize the movies API documentation below.*

## API Endpoints

### Base URL

[https://summit-movie-api.herokuapp.com](https://summit-movie-api.herokuapp.com)

### Movies API

#### Get All Movies

| Web Method | End Point |
| ------- | ------ |
| **GET** | /movie |

Returns information for *all* movies.


| Web Method | End Point |
| ------- | ------ |
| **GET** | /movie/<name> |

Returns information for *one* movie where 'name' is the name of the specific movie title.

| Web Method | End Point |
| ------- | ------ |
| **POST** | /movie |

Creates an entry for a movie.

**Example below:**

```json
{
	"name": "Test",
	"ratings": 1,
	"actors": ["Logan Price", "John Sibo", "Obot"],
	"watched": false
}
```

| Web Method | End Point |
| ------- | ------ |
| **DELETE** | /movie/<name> |

Deletes a movie entry where 'name' is the name of the specific movie title.

**Example below:**

```json
{
	"name": "Test"
}
```

### Reset
####  Reset data back to default

| Web Method | End Point |
| ------- | ------ |
| **DELETE** | /reset |

Reset data back to default data.
