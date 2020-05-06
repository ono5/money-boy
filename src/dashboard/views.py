from django.shortcuts import render, HttpResponse
from django.contrib.auth.decorators import login_required


@login_required
def home_view(request):
    """home view
    """
    return render(request,
                  'dashboard/dashboard.html',
                  {})
