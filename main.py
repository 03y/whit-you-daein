from flask import Flask, jsonify

import os
import psutil
import platform

api = Flask(__name__)

@api.route('/hello', methods=['GET'])
def hello():
    return 'hello world'

@api.route('/uptime', methods=['GET'])
def uptime():
	output = os.popen('uptime').read()
	return output

@api.route('/mem', methods=['GET'])
def mem():
	output = psutil.virtual_memory().percent
	print(output)
	return str(output)

if __name__ == '__main__':
    api.run(host='0.0.0.0')

