from django.db import models
from django.contrib.auth.models import AbstractUser

class Mahasiswa(AbstractUser):
    nim = models.CharField(max_length=8, unique=True)

class Todo(models.Model):
    task = models.CharField(max_length = 180)
    timestamp = models.DateTimeField(auto_now_add = True, auto_now = False, blank = True)
    completed = models.BooleanField(default = False, blank = True)
    updated = models.DateTimeField(auto_now = True, blank = True)
    mahasiswa_id = models.ForeignKey(Mahasiswa, on_delete = models.CASCADE, blank = True, null = True)

    def __str__(self):
        return self.task
