from django.db import models
from django.utils import timezone
from django.contrib.auth.models import User


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
    TYPE = (
        ('simple', 'Simple'),
        ('multi_choice', 'Multi choice')
    )

    STATUS = (
        ('draft', 'Draft'),
        ('published', 'Published')
    )

    text = models.CharField(max_length=250, unique=True)
    type = models.CharField(max_length=15, choices=TYPE, default='simple')
    answer = models.CharField(max_length=1250, default='')
    options = models.CharField(max_length=1250, default='')
    created = models.DateTimeField(default=timezone.now)
    subject = models.ForeignKey(Subject, on_delete=models.CASCADE, related_name='subject')

    def __str__(self):
        return self.text


class Test(models.Model):
    """
    Test with questions
    """
    STATUS = (
        ('draft', 'Draft'),
        ('ready', 'Ready'),
        ('active', 'Active'),
        ('finished', 'Finished')
    )
    user = models.ForeignKey(User, on_delete=models.CASCADE, related_name='tests')
    status = models.CharField(max_length=10, choices=STATUS, default='draft')
    questions = models.ManyToManyField(Question)

    def __str__(self):
        return f'{self.user}\' test'
