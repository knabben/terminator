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

    def post(self, request, format=None):
        self.load_config()

        svc_name = request.data.get("name")
        cli = client.ApiClient()

        api_instance = client.CustomObjectsApi()
        group = 'app.terminator.dev'
        version = 'v1alpha1'

        body = {
            "spec": {
                svc_name: True,
            }
        }

        api_response = api_instance.patch_namespaced_custom_object(
            group, version, 'default', 'terminators', 'deploy', body)

        return Response(data=api_response)

    def delete(self, request, format=None):
        self.load_config()

        svc_name = request.data.get("name")
        cli = client.ApiClient()

        api_instance = client.CustomObjectsApi()
        group = 'app.terminator.dev'
        version = 'v1alpha1'

        body = {
            "spec": {
                svc_name: False,
            }
        }

        api_response = api_instance.patch_namespaced_custom_object(
            group, version, 'default', 'terminators', 'deploy', body)

        return Response(data=api_response)
