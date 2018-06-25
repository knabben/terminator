from django.shortcuts import render
from django.views.generic import TemplateView

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

from kubernetes import client, config


class CRDView(APIView):

    def load_config(self):
        try:
            config.load_kube_config()
        except FileNotFoundError:
            config.load_incluster_config()

    def client_request(self, body):
        cli = client.ApiClient()

        api_instance = client.CustomObjectsApi()
        group = 'app.terminator.dev'
        version = 'v1alpha1'

        api_response = api_instance.patch_namespaced_custom_object(
            group, version, 'default', 'terminators', 'deploy', body)

        return api_response

    def post(self, request, format=None):
        self.load_config()

        name = request.data.get("name")
        response = self.client_request({"spec": {name: True}})

        return Response(data=response)

    def delete(self, request, format=None):
        self.load_config()

        name = request.data.get("name")
        response = self.client_request({"spec": {name: False}})

        return Response(data=response)
