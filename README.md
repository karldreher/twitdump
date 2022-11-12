# Twitdump - The Twitter Dumper
## A utility for downloading your own Twitter content.

`twitdump` is a command-line driven application which can be used to download content from your Twitter account.  Typically, this is used to download images from your own timeline.  

For interactive help with any part of the application, just type the command:  
```
twitdump
```

# Setup
If you don't have one, configure a Twitter Developer account.  This must be associated with the account you want to "dump".

You need to configure the API keys/secrets in a YAML config file as follows:  

```yaml
screenName: my-cool-screenname-no-ampersand
consumerKey: "Consumerkey provided by Twitter Developer account"
consumerSecret: "Secret associated with the key above"
accessToken: "Api access key provided by Twitter Developer account"
accessSecret: "Secret associated with the key above"

```
This can be in a file called `.twitdump.yaml` in your home directory, or in another file you specify with `--config`.  


# Usage

Currently, only the "images" subcommand is supported.  There are several options for this subcommand and it is reccomended you learn about them:
```
twitdump images --help
```

Running this command without the help flag will download the images from your timeline to the directory you're in.  

```
twitdump images
```

