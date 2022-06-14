# Gif App

See changelog [here](https://github.com/luisnquin/gif-app/blob/master/CHANGELOG.md)

## Steps to run:

-   Clone the project to your machine

    ```
     $ git clone https://github.com/luisnquin/gif-app
    ```

-   Generate the public and private keys

    ```
     $ ssh-keygen -t rsa -b 4096 -m PEM -f private.rsa.key
     $ openssl rsa -in private.rsa.key -pubout -outform PEM -out public.rsa.key
    ```

-   Install project dependencies

    ```
     $ go mod tidy
    ```

-   Initialize the database with cache(required to run the application)

    ```
     $ make store
    ```

-   Build and run

    ```
     $ make build && make run
    ```

## Docs

To see the docs, just type:

```console
     $ make docs
```
