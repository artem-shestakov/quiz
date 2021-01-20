from django.contrib import admin
from .models import Subject, Question, Option


@admin.register(Subject)
class SubjectAdmin(admin.ModelAdmin):
    list_display = ('title',)


@admin.register(Question)
class QuestionAdmin(admin.ModelAdmin):
    list_display = ('text', 'answer',)


@admin.register(Option)
class OptionAdmin(admin.ModelAdmin):
    list_display = ('text',)
