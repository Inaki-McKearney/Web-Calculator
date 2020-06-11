from urllib.parse import urlencode
# from src import server.app
import json

def call(client, path, params):
    url = path + '?' + urlencode(params)
    response = client.get(url)
    return json.loads(response.data.decode('utf-8'))

def test_one(client):
    result = call(client, '/modulo', {'x': 3, 'y': 2})
    assert result['error'] == False
    assert result['string'] == "3 % 2 = 1"
    assert result['answer'] == 1

def test_two(client):
    result = call(client, '/modulo', {'x': 26382, 'y': 2723})
    assert result['error'] == False
    assert result['string'] == "26382 % 2723 = 1875"
    assert result['answer'] == 1875

def test_three(client):
    result = call(client, '/modulo', {'x': -50, 'y': 6})
    assert result['error'] == False
    assert result['string'] == "-50 % 6 = 4"
    assert result['answer'] == 4

def test_four(client):
    result = call(client, '/modulo', {'x': 10, 'y': -4})
    assert result['error'] == False
    assert result['string'] == "10 % -4 = -2"
    assert result['answer'] == -2

def test_five(client):
    result = call(client, '/modulo', {'x': 2, 'y': 5})
    assert result['error'] == False
    assert result['string'] == "2 % 5 = 2"
    assert result['answer'] == 2

def test_six(client):
    result = call(client, '/modulo', {'x': 3, 'y': 3})
    assert result['error'] == False
    assert result['string'] == "3 % 3 = 0"
    assert result['answer'] == 0


def test_error_one(client):
    result = call(client, '/modulo', {'y': 5})
    assert result['error'] == True
    assert result['message'] == "x parameter not provided"

def test_error_two(client):
    result = call(client, '/modulo', {'x': 'm', 'y': 5})
    assert result['error'] == True
    assert result['message'] == "x parameter is not a valid integer"

def test_error_three(client):
    result = call(client, '/modulo', {'x': 5})
    assert result['error'] == True
    assert result['message'] == "y parameter not provided"

def test_error_four(client):
    result = call(client, '/modulo', {'x': 5, 'y': 'n'})
    assert result['error'] == True
    assert result['message'] == "y parameter is not a valid integer"

def test_error_five(client):
    result = call(client, '/modulo', {'x': 5, 'y':0})
    assert result['error'] == True
    assert result['message'] == "Divisor (y) must not be 0"

def test_error_six(client):
    result = call(client, '/modulo', {'x': 5, 'y': 10, 'z': 3})
    assert result['error'] == True
    assert result['message'] == "Unexpected parameter(s) provided"
