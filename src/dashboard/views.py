from django.shortcuts import render, HttpResponse


def home_view(request):
    """home view
    """
    return HttpResponse('<html>Hello</html>')
