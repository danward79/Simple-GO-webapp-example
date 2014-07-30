Installing Apple Filing Protocol on Rasbian
====

Before we start. Make sure Raspbian is up to date.

````
sudo apt-get update && sudo apt-get upgrade
````

Use apt to install Netatalk.
```
sudo apt-get install netatalk
```
You may have noticed your Pi appear in your Mac Finder window?

Make sure that the Netatalk is not running, so that we can adjust the settings.
````
sudo /etc/init.d/netatalk stop
sudo nano /etc/netatalk/AppleVolumes.default
````

````
:DEFAULT: options:upriv,usedots,rw,tm
````

````
sudo /etc/init.d/netatalk start
````

Mount usb drive
````
sudo mkdir /mnt/4GBStick
sudo mkfs.ext4 /dev/sda1 -L untitled
````

find drive
```
ls -l /dev/disk/by-uuid/
```

mount it
```
sudo mount /dev/sda1 /mnt/www
sudo chown pi:pi /mnt/www
sudo chmod 777 /mnt/www
```

ntfs-3g for NTFS Drives
vfat for FAT32 Drives
ext4 for ext4 Drives

backup fstab
```
sudo cp /etc/fstab /etc/fstab.backup
```

add mount
```
/dev/sda1 /mnt/www auto defaults,user 0 1
```

```
sudo reboot
```