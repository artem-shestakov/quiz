from quizapp.models import Question
from django.shortcuts import render
from django.http import HttpResponse


def welcome_view(request):
    return render(request, 'index.html')

def questions(request):
    if request.method == "GET":
        questions = Question.objects.all()
        questions_list = []
        for question in questions:
            answers = [answer for answer in question.answer_set.all()]
            questions_list.append({
                'question': question,
                'answers': answers
            })
        context = {
            'questions': questions_list
        }
        return render(request, 'questions.html', context)
    else:
        return HttpResponse("POST")
