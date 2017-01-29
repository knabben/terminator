from django.shortcuts import render
from django.views.generic import TemplateView

from rest_framework.views import APIView
from rest_framework.response import Response


class IndexView(TemplateView):
    template_name = 'index.html'


class ListItem(APIView):

    def get(self, request):
        return Response({
            "data": "hello",
            "more": "data"
        })
