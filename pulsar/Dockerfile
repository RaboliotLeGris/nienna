FROM python:3.9

WORKDIR /pulsar
COPY . .
RUN pip install --no-cache-dir -r requirements.txt

CMD ["python3", "src/pulsar.py"]