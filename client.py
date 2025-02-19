#!/usr/bin/python3

import socket

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

host = socket.gethostname()
port = 888

client.connect((host, port))

msg = client.recv(1024)

client.close()

print(msg.decode('ascii'))