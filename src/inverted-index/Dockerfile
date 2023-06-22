FROM python:3.9.2-buster

RUN apt-get update && apt-get upgrade -y
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
RUN apt-get install -y nodejs 
RUN npm install -g npm@latest
RUN npm install -g nodemon

WORKDIR /usr/app
COPY requirements.txt requirements.txt
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 5000

COPY . .

CMD ["nodemon"]