# Blackhole

We have detected exceedingly high usage of your mock/proxy server, you are rate limited for the time being. Please reach us at support@apiary.io if you need any help restoring access and please help us understand your use case.

## Create new app on Heroku

```
heroku apps:create apiary-blackhole -o apiary
heroku config:set GOVERSION=go1.8
heroku buildpacks:set heroku/go
```

