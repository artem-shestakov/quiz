from quizapp.models import Answer, Question
from django.contrib import admin

class Question(admin.ModelAdmin):
    admin.site.register(Question)

class Answer(admin.ModelAdmin):
    admin.site.register(Answer)
