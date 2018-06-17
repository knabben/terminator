run:
	cd terminator; yarn start &
	python terminator/manage.py runserver 0.0.0.0:8000
