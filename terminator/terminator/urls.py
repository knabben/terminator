from django.conf.urls import url, include
from django.contrib import admin

from quickstart import views
from rest_framework import routers


router = routers.DefaultRouter()
router.register(r'users', views.UserViewSet)
router.register(r'groups', views.GroupViewSet)

urlpatterns = [
    url('api/^', include(router.urls)),
    url('^$', views.BlaView.as_view()),
    url('^api-auth/', include('rest_framework.urls', namespace='rest_framework')),
    url(r'^admin/', admin.site.urls),
]
