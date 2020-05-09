from django.conf import settings
from django.contrib.auth import authenticate, login
from django.shortcuts import redirect, render
from .forms import LoginForm


def login_view(request):
    form = LoginForm()

    if request.method == 'POST':
        form = LoginForm(request.POST)
        if form.is_valid():
            cd = form.cleaned_data
            user = authenticate(request,
                                username=cd['username'],
                                password=cd['password'])
            if user is not None:
                if user.is_active:
                    login(request, user)
                    return redirect(settings.LOGIN_REDIRECT_URL)
                else:
                    return render(request, 'account/login.html', {'errors': 'error', 'form': form})
            else:
                return render(request, 'account/login.html', {'errors': 'error', 'form': form})

    return render(request, 'account/login.html', {'form': form})
