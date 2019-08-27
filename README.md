# sig-webserver

The webserver sits in front of the API and the UI. It reverse proxies http traffic to the other servers based on context path.

```
                 |          | =====>  api 
    browser      |    ws    |
                 |          | =====>  ui

```


