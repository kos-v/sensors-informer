# Sensors informer

This service checks data from lm_sensors (Linux monitoring sensors) and
sends notifications when critical indicators are reached.

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

### Description of configuration parameters

```yaml
# lm_sensors options
lmSensors:
  # Path to the "sensors" command
  command: sensors

# List of channels to which you can send a notification.
# Enabling/disabling a channel is controlled by the "enable" parameter, which by default is false.
channels:
  # Sends a notification to a file, the path to which is specified in the "path" parameter
  # Can be used, for example, when testing, debugging or other cases.
  file:
    # Enable/disable
    enable: true
    # The path to a file
    path: ./notifications.out
  # Sends a notification to the notify-send program, which is often installed in gnome
  # and other graphical unix-shells on a desktop.
  # See https://manpages.ubuntu.com/manpages/xenial/man1/notify-send.1.html
  notifySend:
    # Enable/disable
    enable: false
    # Path to the "notify-send" command
    command: notify-send
    # Specifies the urgency level (low, normal, critical)
    # See notify-send documentation.
    urgency: critical
    # Specifies the timeout in milliseconds at which to expire the notification.
    # See notify-send documentation.
    expireTime: 10
    # Specifies basic extra data to pass. Valid types are int, double, string and byte.
    # See notify-send documentation.
    hint:
  # Sends a notification to a email via smtp
  smtp:
    # Enable/disable
    enable: false
    # Username for authentication
    username: sender@example.local
    # Password for authentication
    password: "password"
    # Connection SMTP host
    host: smtp.example.local
    # Connection port
    port: 465
    # Sender data
    from:
      # Sender's email address
      email: sender@example.local
      # Sender name
      name: Sensors Informer
    # List of email recipients
    to: [ recipient@example.local]
    # Encryption type. Valid values: ssl, tls, starttls.
    encryption: ssl
    # TLS encryption parameters
    tls:
      # InsecureSkipVerify controls whether a client verifies the server's certificate chain and host name.
      # If InsecureSkipVerify is true, TLS  accepts any certificate presented by the server and any host name in that certificate.
      # In this mode, TLS is susceptible to machine-in-the-middle attacks unless custom verification is used.
      insecureSkipVerify: false
      # ServerName is used to verify the hostname on the returned certificates unless InsecureSkipVerify is given.
      # It is also included in the client's handshake to support virtual hosting unless it is an IP address.
      serverName:
  # Sends a notification to your telegram bot
  telegramBot:
    # Enable/disable
    enable: false
    # The token that you received when creating a bot through @BotFather
    token: "0000000000:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    # The id of a chat that you initiated with your bot.
    # The chat ID can only be obtained by calling the telegram api method: https://api.telegram.org/bot{token}/getUpdates
    # where, {token} is a token issued when creating a bot through @BotFather.
    chatId: 000000000

# Sensors options
sensors:
  # Temperature check interval in seconds
  pollingInterval: 10
  # Temperature sensor options
  temperature:
    # Temperature at which a notification will be sent
    critical: 80
    # The unit of measure for the critical parameter. Default: celsius.
    # Valid options: "c" (celsius), "f" (fahrenheit).
    unit: c

# Notification options
report:
  # Format of notification
  format:
    # The format in which the temperature will be displayed in a notification.
    # Valid options: "c" (celsius), "f" (fahrenheit).
    temperatureUnit: с
    # Format of notification title
    title:
      # Title text.
      # If you specify {timestamp}, then the title will contain the date when the notification was generated.
      text: "Critical temperature readings [{timestamp}]:"
  # Time in seconds when a notification can be resubmitted
  repeatTimeout: 1800
```