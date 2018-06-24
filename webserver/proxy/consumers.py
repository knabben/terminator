import json
import time

from datetime import datetime

from asgiref.sync import async_to_sync
from channels.generic.websocket import WebsocketConsumer


class Consumer(WebsocketConsumer):

    def connect(self):
        async_to_sync(self.channel_layer.group_add)('data', self.channel_name)
        self.accept()

    def disconnect(self, close_code):
        async_to_sync(self.channel_layer.group_discard)('data', self.channel_name)

    def receive(self, text_data):
        async_to_sync(self.channel_layer.group_send)(
            'data',
            {
                'type': 'kubernetes',
                'payload': text_data
            }
        )

    def kubernetes(self, event):
        message = event['payload']
        self.send(text_data=json.dumps({'message': message}))
