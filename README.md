# Meow App

## Steps to run:

-   Clone the project to your machine

    ```
    git clone https://github.com/luisnquin/meow-app
    ```

-   Generate the public and private keys
    ```
    ssh-keygen -t rsa -b 4096 -m PEM -f private.rsa.key
    openssl rsa -in private.rsa.key -pubout -outform PEM -out public.rsa.key
    ```
-   Install project dependencies
    ```
    go mod tidy
    ```
-   Initialize the database(in docker)
    ```
    docker compose up -d postgres_db
    ```
-   Build and run
    ```
    make build && make run
    ```
