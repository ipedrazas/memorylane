import os
from flask import Flask, jsonify, request, json
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


@app.route('/api/hello/', methods=['POST'])
def hello():
    name=request.form['yourname']
    email=request.form['youremail']
    return ""


@app.route('/api/articles/<articleid>')
def api_article(articleid):
    return 'You are reading ' + articleid


@app.route('/api/messages', methods = ['POST'])
def api_message():
# -> % curl -H 'Content-Type: application/json' -X POST -d '{"name":"Ivan"}' http://192.168.64.30:32679/api/messages
    if request.headers['Content-Type'] == 'application/json':
        # return "JSON Message: " + json.dumps(request.json)
        event = Event(info=request.json)
        db.session.add(event)
        db.session.commit()
        return jsonify(event=event.serialize())
    else:
        return "415 Unsupported Media Type ;)"


if __name__ == "__main__":
     app.run(
        host='0.0.0.0',
        port=5000,
        debug=True,
        threaded=True)


# docker run -e POSTGRES_PASSWORD=mysecretpassword -d postgres
