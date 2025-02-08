# Wikipedia RecentChanges Discord Bot

## Content
- [Technologies](#technologies-)
- [How to run](#how-to-run)
- [How to use](#how-to-use)

### Technologies 
1. Go Programming Language
2. Docker to deploy bot

### How to run
1. Create a bot, add it to your server.
2. For both steps create ```.env``` file with ```BOT_TOKEN``` variable and insert your discord bot token
#### Using docker:
1. Make sure Docker is installed on your machine
2. Install project on your computer
3. In terminal run ```docker build --tag 'image-name' .```
   This will give you an image on your local machine that you can create a container from.
4. Run ```docker run 'image-name'```. This command will run container based on an image created in the previous step
### How to run locally:
1. In terminal run ```go run main.go``` in terminal you will see ```Bot is running!```.

### How to use
```!recent``` command will make bot to send recent changes made on wikipedia pages. Default language is set to English(en) <br> 
```!setLang [language]``` will change default language parameter. You can see language codes at [Languages](https://en.wikipedia.org/wiki/List_of_Wikipedias) <br>
```!stop``` bot will stop sending messages
