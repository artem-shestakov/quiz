from django.db import models
from django.utils import timezone


class Subject(models.Model):
    """
    Quiz subject
    """
    title = models.CharField(max_length=250)
    created = models.DateTimeField(default=timezone.now)

    class Meta:
        ordering = ('-created',)

    def __str__(self):
        return self.title


class Question(models.Model):
    """
    Subject question
    """
    text = models.CharField(max_length=250)
    answer = models.CharField(max_length=100)
    created = models.DateTimeField(default=timezone.now)
    subject = models.ForeignKey(Subject, on_delete=models.CASCADE, related_name='test_subject')

    def __str__(self):
        return self.text


class Option(models.Model):
    """
    Answer option of question
    """
    text = models.CharField(max_length=250)
    created = models.DateTimeField(default=timezone.now)
    question = models.ForeignKey(Question, on_delete=models.CASCADE, related_name='test_question')

    def __str__(self):
        return self.text
