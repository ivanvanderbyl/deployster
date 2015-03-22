package server

/**
 * A Systemd service template for starting deployster-sidekick as a sidekick
 * service.
 *
 * @type {string}
 */
const dockerSidekickUnitTemplate = `
[Unit]
Description={{.ServiceName}} Sidekick
BindsTo={{.ServiceName}}.service
After={{.ServiceName}}.service

[Service]
EnvironmentFile=/etc/environment
User=core
TimeoutStartSec=0
ExecStartPre=/usr/bin/docker pull {{.SidekickServiceImage}}
ExecStart=/usr/bin/docker run --rm=true --name {{.ServiceName}}-%i {{.SidekickServiceImage}}
`
