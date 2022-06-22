# Kamailio Input Plugin

This plugin gather kamailio metrics from ctl plugin via binrpc

### Configuration:
```toml
[[inputs.kamailio]]
   # Kamailio socket
   socket = "unix:/tmp/kamailio_ctl"
   #Supported modules tm.stats core.shmmem core.tcp_info sl.stats tls.info htable.stats
   modules = "tm.stats core.shmmem core.tcp_info sl.stats tls.info"
   #nodeid  only used for tags
   nodeid = ""
```

### Tags
nodeid

### Metrics output example
```toml
core.shmmem.fragments=14i
core.shmmem.free=34594688i
core.shmmem.max_used=9599160i
core.shmmem.real_used=67108864i
core.shmmem.total=9445504i
core.shmmem.used=0i
core.tcp_info.max_connections=2048i
core.tcp_info.max_tls_connections=2048i
core.tcp_info.opened_connections=0i
core.tcp_info.opened_tls_connections=0i
core.tcp_info.readers=4i
core.tcp_info.write_queued_bytes=0i
sl.stats.200=0i
sl.stats.202=0i
sl.stats.2xx=0i
sl.stats.300=0i
sl.stats.301=0i
sl.stats.302=0i
sl.stats.400=0i
sl.stats.401=0i
sl.stats.403=0i
sl.stats.404=0i
sl.stats.407=0i
sl.stats.408=0i
sl.stats.483=0i
sl.stats.4xx=0i
sl.stats.500=0i
sl.stats.5xx=0i
sl.stats.6xx=0i
sl.stats.xxx=0i
tls.info.clear_text_write_queued_bytes=0i
tls.info.max_connections=2048i
tls.info.opened_connections=0i
tm.stats.2xx=0i
tm.stats.3xx=0i
tm.stats.4xx=0i
tm.stats.5xx=0i
tm.stats.6xx=0i
tm.stats.created=0i
tm.stats.current=0i
tm.stats.delayed_free=0i
tm.stats.freed=0i
tm.stats.replied_locally=0i
tm.stats.total=0i
tm.stats.total_local=0i
tm.stats.waiting=0i
htable.stats.all.0=0i
htable.stats.min.0=0i
htable.stats.slots.0=2048i
htable.stats.name.0=ext_subscription
htable.stats.all.10=0i
htable.stats.min.10=0i
htable.stats.slots.10=256i
htable.stats.all.1=0i
htable.stats.min.1=0i
htable.stats.name.10=nat_uri
htable.stats.all.11=0i
htable.stats.min.11=0i
htable.stats.slots.11=2048i
htable.stats.name.11=cnttcp
htable.stats.slots.1=2048i
htable.stats.all.12=0i
htable.stats.min.12=0i
htable.stats.slots.12=2048i
htable.stats.name.12=via
htable.stats.all.13=0i
htable.stats.min.13=0i
htable.stats.slots.13=2048i
htable.stats.name.13=did
htable.stats.all.14=0i
htable.stats.min.14=0i
htable.stats.slots.14=2048i
htable.stats.name.14=ch
htable.stats.all.15=0i
htable.stats.min.15=0i
htable.stats.slots.15=2048i
htable.stats.name.15=pbx_extension_number
htable.stats.all.16=0i
htable.stats.min.16=0i
htable.stats.slots.16=2048i
htable.stats.name.16=pbx_replacement_did
htable.stats.name.1=user_subscription
htable.stats.all.2=0i
htable.stats.min.2=0i
htable.stats.slots.2=2048i
htable.stats.name.2=base_callee_rU
htable.stats.all.3=0i
htable.stats.min.3=0i
htable.stats.slots.3=2048i
htable.stats.name.3=base_callee
htable.stats.all.4=0i
htable.stats.min.4=0i
htable.stats.slots.4=2048i
htable.stats.name.4=virt_callee
htable.stats.all.5=0i
htable.stats.min.5=0i
htable.stats.slots.5=2048i
htable.stats.name.5=base_caller
htable.stats.all.6=0i
htable.stats.min.6=0i
htable.stats.slots.6=2048i
htable.stats.name.6=virt_caller
htable.stats.all.7=0i
htable.stats.min.7=0i
htable.stats.slots.7=2048i
htable.stats.name.7=pbx_redirect_dst
htable.stats.all.8=0i
htable.stats.min.8=0i
htable.stats.slots.8=16384i
htable.stats.name.8=pbx_redirect_src
htable.stats.all.9=0i
htable.stats.min.9=0i
htable.stats.slots.9=16384i
htable.stats.name.9=pbx_nodeid


```
