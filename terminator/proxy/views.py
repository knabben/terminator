from django.shortcuts import render
from django.views.generic import TemplateView

from rest_framework.views import APIView
from rest_framework.response import Response

from .request import fetch_list_release, protobuf_to_model


class IndexView(TemplateView):
    template_name = 'index.html'


class ListItem(APIView):

    def get(self, request):
        try:
            release = fetch_list_release()
            data = protobuf_to_model(release)
        except Exception as e:
            data = []
        return Response({"releases": data})
