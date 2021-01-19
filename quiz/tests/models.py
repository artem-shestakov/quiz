from django.db import models
from django.utils import timezone
from django.contrib.auth.models import User


class Subject(models.Model):
    title = models.CharField(max_length=250)
    created = models.DateTimeField(auto_now_add=True)

    class Meta:
        ordering = ('-created',)

    def __str__(self):
        return self.title


class Question(models.Model):
    text = models.CharField(max_length=250)
    answer = models.CharField(max_length=100)
    options = models.CharField(max_length=1000)

    def _get_options(self):
        return self.split()

    def _set_options(self, lst: list):
        self.options = ' '.join(lst)

    def _add_options(self, option: str):
        options = self._get_options().append(option)
        self._set_options(options)
