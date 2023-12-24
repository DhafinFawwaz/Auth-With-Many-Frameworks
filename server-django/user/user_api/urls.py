from django.contrib import admin
from django.urls import path
from .views import TodoListApiView, RegisterApiView, LoginApiView, AuthApiView

urlpatterns = [
    path('todo/', TodoListApiView.as_view()),
    path('auth/register/', RegisterApiView.as_view()),
    path('auth/login/', LoginApiView.as_view()),
    path('auth/', AuthApiView.as_view()),
]
