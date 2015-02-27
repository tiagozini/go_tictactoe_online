# Welcome to Go Tic-Tac-Toe Online game

## Getting Started
    Install Golang in your system if you haven't done it yet, and run
    the command below:
        go get github.com/tiagozini/go_tictactoe_online

### Run the web server that provide the game:

    revel run go_tictactoe_online

    Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"Welcome to Tic-Tac-Toe Combate"

### Description of Contents

The default directory structure of the game:

    myapp               App root
      app               App sources
        controllers     App controllers
          app.go        Interceptor registration
          reflesh.go    Interceptor for gameroom interactions
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      gameroom          The game go clases
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.

gameroom

    The gameroom classes directory contains all entities of this game and the business logic of them.
    Each entity has a separated class file and main entity is the class Game.

messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### This application use Golang and Revel. To know more about these technologies
consult the links below:

* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](http://revel.github.io/docs/godoc/index.html).
