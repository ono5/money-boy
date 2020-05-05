from django.urls import path
from . import views

app_name = 'account'

urlpatterns = [
    # post views
    path('login/', views.login_view, name='login')
]