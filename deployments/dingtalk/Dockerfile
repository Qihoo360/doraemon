From  python:3.6.10-alpine3.11
RUN pip install flask && pip install requests && pip install datetime
ADD  doraemontodingtalk.py   /doraemontodingtalk.py
EXPOSE 8100
CMD python /doraemontodingtalk.py
