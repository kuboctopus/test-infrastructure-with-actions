$TTL 3h
@ IN SOA ns1.example.com. admin.example.com. (
  2022031801 ; Serial
  3h ; Refresh
  15m ; Retry
  1w ; Expire
  1h ; Negative Cache TTL
)

www IN A 192.168.1.100
