import proxy.routing
from channels.routing import ProtocolTypeRouter, URLRouter

application = ProtocolTypeRouter({
    'websocket': URLRouter(proxy.routing.websocket_urlpatterns)
})
