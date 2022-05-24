# beat-challenge

This challenge was really a challenge, it is my fourth project in golang and the truth is that I really enjoyed it.

I had to do a lot of research and change the paradigm that I had.

During the middle of the process I wanted to switch it to java, to be honest, it wouldn't have taken that long. However, I am very happy with what I learned and with what I have achieved.

I know that it can be improved and of course, there may be a lack of good practices, but I firmly promise to learn more about golang.

## A little beginning

>If at first you don't succeed, call it version 1.0 __~ Erick Barrera__

### Author
* Erick Barrera - **ebarreral.isc@gmail.com**

## Table of contents

* [Technology stack](#technology-stack)
* [Requirements](#requirements)
* [How to run the application](#how-to-run-the-application)
    - [With Makefile](#with-makefile)
    - [With docker](#with-docker)
    - [With docker-compose](#with-docker-compose)
* [How to run the tests](#how-to-run-the-test)

## Technology stack

- Go
- Docker

## Requirements

- Operating System based on UNIX
- Docker

## How to run the application

To run this project you have 3 options

1. The client that is generated with the Makefile
2. With docker
3. With docker-compose

### With Makefile

For this option, you need to have to __*golang*__ installed on your __*machine*__.

Then you must execute the following steps:

1. Download the dependencies

```shell
$ go mod download
$ go mod tidy
```

2. Create the client

```shell
$ make cli
```

3. Use the cli to process the file

```shell
$ ./bin/fare-cli -f ./input/paths.csv -o my-result.csv
```

4. In the folder named "output" you will see the file that was generated with the name that has been passed as the second parameter with the -o argument

```
├── output
│   └── my-result.csv
```

### With docker

For this build, the image (Dockerfile) has mounted internal volumes so that files can be read between your local machine and the container.

To generate the image you must execute the following commands:

1. Create the image

```shell
$ docker build -t fare-cli .
```

2. Run the image

```shell
$ docker run -v $(pwd)/input/:/app/input -v $(pwd)/output:/app/output -e input=paths.csv -e output=another-result.csv fare-cli
```

3. The *$(pwd)* command will be in charge of linking your current path to the container volume and the *-e* parameters will be the name of the files

```
$(pwd)
├── input
│   └── paths.csv
├── output
│   └── another-result.csv
```

4. In the folder named "output" you will see the file that was generated with the name that has been passed as the second parameter with the -o argument

```
├── output
│   └── another-result.csv
```

### With docker-compose

Practically it is the automation of the process with docker, here the only thing we will use will be the docker-compose up command.

```shell
$ docker-compose up
```

## How to run the tests

If you have golang installed on your machine you can use the Makefile file.

```shell
$ make test
```

If you have used the docker options, both internally already run the tests, if the tests fail, the image cannot be built.

```
/* Dockerfile */
VOLUME ${WORK_DIR}/input
VOLUME ${WORK_DIR}/output

RUN make cli

/* Make file */
cli:
	$(MAKE) test
	$(GO) build -o $(CLI_DIR) $(MAIN)

test:
	$(GO) test -v $(TEST_DIR_RECURSIVELY) --parallel 10
```