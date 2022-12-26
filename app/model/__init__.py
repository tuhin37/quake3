from typing import Optional
from pydantic import BaseModel, validator
import re
from varname import nameof
import os

name_reg = re.compile(r"[a-zA-Z0-9_-]{1,256}$")
map_reg = re.compile(r"[a-zA-Z0-9_-]{1,256}$")
hostname_reg = re.compile(r"[a-zA-Z0-9_-]{1,256}$")
message_reg = re.compile(r"^[a-zA-Z0-9._ \t-]{1,255}$")
maxClients_reg = re.compile(r"^[1]{1}[0-6]{1}$|[1-9]{1}$")
pure_reg = re.compile(r"\b(?:true|false)\b$")
quadfactor_reg = re.compile(r"^[1-4]{1}$")
friendlyFire_reg = re.compile(r"\b(?:true|false)\b$")
gameType_reg = re.compile(r"\b(?:FFA|CTF|TDM|Tourney)\b$")
timelimit_reg = re.compile(r"^[0-9]+$")
fraglimit_reg = re.compile(r"^[0-9]+$")
bot_enable_reg = re.compile(r"^[0-1]{1}$")
bot_nochat_reg = re.compile(r"^[0-1]{1}$")
bot_skill_reg = re.compile(r"^[0-9]{1}$")
bot_minplayers_reg = re.compile(r"^[0-9]{1}$")
weaponrespawn_reg = re.compile(r"^[0-9]+$")
inactivity_reg = re.compile(r"^[0-9]+$")
forcerespawn_reg = re.compile(r"^[0-1]{1}$")
rconpassword_reg = re.compile(r"[a-zA-Z0-9_-]{1,256}$")
rate_reg = re.compile(r"^[0-9]+$")
snaps_reg = re.compile(r"^[0-9]+$")
maxpackets_reg = re.compile(r"^[0-9]+$")
packetdup_reg = re.compile(r"^[0-1]{1}$")

converter = {}
converter[gameType_reg] = {"FFA": 0, "Tourney": 1, "TD": 3, "CTF": 4}


def setter(name, regex_checker, nameof):
    print('{} checking value {}'.format(nameof, name))
    val = regex_checker.match(name)
    if val:
        got_str = val.group()
        if got_str:
            if regex_checker in converter:
                return converter[regex_checker][got_str]
            return str(got_str)
    else:
        print('Invalid parameter for {}'.format(nameof))
        raise ValueError("Invalid parameter: {}".format(nameof))


class Expected(BaseModel):
    name: Optional[str]
    map: Optional[str]
    hostname: Optional[str]
    message: Optional[str]
    maxClients: Optional[str]
    pure: Optional[str]
    quadfactor: Optional[str]
    friendlyFire: Optional[str]
    gameType: Optional[str]
    timelimit: Optional[str]
    fraglimit: Optional[str]
    bot_enable: Optional[str]
    bot_nochat: Optional[str]
    bot_skill: Optional[str]
    bot_minplayers: Optional[str]
    weaponrespawn: Optional[str]
    inactivity: Optional[str]
    forcerespawn: Optional[str]
    rconpassword: Optional[str]
    rate: Optional[str]
    snaps: Optional[str]
    maxpackets: Optional[str]
    packetdup: Optional[str]
    api_key: str  # mandatory

    class Config:
        validate_assignment = True
    
    @validator("api_key", pre=True)
    def set_api_key(cls, api_key):
        if api_key != os.getenv('API_KEY'):
            raise ValueError('Invalid API Key')
        return api_key

    @validator("name", pre=True)
    def set_name(cls, name):
        return setter(name, name_reg, nameof(name))

    @validator("map", pre=True)
    def set_map(cls, map):
        return setter(map, map_reg, nameof(map))

    @validator("hostname", pre=True)
    def set_hostname(cls, hostname):
        return setter(hostname, hostname_reg, nameof(hostname))

    @validator("message", pre=True)
    def set_message(cls, message):
        return setter(message, message_reg, nameof(message))

    @validator("maxClients", pre=True)
    def set_maxClients(cls, maxClients):
        return setter(maxClients, maxClients_reg, nameof(maxClients))

    @validator("pure", pre=True)
    def set_pure(cls, pure):
        return setter(pure, pure_reg, nameof(pure))

    @validator("quadfactor", pre=True)
    def set_quadfactor(cls, quadfactor):
        return setter(quadfactor, quadfactor_reg, nameof(quadfactor))

    @validator("friendlyFire", pre=True)
    def set_friendlyFire(cls, friendlyFire):
        return setter(friendlyFire, friendlyFire_reg, nameof(friendlyFire))

    @validator("gameType", pre=True)
    def set_gameType(cls, gameType):
        return setter(gameType, gameType_reg, nameof(gameType))

    @validator("fraglimit", pre=True)
    def set_fraglimit(cls, fraglimit):
        return setter(fraglimit, fraglimit_reg, nameof(fraglimit))

    @validator("bot_enable", pre=True)
    def set_bot_enable(cls, bot_enable):
        return setter(bot_enable, bot_enable_reg, nameof(bot_enable))

    @validator("bot_nochat", pre=True)
    def set_bot_nochat(cls, bot_nochat):
        return setter(bot_nochat, bot_nochat_reg, nameof(bot_nochat))

    @validator("bot_skill", pre=True)
    def set_bot_skill(cls, bot_skill):
        return setter(bot_skill, bot_skill_reg, nameof(bot_skill))

    @validator("bot_minplayers", pre=True)
    def set_bot_minplayers(cls, bot_minplayers):
        return setter(bot_minplayers, bot_minplayers_reg, nameof(bot_minplayers))

    @validator("weaponrespawn", pre=True)
    def set_weaponrespawn(cls, weaponrespawn):
        return setter(weaponrespawn, weaponrespawn_reg, nameof(weaponrespawn))

    @validator("inactivity", pre=True)
    def set_inactivity(cls, inactivity):
        return setter(inactivity, inactivity_reg, nameof(inactivity))

    @validator("timelimit", pre=True)
    def set_timelimit(cls, timelimit):
        return setter(timelimit, timelimit_reg, nameof(timelimit))

    @validator("forcerespawn", pre=True)
    def set_forcerespawn(cls, forcerespawn):
        return setter(forcerespawn, forcerespawn_reg, nameof(forcerespawn))

    @validator("rconpassword", pre=True)
    def set_rconpassword(cls, rconpassword):
        return setter(rconpassword, rconpassword_reg, nameof(rconpassword))

    @validator("rate", pre=True)
    def set_rate(cls, rate):
        return setter(rate, rate_reg, nameof(rate))

    @validator("snaps", pre=True)
    def set_snaps(cls, snaps):
        return setter(snaps, snaps_reg, nameof(snaps))

    @validator("maxpackets", pre=True)
    def set_maxpackets(cls, maxpackets):
        return setter(maxpackets, maxpackets_reg, nameof(maxpackets))

    @validator("packetdup", pre=True)
    def set_packetdup(cls, packetdup):
        return setter(packetdup, packetdup_reg, nameof(packetdup))
