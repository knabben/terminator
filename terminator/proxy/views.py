from django.shortcuts import render
from django.views.generic import TemplateView

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

from .request import fetch_list_release, delete_release, protobuf_to_model


class IndexView(TemplateView):
    template_name = 'index.html'


class ListItem(APIView):

    def get(self, request):
        """ Fetch all chars from Tiller """
        try:
            grpc_status = status.HTTP_200_OK
            release = fetch_list_release()
            data = protobuf_to_model(release)
        except Exception as e:
            grpc_status, data = status.HTTP_404_NOT_FOUND, []
        return Response({"releases": data}, status=grpc_status)

    def delete(self, request):
        """ Delete a specific release by name """
        try:
            grpc_status = status.HTTP_200_OK
            release_name = request.GET['release_name']
            if release_name is None:
                raise ValueError
            delete_release(release_name)
        except Exception:
            grpc_status = status.HTTP_404_NOT_FOUND
        return Response(status=grpc_status)
