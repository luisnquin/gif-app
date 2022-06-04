#!/usr/bin/env python

from models.profile import Profile

from typer import Typer, Option, echo
from datetime import datetime
from typing import List
from faker import Faker


fake: object = Faker()
app: object = Typer()


def create_users_stmts(n: int) -> List[str]:
    users: List[str] = []

    for _ in range(n):
        profile: object = Profile.parse_obj(fake.profile())
        stmt_template: str = "INSERT INTO users(username, firstname, lastname, email, password, birthday) VALUES('%s', '%s', '%s', '%s', '%s', '%s');"

        firstname: str = profile.name.split(" ")[0]
        lastname: str = " ".join(profile.name.split(" ")[1:])

        users.append(stmt_template %
                     (profile.username, firstname, lastname, profile.mail, fake.password(), datetime.combine(profile.birthdate, datetime.min.time()).utcnow()))

    return users


@app.command()
def main(length: int = Option(default=10, show_default=True),
         stdout: bool = Option(default=False, show_default=True),
         output_file: str = Option(default="./tools/automock/etc/mock.sql", )) -> None:

    if stdout:
        for user_stmt in create_users_stmts(length):
            print(user_stmt)

        return

    with open(output_file, mode="w") as file:
        for user_stmt in create_users_stmts(length):
            file.write(user_stmt+"\n")


if __name__ == '__main__':
    app()
