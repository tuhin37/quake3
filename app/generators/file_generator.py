from uuid6 import uuid7
from app.model import Expected
import os


def create_file(expected: Expected):
    """this will create and return the paths to the files"""
    folder_name = "/usr/local/games/quake3/baseq3/"
    name = str(uuid7()) + ".cfg"
    expected_file_name = os.path.join(folder_name, name)
    file_contents = f"""
    seta bot_enable {expected.bot_enable}
seta bot_nochat {expected.bot_nochat}    
seta g_spskill {expected.bot_skill}      
seta bot_minplayers {expected.bot_minplayers}
set dm1 "map {expected.map}; set nextmap vstr dm1"
vstr dm1

seta sv_hostname "{expected.hostname}"
seta g_motd "{expected.message}"     
seta sv_maxclients {expected.maxClients}
seta sv_pure {expected.pure}            
seta g_quadfactor {expected.quadfactor} 
seta g_friendlyFire {expected.friendlyFire} 


seta g_gametype {expected.gameType}
seta g_teamAutoJoin 0           
seta g_teamForceBalance 0       
seta timelimit {expected.timelimit}
seta capturelimit 8             
seta fraglimit {expected.fraglimit}


seta g_weaponrespawn {expected.weaponrespawn}
seta g_inactivity {expected.inactivity}      
seta g_forcerespawn {expected.forcerespawn}  
seta g_log server.log           
seta logfile 3                  
seta rconpassword "{expected.rconpassword}"      

seta rate "{expected.rate}"               
seta snaps "{expected.snaps}"                 
seta cl_maxpackets "{expected.maxpackets}"    
seta cl_packetdup "{expected.packetdup}"      
    """
    with open(expected_file_name, "w") as writer:
        writer.write(file_contents)
    return name
