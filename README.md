# Environ:

Multi platform library for system dependend directories.

Based on the github project [https://github.com/shibukawa/configdir](https://github.com/shibukawa/configdir)

This library is designed to get the default directories of the operation
system like HOME, TMP or configuration directories. It implements the
different defaults of the operation system and gives an unified api to
access those directories or files.

Implementation details per operation system:
1. Windows:

   All the registry entries are based on the path `Computer\HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`

    |                    | Var-Name                    | Default                              | Reqistry key |
    | ------------------ | --------------------------- | ------------------------------------ | ------------ |
    | HOME               | $env:UserProfile            | C:\User\username                     |              |
    | UserConfig         | $env:AppData                | C:\Users\username\AppData\Roaming    |              |
    | GlobalConfig       | $env:AllUsersProfile        | C:\ProgramData                       |              |
    | UserTemp           | $env:Temp                   | C:\Users\username\AppData\Local\Temp |              |
    | UserCache          | $env:LocalAppData           | C:\Users\username\AppData\Local      |              |
    | GlobalCache        | $env:ProgramData            | C:\ProgramData                       |              |
    | UserProgramData    | $env:LocalAppData           | C:\Users\username\AppData\Local      |              |
    | GlobalProgramData  | $env:ProgramFiles           | C:\ProgramFiles                      |              |
    | Desktop            |                             | C:\Users\username\Desktop            | Desktop      |
    | Documents          |                             | C:\Users\username\Documents          | Personal     |
    | Downloads          |                             | C:\Users\username\Downloads          | {374DE290-123F-4565-9164-39C4925E467B} |
    | Pictures           |                             | C:\Users\username\Pictures           | My Pictures  |
    | Music              |                             | C:\Users\username\Music              | My Music     |
    | Videos             |                             | C:\Users\username\Videos             | My Video     |
    | SaveGames          |                             | C:\Users\username\Saved Games        | {4C5C32FF-BB9D-43B0-B5B4-2D72E54EAAA4} |

2. Linux / Unix / BSD:

    |                    | Var-Name                    | Default                              |
    | ------------------ | --------------------------- | ------------------------------------ |
    | HOME               | $HOME                       | /home/{username}                     |
    | UserConfig         | $XDG_CONFIG_HOME            | $HOME/.config                        |
    | GlobalConfig       | $XDG_CONFIG_DIRS            | /etc/xdg                             |
    | UserTemp           | $XDG_RUNTIME_DIR            |                                      |
    | UserCache          | $XDG_CACHE_HOME             | $HOME/.cache                         |
    | GlobalCache        |                             | /var/cache                           |
    | UserProgramData    | $XDG_DATA_HOME              | $HOME/.local/share                   |
    | GlobalProgramData  | $XDG_DATA_DIRS              | /usr/share                           |
    | Desktop            | $XDG_DESKTOP_DIR            | $HOME/Desktop                        |
    | Documents          | $XDG_DOCUMENTS_DIR          | $HOME/Documents                      |
    | Downloads          | $XDG_DOWNLOAD_DIR           | $HOME/Downloads                      |
    | Pictures           | $XDG_PICTURES_DIR           | $HOME/Pictures                       |
    | Music              | $XDG_MUSIC_DIR              | $HOME/Music                          |
    | Videos             | $XDG_VIDEOS_DIR             | $HOME/Videos                         |
    | SaveGames          | $XDG_DATA_HOME              | $HOME/.local/share                   |

3. macOS:

    |                    | Var-Name                    | Default                              |
    | ------------------ | --------------------------- | ------------------------------------ |
    | HOME               | $HOME                       | /Users/username                      |
    | UserConfig         |                             | $HOME//Library/Application Support   |
    | GlobalConfig       |                             | /Library/Application Support         |
    | UserTemp           | $TMPDATA                    | /var/folders/xl/random_stuff/T/      |
    | UserCache          |                             | $HOME/Library/Cache                  |
    | GlobalCache        |                             | /Library/Cache                       |
    | UserProgramData    |                             | $HOME//Library/Application Support   |
    | GlobalProgramData  |                             | /Library/Application Support         |
    | Desktop            |                             | $HOME/Desktop                        |
    | Documents          |                             | $HOME/Documents                      |
    | Downloads          |                             | $HOME/Downloads                      |
    | Pictures           |                             | $HOME/Pictures                       |
    | Music              |                             | $HOME/Music                          |
    | Videos             |                             | $HOME/Videos                         |
    | SaveGames          |                             | $HOME/Library/Application Support   |

Non existing values will be filled with the most similar value. If an OS has no global temp
 instead of nothing environ will return the local temp to enable a simple workflow
 where the developer has not to bother about checking for existence of that configuration
 key.

Regarding MacOS I have very little information about program structures
 and how program data is distinguished from configuration files. To get
 a little bit more structure into the application directories configuration
 file will go into a subfolder `config` and data files will go into a subfolder
 `data`.

License
=======

MIT
