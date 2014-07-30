Explaination of MacOSX Environmental Variables
===========

* http://stackoverflow.com/questions/603785/environment-variables-in-mac-os-x/4567308#4567308

There are several places where you can set environment variables.

*  ~/.profile: use this for variables you want to set in all programs launched from the terminal (note that, unlike on Linux, all shells opened in Terminal.app are login shells).
* ~/.bashrc: this is invoked for shells which are not login shells. Use this for aliases and other things which need to be redefined in subshells, not for environment variables that are inherited.
*	/etc/profile: this is loaded before ~/.profile, but is otherwise equivalent. Use it when you want the variable to apply to terminal programs launched by all users on the machine (assuming they use bash).
* ~/.MacOSX/environment.plist: this is read by loginwindow on login. It applies to all applications, including GUI ones, except those launched by Spotlight in 10.5 (not 10.6). It requires you to logout and login again for changes to take effect. This file is no longer supported as of OS X 10.8.
* your user's launchd instance: this applies to all programs launched by the user, GUI and CLI. You can apply changes at any time by using the setenv command in launchctl. In theory, you should be able to put setenv commands in ~/.launchd.conf, and launchd would read them automatically when the user logs in, but in practice support for this file was never implemented. Instead, you can use another mechanism to execute a script at login, and have that script call launchctl to set up the launchd environment.
* /etc/launchd.conf: this is read by launchd when the system starts up and when a user logs in. They affect every single process on the system, because launchd is the root process. To apply changes to the running root launchd you can pipe the commands into sudo launchctl.

The fundamental things to understand are:

* environment variables are inherited by a process's children at the time they are forked.
* the root process is a launchd instance, and there is also a separate launchd instance per user session.
* launchd allows you to change its current environment variables using launchctl; the updated variables are then inherited by all new processes it forks from then on.