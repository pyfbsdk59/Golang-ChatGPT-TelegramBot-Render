# Golang-ChatGPT-TelegramBot-Render
# 一個Golang ChatGPT TelegramBot專案，快速建置於平台Render。


### [English](https://github.com/pyfbsdk59/Golang-ChatGPT-TelegramBot-Render/blob/main/README.md)
### [日本語](https://github.com/pyfbsdk59/Golang-ChatGPT-TelegramBot-Render/blob/main/README_jp.md)


#### 1. 本專案參考了以下前輩和官方的方案改成製作：
https://github.com/kkdai/chatgpt<br>
https://github.com/go-telegram-bot-api/telegram-bot-api<br>


#### 2. 請先fork本專案到自己的Github帳號裡。然後前往Render網站中設定，選擇新增「Web Services」，可用Github帳號匯入此專案，然後設定自己的名稱和選擇免費free方案。記得按下方「Advanced」，設定環境變數。


#### 3. 必須在Render的Environment Variables設定3個環境變數，分別是OPENAI_TOKEN和OPENAI_MAXTOKENS，分別是OPENAI的api key和回答文字的token數量上限（設越多回答的文字量越多但若超過免費額度越花錢，一開始可設200-400左右），接著尚有telegram的環境變數TELEGRAM_BOT_TOKEN。設定好後開始部屬，可能要花上一些時間。



------
#### telegram和openai api設置請參考： 
https://github.com/howarder3/GPT-Linebot-python-flask-on-vercel<br>
https://ithelp.ithome.com.tw/articles/10245264<br>
https://tcsky.cc/tips-01-telegram-chatbot/
