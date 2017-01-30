import grpc

from hapi.services import tiller_pb2_grpc, tiller_pb2

from collections import defaultdict
from proxy.models import Release

from datetime import datetime


def get_tiller_settings():
    from django.conf import settings
    return settings.TILLER_HOST

def fetch_list_release():
    channel = grpc.insecure_channel(get_tiller_settings())
    stub = tiller_pb2_grpc.ReleaseServiceStub(channel)
    response = stub.ListReleases( tiller_pb2.ListReleasesRequest())
    return response.next()

def extract_attributes(releases, attrs):
    """ Iterate through Releases GRPC models and extracts
        attributes, applying the function after."""
    output = []
    for release in releases:
        tmp_rel = defaultdict(dict)
        for key, key_out, func in attrs:
            try:
                if '.' not in key:
                    tmp_attr = getattr(release, key)
                else:
                    tmp_attr = None
                    for n in key.split('.'):
                        tmp_attr = getattr(release, n) if not tmp_attr else \
                                   getattr(tmp_attr, n)

                # Apply function if exists
                tmp_rel[key_out] = tmp_attr if func is None else func(tmp_attr)

            except AttributeError:
                continue
        output.append(tmp_rel)
    return output


def protobuf_to_model(protobuf):
    """ Move ProtoBuffer to models """
    # TODO - Mock protobuf responses to create test
    #assert type(protobuf) == tiller_pb2.ListReleasesResponse
    attrs = [
        ("info.first_deployed.seconds", "first_deploy",
         lambda x: datetime.fromtimestamp(float(x))),
        ("info.last_deployed.seconds", "last_deploy",
         lambda x: datetime.fromtimestamp(float(x))),
        ("name", "name", None),
        ("namespace", "namespace", None),
        ("version", "version", None),
    ]

    protos = []
    for proto in extract_attributes(protobuf.releases, attrs):
        release, created = Release.objects.get_or_create(**proto)
        protos.append(proto)
    return protos

