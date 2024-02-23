# Use the official Python image as the base image
FROM python:3.9

# Set the working directory in the container
COPY server /server
WORKDIR /server
RUN pip install -r requirements.txt
CMD ["python", "app.py", "run", "-h", "gopiko.fr", "-p", "5000"]

