# Goodm

<!-- TABLE OF CONTENTS -->

  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#connection">Connection</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

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
        "github.com/angrypufferfish/goodm/src/connection"
    )

    func Connect() {

        var goodm connection.Goodm
        ctx := context.Background()

        client, err := goodm.Connect("mongodb://localhost:27017", 10000)

        ///!
        defer goodm.Disconnect()

        if err != nil {
            panic(err)
        }
        client.UseDatabase("test", &ctx)
    }

   ```

<!-- USAGE EXAMPLES -->
## Usage

   ```go
    import (
        "context"
        "github.com/angrypufferfish/goodm/src/connection"
    )

    func Connect() {

        var goodm connection.Goodm
        ctx := context.Background()

        client, err := goodm.Connect("mongodb://localhost:27017", 10000)

        ///!
        defer goodm.Disconnect()

        if err != nil {
            panic(err)
        }
        client.UseDatabase("test", &ctx)
    }

   ```

<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue.


<!-- LICENSE -->
## License

Distributed under the APACHE LICENSE 2.0 See `LICENSE.txt` for more information.