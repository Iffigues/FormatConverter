from flask import Flask
from magika import Magika
from pathlib import Path
import os
import json

app = Flask(__name__)

magika = Magika()

def get_dir_content(a):
    e = []
    path = "./"

    try:
        dir_contents = os.listdir(path + a)
    except Exception as e:
        raise CustomError(e)

    for item in dir_contents:
        try:
            e.append(path + a + "/" + item)
        except Exception as e:
            raise CustomError(e)
    return e

def get_type(a):
    t = []
    try:
        e = get_dir_content(a)
    except Exception as e:
        raise CustomError(e)

    for i in e:
        try:
            result =  magika.identify_path(Path(i))
        except Exception as e:
           raise CustomError(e)
        
        z = {
                "path":result.path,
                "ct_label":result.output.ct_label
        }
        t.append(z)
    return t

@app.route(r'/path/<path:path>')
def hello_world(path):
    try:
        response = get_type(path)
    except Exception as e:
        return "error", 500
    return json.dumps(response), 200, {'Content-Type': 'application/json'}

# Run the app if executed directly
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)

