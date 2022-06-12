#!/usr/bin/env python

from config_model import Configuration

from psycopg2 import connect
from typing import List
from json import loads


def load_configuration() -> Configuration:
    with open(file="./config-server.json", mode="r") as file:
        config_obj: dict = loads(file.read())

        config: object = Configuration.parse_obj(obj=config_obj)

        return config


def read_migration_file() -> List[str]:
    with open(file="./schemas.sql", mode="r") as file:
        statements: List[str] = []
        statement: str = ""

        for line in file.readlines():
            line = line.replace("\n", "")
            statement += line

            if line[-1] == ";":
                statements.append(statement)
                statement = ""

        return statements


def main() -> None:
    statements: List[str] = read_migration_file()
    config: object = load_configuration()

    with connect(dsn=config.database.in_local_dsn) as connection:
        with connection.cursor() as cursor:
            cursor.execute("DROP SCHEMA public CASCADE;")
            cursor.execute("CREATE SCHEMA public;")

            for statement in statements:
                cursor.execute(statement)

            connection.commit()


if __name__ == '__main__':
    main()
