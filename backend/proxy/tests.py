from django.test import TestCase, Client
from django.urls import reverse

from .request import fetch_list_release, get_tiller_settings, protobuf_to_model, \
extract_attributes, delete_release

from unittest.mock import MagicMock
from unittest import mock

from rest_framework.test import APITestCase
from datetime import datetime


class TestItems(TestCase):
    """ Mocking GRPC engine """

    @mock.patch("grpc.insecure_channel")
    def test_fetch_items(self, inc):
        output = fetch_list_release()
        assert len(inc.mock_calls) == 12

    @mock.patch("grpc.insecure_channel")
    def test_delete_items(self, inc):
        output = delete_release("fake_release")
        assert len(inc.mock_calls) == 11

    @mock.patch("grpc.insecure_channel")
    def test_host_from_settings(self, inc):
        tiller_host = get_tiller_settings()
        output = fetch_list_release()
        assert tiller_host in str(inc.mock_calls[0])

    def test_extract_attributes_complete(self):
        mk = MagicMock()
        mk.info.first_deployed = 1485791240
        mk.data.data1.data2 = "10"
        mk.plata = 40.0

        attrs = [
            ("info.first_deployed", "f", datetime.fromtimestamp),
            ("data.data1.data2", "k", int),
            ("plata", "h", None)
        ]
        output = extract_attributes([mk, mk, mk], attrs)
        assert output[0]['f'] == datetime(2017, 1, 30, 15, 47, 20)
        assert output[1]['k'] == 10
        assert output[2]['h'] == 40.0


class TestEndpoints(TestCase):

    def setUp(self):
        self.client = Client()

    def empty_list():
        return []

    def raise_exception():
        raise Exception

    def ensure_ok_empty(self, response):
        assert response.status_code == 404
        assert response.json() == {"releases": []}

    @mock.patch("proxy.views.fetch_list_release", empty_list)
    def test_endpoint_empty_list(self):
        response = self.client.get(reverse('items'))
        self.ensure_ok_empty(response)

    @mock.patch("proxy.views.fetch_list_release", raise_exception)
    def test_endpoint_exception(self):
        response = self.client.get(reverse('items'))
        self.ensure_ok_empty(response)

    @mock.patch("proxy.views.delete_release", lambda x: x)
    def test_delete_endpoint(self):
        response = self.client.delete('/items/?release_name=some-name')
        assert response.status_code == 200
