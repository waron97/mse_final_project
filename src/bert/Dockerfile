FROM python:3.9.2-buster

WORKDIR /usr/app
COPY requirements.txt requirements.txt
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 5000

COPY . .

CMD ["python", "-m", "flask", "--app", "main", "run", "--host", "0.0.0.0"]