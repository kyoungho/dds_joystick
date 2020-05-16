# dds_joystick
This application publishes joystick (Dualshock3) events over DDS, leveraging [Go Connector](https://github.com/rticommunity/rticonnextdds-connector-go) and [Gobot](https://github.com/hybridgroup/gobot). 

### Getting Started
RTI Go Connector requires [Git LFS](https://github.com/git-lfs/git-lfs/wiki/Installation) to check out the Connector C library files properly. 
```bash
$ curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | sudo bash
$ sudo apt-get install git-lfs
$ git lfs install
```


