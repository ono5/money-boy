from django.shortcuts import render, HttpResponse


def home_view(request):
    """home view
    """
    return render(request,
                  'dashboard/dashboard.html',
                  {})
