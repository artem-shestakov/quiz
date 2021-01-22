from django.contrib import admin
from .models import Subject, Question, Test


@admin.register(Subject)
class SubjectAdmin(admin.ModelAdmin):
    list_display = ('title',)


@admin.register(Question)
class QuestionAdmin(admin.ModelAdmin):
    list_display = ('text', 'subject', 'created')
    list_filter = ('subject', 'created')


@admin.register(Test)
class TestAdmin(admin.ModelAdmin):
    list_display = ('user', 'status')

