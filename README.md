# telegrambothellow


Based on https://dev.to/talentlessguy/create-and-deploy-golang-telegram-bot-2dl0 tutorial


1. Create repository

echo "# telegrambothellow" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin https://github.com/jejimenez/telegrambothellow.git
git push -u origin master


2. Create Heroku project 

Go to Settings, there connect your github repository to Heroku app

3. Now add our telegram token and app url to Config Vars. 
GOVERSION = 1.13
PUBLIC_URL = <herokuprojectname>.herokuapp.com
TOKEN = <Telegram API KEY>

4. Add heroku to local project in root git project path
heroku git:remote -a telegram-bot-hellow

5. Create go files 
go mod init https://github.com/jejimenez/telegrambothellow.git

6. Create the buildpack
heroku buildpacks:set heroku/go

7. Commit and push git 
git add .
git commit -m "Heroku setup"
git push
git push heroku master