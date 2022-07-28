import discord
from discord.ext import commands, tasks
from pymongo import MongoClient
import redis
import json
import os
from dotenv import load_dotenv

#load token
BASEDIR = os.path.abspath(os.path.dirname(__file__))
load_dotenv(os.path.join(BASEDIR, 'config.env'))
TOKEN=os.getenv("TOKEN")

#mongo
client = MongoClient('localhost', 27017)
db = client["CosmosHubSubcribeEvent"]
collection = db["config"]

#redis
r = redis.Redis(host="localhost", port='6379')
pub = r.pubsub()
pub.psubscribe("all")


#discord
intents = discord.Intents.default()
intents.members = True

bot = commands.Bot(command_prefix='!',
                   intents=intents)

bot.remove_command("help")

def get_warning_embed(detail):
    embed = discord.Embed(title="Warning", color=0xFF0000)
    embed.add_field(
        name="Detail:", value=detail, inline=False)
    return embed

def get_high_embed(detail):
    embed = discord.Embed(title="High", color=0xFF8000)
    embed.add_field(
        name="Detail:", value=detail, inline=False)
    return embed

def get_medium_embed(detail):
    embed = discord.Embed(title="Medium", color=0xFFFF00)
    embed.add_field(
        name="Detail:", value=detail, inline=False)
    return embed

def get_low_embed(detail):
    embed = discord.Embed(title="Low", color=0x00FF00)
    embed.add_field(
        name="Detail:", value=detail, inline=False)
    return embed

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


@tasks.loop(seconds=1)
async def my_background_task():
    data = pub.get_message()
    if data:
        if type(data["data"]) is bytes:
            json_data=json.loads(data["data"].decode('utf-8'))
            amount=int(json_data["amount"][0:-5])
            warning_list=collection.find({"warning": {"$lt": amount}})
            medium_list=collection.find({"medium": {"$lt": amount}, "high":{"$gt":amount}})
            high_list=collection.find({"high": {"$lt": amount}, "warning":{"$gt":amount}})
            low_list=collection.find({"low": {"$lt": amount}, "medium":{"$gt":amount}})
            for x in warning_list:
                user = await bot.fetch_user(x["_id"])
                await user.send(embed=get_warning_embed(json_data))
            for x in medium_list:
                user = await bot.fetch_user(x["_id"])
                await user.send(embed=get_medium_embed(json_data))
            for x in high_list:
                user = await bot.fetch_user(x["_id"])
                await user.send(embed=get_high_embed(json_data))
            for x in low_list:
                user = await bot.fetch_user(x["_id"])
                await user.send(embed=get_low_embed(json_data))


@my_background_task.before_loop
async def my_background_task_before_loop():
    await bot.wait_until_ready()

my_background_task.start()

bot.run(TOKEN)
