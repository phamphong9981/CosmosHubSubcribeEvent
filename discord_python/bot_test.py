
from turtle import title
import discord
from discord.ext import commands, tasks
from pymongo import MongoClient
import redis


client = MongoClient('localhost', 27017)
db = client["CosmosHubSubcribeEvent"]
collection = db["config"]
intents = discord.Intents.default()
intents.members = True
r = redis.Redis(host="localhost", port='6379')
pub = r.pubsub()
pub.psubscribe("*")

bot = commands.Bot(command_prefix='!',
                   intents=intents)

bot.remove_command("help")


@bot.event
async def on_command_error(ctx, error):
    if isinstance(error, discord.ext.commands.errors.CommandNotFound):
        await ctx.send("That command wasn't found! Please type `!help` to get help.")
    elif isinstance(error, discord.ext.commands.errors.MissingRequiredArgument):
        await ctx.send("Missing required arguments. Please type `!help` to get help")
    else:
        raise error


@bot.command()
async def help(ctx):
    embed = discord.Embed(title="List command", color=0xFF5733,
                          description="Welcome to help section. This is an embed that will show list command working with bot")
    embed.add_field(
        name="!info", value="Get your registered threshold log", inline=False)
    embed.add_field(name="!config <low> <medium> <high> <warning>",
                    value="Get your registered threshold log", inline=False)
    await ctx.send(embed=embed)


@bot.command()
async def info(ctx):
    myquery = {"_id": ctx.message.author.id}
    info = collection.find_one(myquery)
    if info:
        embed = discord.Embed()
        embed.add_field(
            name="Low", value=info["low"], inline=False)
        embed.add_field(name="Medium",
                        value=info["medium"], inline=False)
        embed.add_field(
            name="High", value=info["high"], inline=False)
        embed.add_field(
            name="Warning", value=info["warning"], inline=False)
        await ctx.send(embed=embed)
    else:
        await ctx.send("You havent config your threshold. Please type `!help` to get help.")


@bot.command()
async def config(ctx, low: int, medium: int, high: int, warning: int):
    collection.update_one({'_id': ctx.message.author.id}, {"$set": {
                          '_id': ctx.message.author.id, "low": low, "medium": medium, "high": high, "warning": warning}}, True)

# @bot.command()
# async def status(ctx):
#     embed = discord.Embed(title="List status", color=0xFF5733,
#                           description="Welcome to help section. This is an embed that will show list available status")
#     embed.add_field(
#         name="!info", value="Get your registered threshold log", inline=False)
#     embed.add_field(name="!config <status> <threshold>",
#                     value="Get your registered threshold log", inline=False)
#     await ctx.send(embed=embed)

@tasks.loop(seconds=1)
async def my_background_task():
    data=pub.get_message()
    if data:
        user = await bot.fetch_user(983558322819059722)
        await user.send('TEST!')


@my_background_task.before_loop
async def my_background_task_before_loop():
    await bot.wait_until_ready()

my_background_task.start()

bot.run("OTY1MjU0NTA0MjQ3MzM3MDIw.Gb7bxS.TFAaohtFr5l-xhYdAZh-p8ePSlXjSQm23xSJj0")
