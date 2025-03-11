# Janjiyan

**Janjiyan** is a personal project aimed to create an application that are capable of managing appointments and invitations as well as their relation with users. This application serve as a back-end application and provide only APIs to be later used by any front-end applications.

## Table of Contents

- [Dependencies](#dependencies)
- [Installation](#installation)
- [Project Structure and Sepcification](#project-structure-and-specification)
  - [Domains](#domains)
  - [Routes](#routes)
  - [Utilities](#utilities)
  - [Relations](#relationship-between-entities)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [Subroutes](#subroutes)

## Dependencies

This project utilize following frameworks and package:

- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [GoDotEnv](https://pkg.go.dev/github.com/joho/Godotenv)
- [jwt-go](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)

## Installation

1. Clone this repository.

2. Run

    ```bash
    go install
    cp .env.example .env
    ```

3. Provide a database for this application and database connection information in the `.env` file

4. Run the `main.go` file as shown as below.

    ```bash
    # run without seeding
    go run path/to/main.go

    # add -s flag to run with seeding
    go run path/to/main.go -- -s
    ```

## Project Structure and Specification

This project is structured as the graphic below.

```txt
.
└── Root/
    ├── domains/
    │   └── [domain]/
    │       ├── model.go
    │       ├── database.go
    │       └── controller.go
    ├── routes/
    │   ├── routes.go
    │   └── [supporting files]
    ├── utilities/
    │   └── [utility]
    └── main.go

```

### Domains

The domains directory contains several packages meant to handle entities involved in this application. There are 3 entities involved in this application: user, appointment, and invitation. Invitation also serve as conjunction table between user and appointment.

Each entities has a model file, database file, and controller file. The model files defines entities structure both for database transaction and http response as well as handling the conversion between both types. Database files handle database transaction for each entities, as well as their assigned relations. Controllers handle data processing, and determines authorization.

### Routes

Routes is a directory that handles routing and middleware. each route (except one) has their dedicated function that extract data needed such as entities ID and token issuer and directly calls the related controller function. Inside the directory, each routes is divided into 4 categories:

- auth (any routes processing authentication process such as login and register)
- user (directly calls user's controller, may pass appointment's and invitation's delete function for cascade deletion)
- appointment (directly calls appointment's controller, may pass invitation's delete function for cascade deletion)
- invitation (directly calls invitation's controller)

### Utilities

Utilities is a directory contains many helper functions to help the application run properly. There are 5 packages inside:

- auth (handles password hashing and JWT token generation and validation)
- database (handle database connection)
- errorlog (handles log recording. Primarily useful for debugging)
- migration (handles database migration and tables update)
- seeder (handle database seeding).

### Relationship Between Entities

The relation between entities is listed as follows:

- A user may has many appointments
- A user may has many invitations
- An appointment only has a user as a creator
- An appointment may has many invitations
- An invitation has a user and an appointment

To prevent a high degree of coupling, relations between entities is listed as:

- User entity, being the parent of all entities, to not have any relations with other entities.
- Appointment entity, being a child of user, is allowed to have a 'belong to' association and may use user's conversion and controller functions if necessary.
- Invitation entity may have 'belong to' association with both user and appointment. Invitation also allowed to use user's and appointment's conversion and controller functions if necessary.

#### Cascade Deletion

Cascade deletion for entities that has a child relation with another entities is achieved by the **creation of child entities deletion methods by their parent's foreign key**. Example is given below.

```go
// controller.go
func DeleteByAppointment(appointmentID int) {
 deleteByAppointmentDB(appointmentID)
}

// database.go
func deleteByAppointmentDB(appointmentID int) {
 db := database.GetDB()
 db.Delete(&Invitation{}, "appointment_id = ?", appointmentID)
}
```

Later, these deletion methods can be passed as an arguments for entities that has child entities.

```go
// routes/appointment.go
func deleteAppointment(ctx *gin.Context) {
// ... //
 err = appointment.Delete(issuer, invitation.DeleteByAppointment, appointmentToBeDeleted)
// ... //
}

// controller.go
func Delete(issuer string, invitationDeletion func(int), appointment Appointment) error {
 check := readDB(appointment.ID)
 if issuer != check.User.Username {
  return errors.New("you are not the creator of this appointment")
 }
 invitationDeletion(appointment.ID)
 deleteDB(appointment.ID)
 return nil
}
```

## API Documentation

### Authentication

This application use JWT (JSON Web Token) that contains username as issuer and valid for 1 hour long. The token may be placed in header as bearer token.

### Subroutes

#### `/auth`

- `/register`

  - Method: POST
  - URI: `/auth/register`
  - Auth: Not Required
  - Request body:

    ```json
    {
        "name":string,
        "username":string,
        "timezone":int,
        "password":string
    }
    ```

  - Response

    ```json
    {
        "token":string,
        "user":
        {
            "ID":int,
            "name":string,
            "username":string,
            "timezone":int
        }
    }
    ```

- `/login`

  - Method: GET
  - URI: `/auth/login`
  - Auth: Not Required
  - Request body:

    ```json
    {
        "username":string,
        "password":string
    }
    ```

  - Response

    ```json
    {
        "token":string
    }
    ```

#### `/user`

- `/update`

  - Method: PUT
  - URI: `/user/update`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "name":(optional)string,
        "username":(optional)string,
        "timezone":(optional)int,
        "password":(optional)string
    }
    ```

  - Response

    ```json
    {
        "token":string,
        "user":
        {
            "ID":int,
            "name":string,
            "username":string,
            "timezone":int
        }
    }
    ```

- `/delete`

  - Method: DELETE
  - URI: `/user/delete`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    {
        "message":"user deleted"
    }
    ```

#### `/appointment`

- `/create`

  - Method: POST
  - URI: `/appointment/create`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "title":string,
        "start":string(serialized version of go time package),
        "end":string(serialized version of go time package)
    }
    ```

  - Response

    ```json
    {
        "ID":int,
        "title":string,
        "start":string(serialized version of go time package),
        "end":string(serialized version of go time package),
        "creatorID":int
    }
    ```

- `/{id}`

  - Method: GET
  - URI: `/appointment/{id}`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    {
        "ID":int,
        "title":string,
        "start":string(serialized version of go time package),
        "end":string(serialized version of go time package),
        "creatorID":int
    }
    ```

- `/{id}/members`

  - Method: GET
  - URI: `/appointment/{id}/members`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    {
        "appointment":
        {
            "ID":int,
            "title":string,
            "start":string(serialized version of go time package),
            "end":string(serialized version of go time package),
            "creatorID":int
        },
        "creator":string,
        "accepted":
        [
            string
        ],
        "invited":
        [
            string
        ]
    }
    ```

- `/update`

  - Method: PUT
  - URI: `/appointment/update`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "ID":int,
        "title":(optional)string,
        "start":(optional)string(serialized version of go time package),
        "end":(optional)string(serialized version of go time package),
    }
    ```

  - Response

    ```json
    {
        "ID":int,
        "title":string,
        "start":string(serialized version of go time package),
        "end":string(serialized version of go time package),
        "creatorID":int
    }
    ```

- `/delete`

  - Method: DELETE
  - URI: `/appointment/delete`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "ID":int
    }
    ```

  - Response

    ```json
    {
        "message":"appointment deleted"
    }
    ```

#### `/appointments`

- `/created`

  - Method: GET
  - URI: `/appointments/created`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    [
        appointment(see previous documentation)
    ]
    ```

- `/invited`

  - Method: GET
  - URI: `/appointments/invited`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    [
        {
            "ID":int,
            "message":string,
            "appointment_id":int,
            "appointment":appointment(see previous documentation),
            "invitee_id":int,
            "accepted":bool
        }
    ]
    ```

#### `/invitation`

- `/create`

  - Method: POST
  - URI: `/invitation/create`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "message":string,
        "appointment_id":int,
        "invitee_id":int
    }
    ```

  - Response

    ```json
    {
        "ID":int,
        "message":string,
        "appointment_id":int,
        "appointment":appointment(see previous documentation),
        "invitee_id":int,
        "accepted":bool,
        "invited_user":user(see previous documentation)
    }
    ```

- `/{id}`

  - Method: GET
  - URI: `/invitation/{id}`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    {
        "ID":int,
        "message":string,
        "appointment_id":int,
        "appointment":appointment(see previous documentation),
        "invitee_id":int,
        "accepted":bool,
        "invited_user":user(see previous documentation)
    }
    ```

- `/{id}/accept`

  - Method: POST
  - URI: `/invitation/{id}/accept`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    {
        "ID":int,
        "message":string,
        "appointment_id":int,
        "appointment":appointment(see previous documentation),
        "invitee_id":int,
        "accepted":bool
    }
    ```

- `/update`

  - Method: PUT
  - URI: `/invitation/update`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "ID":int,
        "Message":string
    }
    ```

  - Response

    ```json
    {
        "ID":int,
        "message":string,
        "appointment_id":int,
        "appointment":appointment(see previous documentation),
        "invitee_id":int,
        "accepted":bool
    }
    ```

- `/delete`

  - Method: DELETE
  - URI: `/invitation/delete`
  - Auth: **Required**
  - Request body:

    ```json
    {
        "ID":int
    }
    ```

  - Response

    ```json
    {
        "message":"invitation deleted"
    }
    ```

#### `/invitations`

- `/created`

  - Method: GET
  - URI: `/invitations/created`
  - Auth: **Required**
  - Request body:

    ```json
    -
    ```

  - Response

    ```json
    [
        {
            "ID":int,
            "message":string,
            "appointment_id":int,
            "appointment":appointment(see previous documentation),
            "invitee_id":int,
            "accepted":bool,
            "invited_user":user(see previous documentation)
        }
    ]
    ```
