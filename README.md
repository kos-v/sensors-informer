# Sensors informer

## Building, installation, etc

### Compilation

To compile, you must have installed Go at least version 16.

1. `git clone https://github.com/kos-v/sensors-informer.git`

2. `cd sensors-informer`

3. `make`

### Installation

You can install this service and the global configuration file to the system directory.

Run (you may need root-rights): `make install`

### Uninstallation

To uninstall run (you may need root-rights): `make uninstall`

## Configuring

To start the service, you need a configuration file, a template of which is in the `./etc/conf/` repository folder. When installing the service to the system directory, the template will be copied to the directory `/etc/sensors-informer`, after which you need to configure the necessary channels and other parameters.  

It is also possible to load a local config file, for this put the config in the same directory with the executable file.  

In addition, you can pass the path to the config via the `-config=/path/config.yml` argument when starting the service.

Only one config can be loaded and with this priority:

1. Specific config passed via command line argument  `-config=/path/config.yml`.  
2. Local config: `./config.yml`.  
3. Global config: `/etc/sensors-informer/config.yml`