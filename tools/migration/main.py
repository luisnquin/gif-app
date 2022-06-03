from config_model import Configuration

from psycopg2 import connect
from json import loads


def load_configuration() -> Configuration:
    with open(file="./config-server.json", mode="r") as file:
        config_obj: dict = loads(file.read())

        config: object = Configuration.parse_obj(obj=config_obj)

        return config


def main() -> None:
    config: object = load_configuration()

    with connect(dsn=config.database.in_local_dsn) as connection:
        with connection.cursor() as cursor:
            cursor.execute("SELECT 1;")
            result: List[Tuple] = cursor.fetchall()

            print(result)


if __name__ == '__main__':
    main()
