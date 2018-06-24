from django.conf.urls import url, include
from django.contrib import admin

from proxy import views

urlpatterns = [
    url('^crd$', views.CRDView.as_view()),
]
