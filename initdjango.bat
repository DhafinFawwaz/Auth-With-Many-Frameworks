echo [Initializing Django]
pip install virtualenv
cd server-django
python -m venv venv
venv\Scripts\activate
cd user
pip install -r requirements.txt
python manage.py migrate --run-syncdb
python manage.py makemigration
python manage.py migrate
cd ../../
deactivate
