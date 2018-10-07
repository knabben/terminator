from django.conf.urls import url
from . import consumers

websocket_urlpatterns = [
    url('^ws/events/$', consumers.Consumer)
]
