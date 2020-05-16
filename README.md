# dds_joystick
This application sends and receives joystick (Dualshock3) events over DDS, leveraging [Go Connector](https://github.com/rticommunity/rticonnextdds-connector-go) and [Gobot](https://github.com/hybridgroup/gobot). 

### Getting Started
This requires `sdl2` to be installed. 
```bash
$ wget https://www.libsdl.org/release/SDL2-2.0.8.tar.gz
$ tar -zxvf SDL2-2.0.8.tar.gz
$ cd SDL2-2.0.8/
$ ./configure && make && sudo make install
```

RTI Go Connector requires [Git LFS](https://github.com/git-lfs/git-lfs/wiki/Installation) to check out the Connector C library files properly. 
```bash
$ curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | sudo bash
$ sudo apt-get install git-lfs
$ git lfs install
```

Build and run:
```bash
$ go get -v github.com/kyoungho/dds_joystick
$ go build ~/go/src/github.com/kyoungho/dds_joystick/writer/js_writer.go
$ export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:~/go/src/github.com/rticommunity/rticonnextdds-connector-go/rticonnextdds-connector/lib/armv6vfphLinux3.xgcc4.7.2
$ ./js_writer
```

You should replace `armv6vfphLinux3.xgcc4.7.2` with your architecture. 

