import json
import time

from datetime import datetime

from asgiref.sync import async_to_sync
from channels.generic.websocket import WebsocketConsumer


class Consumer(WebsocketConsumer):

    def connect(self):
        print("CONNECT", self)
        print(self.channel_layer)
        async_to_sync(self.channel_layer.group_add)('data', self.channel_name)
        self.accept()

    def disconnect(self, close_code):
        print("DISCONNECT")
        async_to_sync(self.channel_layer.group_discard)('data', self.channel_name)

    def receive(self, text_data):
        print("receive", datetime.now())
        print(text_data)

        async_to_sync(self.channel_layer.group_send)(
            'data',
            {
                'type': 'chat',
                'message': text_data
            }
        )

    def chat(self, event):
        message = event['message']

        print("send to websocket", datetime.now())
        self.send(text_data=json.dumps({'message': message}))
