from django.db import models


class Release(models.Model):
    name = models.CharField(max_length=255)
    namespace = models.CharField(max_length=255)
    version = models.IntegerField()
    first_deploy = models.DateTimeField()
    last_deploy = models.DateTimeField()
