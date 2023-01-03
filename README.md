
# Golang-ChatGPT-TelegramBot-Render
# A Golang ChatGPT TelegramBot project, quickly built on the platform Render


### [繁體中文](https://github.com/pyfbsdk59/Golang-ChatGPT-TelegramBot-Render/blob/main/README_tw.md)<br>
### [日本語](https://github.com/pyfbsdk59/Golang-ChatGPT-TelegramBot-Render/blob/main/README_jp.md)


#### 1. This project refers to the following ones and official docs below.
https://github.com/kkdai/chatgpt<br>
https://github.com/go-telegram-bot-api/telegram-bot-api<br>

#### 2. Please fork this project to your own Github account first. Go to Render website, choose to add "Web Services", import this project with your Github account, and then set your own project name and choose a free plan. Remember to click "Advanced" below to set the environment variables.




#### 3. Four environment variables must be set in Render's Environment Variables in this project, namely OPENAI_TOKEN and OPENAI_MAXTOKENS, which are OPENAI's api key and the upper limit of the token number of the answer text (the more you set, the more text of the answer you get, but if you exceed the free quota, you will spend more money. It can be set around 200-400 at the beginning), and then there are two other environment variables CHANNEL_SECRET and CHANNEL_TOKEN of line. It may take some time to start deploying after setting. After success, copy your own URL to the Line developer webpage to set the Webhook URL. For example, 

https://xxx.onrender.com/callback

------
#### Line and OPENAI api settings: 
https://github.com/howarder3/GPT-Linebot-python-flask-on-vercel


