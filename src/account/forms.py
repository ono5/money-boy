from django import forms


class LoginForm(forms.Form):
    """Login Form to use login page
    """
    username = forms.CharField()
    password = forms.CharField(widget=forms.PasswordInput)
