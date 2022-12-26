Title    : Excessive (Mr. Pants' Excessive Overkill for Quake III Arena)
Filename : excessive_q3_server_004.zip
Version  : 004
Date     : 3-1-04
Author   : Dan "Mr. Pants" Schoenblum <mrpants@gamespy.com>
Web Page : http://www.planetquake.com/excessive

-----
Intro
-----
Excessive is a mod for Quake III Arena.  It's a server-side only mod, so if
you only want to play on Excessive servers, you do not need to install this.
If you plan on running an Excessive server, or want to play Excessive
single-player then this is for you.

-------------------------
Installation Instructions
-------------------------
This is the server package for Excessive.  Aside from the contents of this 
package, all you need to setup an Excessive server is Quake III Arena.

First make sure you have the latest version of Quake III, which is 1.32.
It's available at:
http://www.fileplanet.com/files/60000/61607.shtml

One you have the latest patch installed, extract the contents of the
Excessive zip into your Quake III directory (the directory with quake3.exe
in it).  Make sure that you preserve the directory structure.  If you are
using pkunzip, use the -d option when unzipping.

The Excessive mod will be placed in a subdirectory "excessive" off of the main 
Quake III directory.  Excessive puts three files in your main Quake III
directory that can be used to play Excessive.  Two can be used to start servers:
excessive_dedicated_server.bat is used to run a dedicated server (just a
server), and excessive_listen_server.bat is used to run a listen server (server
and client).  If excessive_single_player.bat is used to play Excessive
single-player.

-------------------------
Running Excessive Servers
-------------------------
To run an Excessive server, you need to add the following option to the
command line when running the server:
+set fs_game Excessive
So, a typical command line will look something like this:
quake3.exe +set dedicated 2 +set fs_game excessive +sv_pure 0 +map q3dm4

Using "+set dedicated 2" causes the server to report itself to the backend
so that it shows up in server lists.  "+set dedicated 1" will run a dedicated
LAN server that does not report itself to the backend.

If when you run the server you get an error "fs_game is write protected", make 
sure that you are loading a map on the command line with the +map option.

-------------------------
Finding Excessive Servers
-------------------------
The best way to find Excessive servers is with GameSpy Arcade.  It can be
downloaded off of the GameSpy Arcade site:
http://www.gamespyarcade.com/
You can also use GameSpy3D <http://gamespy3d.com>, All Seeing Eye
<http://udpsoft.com/eye>, or the in-game browser.

-------------------------------------
Server Console Variables and Commands
-------------------------------------
Excessive has a set of variables and commands that can be used by the server
operator to control certain aspects of the game.

ex_instant (0/1): If this is 1, then weapon changes occur instantly.  If it is
0, then weapon changes take as long as they take in normal Quake III.  The
default is 1.

ex_quadstart (0/1): If this is 1, it acts like Release 2 of Excessive where
everyone gets quad damage on level changes.  If set to 0, then noone gets it.
It is 1 by default.

ex_joust (0/1/2): If 1, this turns on "Joust" mode, which allows players to fly.
To fly, players just need to keep pressing the jump button.  If it is 2, then
players can jump further/faster and get more air control, but cannot fly.  This
can be changed at any time, and is 2 by default.

ex_quadfly (0/1): If this is 1, then any player with quad damage can fly, as if
ex_joust was set to 1 for that player only.  This is 1 by default.

ex_motd <filename>: Server operators have the ability to add to Excessive's MOTD
(this is the message that flashes on the center of the screen when a player first
joins the server).  This variable specifies a file containing test that will get
appended to the Excessive MOTD.  This can be used to use a different MOTD on
different servers running out of the same Quake III folder.  The file can be more
than one line, and it can contain colors, just like colored nicknames.  To use
color, simply put in one of these markers wherever you want to start using a new
color:
black: ^0   red:  ^1   green:   ^2   yellow: ^3
blue:  ^4   cyan: ^5   magenta: ^6   white:  ^7
The default for this variable is motd.cfg.

ex_<weapon> (0/1): Allows specific weapons to be disabled on the server.  Set
a weapon to 0 to disable it.  All weapons are enabled (1) by default.  The
weapon specific variables are:
ex_gauntlet   ex_machinegun   ex_shotgun   ex_grenade   ex_rocket
ex_lightning  ex_rail         ex_plasma    ex_bfg       ex_hook
Changes to this variable do not take effect until the next map change or restart.

ex_spawn (weapon number/name): Sets the weapon to be used when a player spawns.
The spawn weapon can be set as a number, in which case the number used to pick
that weapon should be used (1 is gauntlet, 2 is machinegun, etc.).  Or, it can
be set as a name ("gauntlet", "machinegun", etc.).  When using a name, you only
need to type enough to uniquely identify that weapon, which is at most the first
2 characters of the weapon's name ("ga", "m", "s", "gr", "ro", "l", "ra", "p",
"b", "h").  If this is set to 0 or a blank string, then Excessive will choose a
spawn weapon.  The default value is an empty string.  Changes to this variable
do not take effect until the next map change or restart.

ex_reset: This is a command which resets all of the weapon settings.  All of the
ex_<weapon> variables and ex_spawn are set back to their defaults.  It is used to
enable all weapons.  These changes will not take effect until the next map change
or restart.

ex_clear: This is a command which clears all of the weapons settings.  All of the
ex_<weapon> variables are set to 0, and ex_spawn is set back to an empty string.
This is used to disable all weapons.  These changes will not take effect until
the next map change or restart.

ex_settings: This is a command that is used to see all of the current settings for
the ex_<weapon> and ex_spawn variables.  Note that it prints the variables as they
were when the map started, not what they are currently set to.

--------------------
Single-Player (Bots)
--------------------
To play Excessive single-player, first find then run the excessive_single_player.bat
file in the Quake III folder.  This will start up Quake III with Excessive loaded.
Then just start a single-player game and it will be Excessive.

----
Help
----
For help, try the Excessive forums:
http://forumplanet.com/planetquake/excessive/

--------------------------------------
Copyright and Distribution Permissions
--------------------------------------

This mod is freely distributable provided that this readme is distributed
as well and is unchanged.

Copyright Daniel E. Schoenblum 2000-2004.

DISCLAIMER: THE PROGRAM IS DEFINED AS THE QUAKE III MOD "EXCESSIVE" AND
  ALL FILES CONTAINED WITHIN.
  BECAUSE THE PROGRAM IS LICENSED FREE OF CHARGE, THERE IS NO WARRANTY
  FOR THE PROGRAM, TO THE EXTENT PERMITTED BY APPLICABLE LAW.  EXCEPT WHEN
  OTHERWISE STATED IN WRITING THE COPYRIGHT HOLDERS AND/OR OTHER PARTIES
  PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY OF ANY KIND, EITHER EXPRESSED
  OR IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
  MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.  THE ENTIRE RISK AS
  TO THE QUALITY AND PERFORMANCE OF THE PROGRAM IS WITH YOU.  SHOULD THE
  PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF ALL NECESSARY SERVICING,
  REPAIR OR CORRECTION.

------------
Availability
------------
This modification is available from the following places:
WWW   : http://www.planetquake.com/excessive

-------
Contact
-------
The author, Dan "Mr. Pants" Schoenblum, can be contacted by:
e-mail : mrpants@gamespy.com
ICQ    : 46786530

---------------
Version History
---------------
1 - Initial Public Release

2 - Second Public Release
Changes:
Excessified BFG -  spews short-fused grenades on impact
Excessified Lightning-Gun - massive amounts of explosive damage
Everyone gets 30 seconds of Quad on level changes
Shotgun fires twice as fast
Machinegun fires slower, but with more damage
Health and ammo regenerate
Health/ammo/weapon item pickups are removed
When you die, you explode

3 - Third Public Release
Changes:
Can disable starting maps with quad (set ex_quadstart to 0)
Can turn on Joust mode (set ex_joust to 1)
Can append to MOTD using motd.cfg
Compatible with 1.27 Q3
BFG grenades are now affected by quad
BFG grenades now uses the BFG splash MoD (was BFG)
Shotgun now uses the shotgun MoD (was rail)

4 - Final Public Release
Changes:
Grappling hook
Instant weapon changing (ex_instant)
Added the ability to specify the MOTD filename (ex_motd)
Enable/disable specific weapons (ex_<weapon>)
Set a spawn weapon (ex_spawn)
Weapon commands (ex_reset/ex_clear/ex_settings)
Removed falling damage
Death by BFG splash correctly identified
Users with quad can fly (ex_quadfly)
New joust mode (ex_joust 2)