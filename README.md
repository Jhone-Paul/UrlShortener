# UrlShortener
A http handler that will look at the path of any incoming web request and determine if it should redirect the user.

For example if we had a redirect set up for ```/git``` to ```https://github.com/Jhone-Paul``` the code looks for any incoming web requests with ```/git``` and then redirects.
