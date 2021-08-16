from django.db import models


class Question(models.Model):
    class QuestionType(models.TextChoices):
        SINGLE_CHOICE = "single_chioce"
        MULTIPLE_CHOICE = "miltiple_choice"

    type = models.CharField(
        max_length=15,
        verbose_name="The type of the question",
        choices=QuestionType.choices
    )
    content = models.TextField(
        help_text="The question text"
    )
    created_at = models.DateTimeField(
        auto_now_add=True,
        help_text="The date and time the question was created"
    )
    updated_at = models.DateTimeField(
        auto_now_add=True,
        help_text="The date and time the question was last updated"
    )

    def __str__(self) -> str:
        return self.content


class Answer(models.Model):
    content = models.TextField(
        help_text="Question answer"
    )
    is_correct = models.BooleanField(
        default=False,
        help_text="True if answer correct to the question"
    )
    created_at = models.DateTimeField(
        auto_now_add=True,
        help_text="The date and time the question was created"
    )
    updated_at = models.DateTimeField(
        auto_now_add=True,
        help_text="The date and time the question was last updated"
    )
    question = models.ForeignKey(
        Question,
        on_delete=models.PROTECT,
        help_text="The question that answer is to"
    )

    def __str__(self) -> str:
        return self.content
    