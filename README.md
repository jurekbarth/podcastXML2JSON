# Podcast XML Feed to JSON API
This Go function exposes a simple function that converts your Podcast XML RSS Feed into a more web friendly JSON format.

## Development
1. Download the Zeit Now cli with `npm i -g now@latest`
2. Run `now dev`
3. Open [`http://localhost:3000/api?feed=https://anchor.fm/s/119f3bc8/podcast/rss`](http://localhost:3000/api?feed=https://anchor.fm/s/119f3bc8/podcast/rss)


## Deployment
In order to use Zeit Now you need to register for an account over at [Zeit](https://zeit.co/). After you registered and logged in with the cli you can run `now deploy` to deploy your function.
