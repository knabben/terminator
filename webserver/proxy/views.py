from django.shortcuts import render
from django.views.generic import TemplateView

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

class IndexView(TemplateView):
    template_name = 'index.html'
