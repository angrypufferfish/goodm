# Goodm

<!-- TABLE OF CONTENTS -->

  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#connection">Connection</a></li>
    <li><a href="#define-document">Define Document</a></li>
    <li><a href="#save-a-document">Save a Document</a></li>
    <li><a href="#save-and-serialize-output">Save and serialize output</a></li>
    <li><a href="#list-documents">List documents</a></li>
    <li><a href="#operators">Operators</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>



<!-- ABOUT THE PROJECT -->
## About The Project

Golang ODM based on the official MongoDB driver.


<!-- GETTING STARTED -->
## Installation

```sh
go get github.com/angrypufferfish/goodm
```

## Connection

```go
import (
  "context"
  "github.com/angrypufferfish/goodm"
)

func Connect() {

  var iGoodm goodm.Goodm
  ctx := context.Background()

  client, err := iGoodm.Connect("mongodb://localhost:27017", 10000)

  ///!
  defer iGoodm.Disconnect()

  if err != nil {
    panic(err)
  }
  client.UseDatabase("test", &ctx)
}

```

## Define Document

```go
package example

import (
  "github.com/angrypufferfish/goodm/src/base"
)

type UserInfo struct {
  Address string `json:"address" bson:"address"`
}

type User struct {
  //Define collection name on goodm tag
  base.BaseDocument `json:",inline" bson:",inline" goodm:"users"`

  LastName  string     `json:"lastName" bson:"lastName"`
  FirstName string     `json:"firstName" bson:"firstName"`
  Info  UserInfo       `json:"info" bson:"info"`
}

```

## Usage

### Save a document

```go
package example

import (
  "github.com/angrypufferfish/goodm"
)

user := &User{
	FirstName: "Mario",
	LastName:  "Rossi",
	Info: UserInfo{
		Address: "via campo",
	},
}

savedUser := goodm.Create[User](user)

```

### Save and serialize output

```go
package example

import (
  "github.com/angrypufferfish/goodm"
)

type UserName struct {
	FirstName string            `json:"firstName" bson:"firstName"`
	LastName string             `json:"lastName" bson:"lastName"`
}

user := &User{
	FirstName: "Mario",
	LastName:  "Rossi",
	Info: UserInfo{
		Address: "via campo",
	},
}

savedUser := goodm.CreateAndSerialize[User, UserName](user)
//output usr
//{
//  "firstName": "Mario",
//  "lastName": "Rossi",
//}
```

### List

```go
package example

import (
  "github.com/angrypufferfish/goodm"
)

type UserName struct {
	FirstName string            `json:"firstName" bson:"firstName"`
	LastName string             `json:"lastName" bson:"lastName"`
}

//Output users: []User
users := goodm.FindAll[User]()

//Output serializedUsers: []UserName
serializedUsers := goodm.FindAllAndSerialize[User, UserName]()

```

### Operators

```go
package example

import (
  "github.com/angrypufferfish/goodm"
  "github.com/angrypufferfish/goodm/src/query"
)


/// {"$and": [
///    {"lastName": {"$ne": "Doe"}},
///    {"firstName": {"$eq": "Mario"}}
///  ]}
users := goodm.Find[User](
  query.And(
    ///{"lastName": {"$ne": "Doe"}}
		query.Ne("lastName", "Doe"),
    ///{"firstName": {"$eq": "Mario"}}
		query.Eq("firstName", "Mario"),
	),
)


```

<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue.


<!-- LICENSE -->
## License

Distributed under the APACHE LICENSE 2.0 See `LICENSE.txt` for more information.