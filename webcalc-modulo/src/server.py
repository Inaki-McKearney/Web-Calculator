from flask import Flask
from flask import request
from flask import Response
from flask import jsonify
# from flask_cors import CORS

import json

# def create_app(test_config=None):
app = Flask(__name__)
# CORS(app)


def bad_request_error(message):
    error = json.dumps({'error':True,'message':message})
    return Response(response=error, status=400, mimetype="application/json")


@app.route('/modulo', methods=['GET'])
def modulo():
    try:
        x = int(request.args.get('x'))
    except TypeError:
        return bad_request_error("x parameter not provided")
    except ValueError:
        return bad_request_error("x parameter is not a valid integer")

    try:
        y = int(request.args.get('y'))
    except TypeError:
        return bad_request_error("y parameter not provided")
    except ValueError:
        return bad_request_error("y parameter is not a valid integer")
    if y == 0:
        return bad_request_error("Divisor (y) must not be 0")

    args = request.args.to_dict()
    del args['x']
    del args['y']
    if len(args) > 0:
        return bad_request_error("Unexpected parameter(s) provided")

    answer = x % y
    
    response = {
        'error': False,
        'string': f'{x} % {y} = {answer}',
        'answer': answer
    }

    return jsonify(response)

    # return app


# if __name__ == '__main__':
#     app.run(debug=True,host='0.0.0.0')
