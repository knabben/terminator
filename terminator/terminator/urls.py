from django.conf.urls import url, include
from django.contrib import admin

from proxy import views


urlpatterns = [
    url('^$', views.IndexView.as_view()),

    # APIs
    url('items/', views.ListItem.as_view()),

    url('^api-auth/', include('rest_framework.urls', namespace='rest_framework')),
    url(r'^admin/', admin.site.urls),
]
