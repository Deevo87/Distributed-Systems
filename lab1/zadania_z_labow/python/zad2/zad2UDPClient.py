import socket;

serverIP = "127.0.0.1"
serverPort = 9008
msg = "Żółta gęś"
data = msg.encode('utf-8')

print('PYTHON UDP CLIENT')
client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
client.sendto(bytes(msg, 'utf-8'), (serverIP, serverPort))