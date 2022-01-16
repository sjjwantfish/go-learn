# 生成数据前获取info memory

# Memory
used_memory:1074880
used_memory_human:1.03M
used_memory_rss:2531328
used_memory_rss_human:2.41M
used_memory_peak:1135328
used_memory_peak_human:1.08M
used_memory_peak_perc:94.68%
used_memory_overhead:1029280
used_memory_startup:1011840
used_memory_dataset:45600
used_memory_dataset_perc:72.34%
allocator_allocated:1029760
allocator_active:2493440
allocator_resident:2493440
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.42
allocator_frag_bytes:1463680
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.02
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.46
mem_fragmentation_bytes:1501568
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0


# 使用脚本生成数据并插入
脚本中控制key的字节长度
```shell
python3 gen_data.py
```

# 获取插入数据后的info memory
从文件夹下对应字节的文件中查看

# 使用redis-benchmark工具测试
从文件夹下对应字节的文件中查看


# 统计
| key字节数 | 每个key平均大小 |
| ---       | ---             |
| 10        | 191kb           |
| 20        | 103kb           |
| 50        | 46kb            |
| 100       | 282kb           |
| 200       | 187kb           |
| 1k        | 114kb           |
| 5k        | 101kb           |



