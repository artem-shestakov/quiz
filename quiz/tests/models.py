from django.db import models


class Subject(models.Model):
    """
    Quiz subject
    """
    title = models.CharField(max_length=250)
    created = models.DateTimeField(auto_now_add=True)

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
    created = models.DateTimeField(auto_now_add=True)
    Subject = models.ForeignKey(Subject, on_delete=models.CASCADE, related_name='questions')

    def __str__(self):
        return self.text


class Option(models.Model):
    """
    Answer option of question
    """
    text = models.CharField(max_length=250)
    created = models.DateTimeField(auto_now_add=True)
    question = models.ForeignKey(Question, on_delete=models.CASCADE, related_name='test_question')

    def __str__(self):
        return self.text
