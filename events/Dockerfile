FROM python:3.6.1-alpine
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

RUN apk update \
  && apk add --virtual build-deps gcc python3-dev musl-dev \
  && apk add postgresql-dev \
  && pip install psycopg2 flask flask-cors Flask-SQLAlchemy \
  && apk del build-deps

COPY . /src
WORKDIR /src

CMD ["python", "./api.py"]