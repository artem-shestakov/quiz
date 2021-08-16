# Generated by Django 3.2.6 on 2021-08-16 14:55

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('quizapp', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='question',
            name='type',
            field=models.CharField(choices=[('single_chioce', 'Single Choice'), ('miltiple_choice', 'Multiple Choice')], max_length=15, verbose_name='The type of the question'),
        ),
    ]
