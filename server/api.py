import os
from flask import Flask, jsonify
from flask_cors import CORS
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.dialects.postgresql import JSON


USER = os.getenv('USER')
PASSWORD = os.getenv('PASSWORD')
DB = os.getenv('DB')
HOST = os.getenv('HOST')

app = Flask(__name__)
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://' + USER + ':' + PASSWORD + '@' + HOST + '/' + DB
db = SQLAlchemy(app)

CORS(app)


class Event(db.Model):
    __tablename__ = 'events'

    id = db.Column(db.Integer, primary_key=True)
    info = db.Column(JSON)
    

    def __init__(self, info):
        self.info = info

    def __repr__(self):
        return '<id {}>'.format(self.id)
    
    def serialize(self):
        return {
            'id': self.id, 
            'info': self.info,
        }


@app.route('/')
def index():
    return "", 200

@app.route('/api/')
def api():
    # This is a dummy list, 2 nested arrays containing some
    # params and values
    list = [
        {'name': "phone", 'location': "table"},
        {'name': "laptop", 'location': "rucksack"},
        {'name': "iPad", 'location': "home"},
    ]
    # jsonify will do for us all the work, returning the
    # previous data structure in JSON
    return jsonify(results=list)


@app.route('/api/events/')
def list_events():
    events = Event.query.all()
    return jsonify(results=[e.serialize() for e in events])


if __name__ == "__main__":
     app.run(
        host='0.0.0.0',
        port=5000,
        debug=True,
        threaded=True)


# docker run -e POSTGRES_PASSWORD=mysecretpassword -d postgres
