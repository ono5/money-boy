import pytest
from django.test import TestCase
from ..forms import LoginForm


class LoginFormTest(TestCase):
    """Login Form Test
    """

    def test_login_form_is_valid(self):
        """Test Login Form is valid
        """
        # Arrange ---
        data = {
            'username': 'admin',
            'password': 'XXXXXXXXXXX',
        }

        # Act ---
        login_form = LoginForm(data)

        # Assert ---
        assert login_form.is_valid()

    def test_login_form_is_invalid_without_username(self):
        """Test Login Form is invalid without username
        """
        # Arrange ---
        data = {
            'username': '',
            'password': 'XXXXXXXXXXX',
        }

        # Act ---
        login_form = LoginForm(data)

        # Assert ---
        assert not login_form.is_valid()

    def test_login_form_is_invalid_without_password(self):
        """Test Login Form is invalid without password
        """
        # Arrange ---
        data = {
            'username': 'admin',
            'password': '',
        }

        # Act ---
        login_form = LoginForm(data)

        # Assert ---
        assert not login_form.is_valid()
