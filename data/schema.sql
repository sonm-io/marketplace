CREATE TABLE IF NOT EXISTS orders
(
  idx INTEGER PRIMARY KEY AUTOINCREMENT,
  id TEXT,
  type INTEGER DEFAULT 0 NOT NULL,
  supplier_id TEXT,
  buyer_id TEXT,
  price TEXT DEFAULT 0 NOT NULL,
  slot_buyer_rating INTEGER DEFAULT 0 NOT NULL,
  slot_supplier_rating INTEGER DEFAULT 0 NOT NULL,
  slot_duration INTEGER DEFAULT 0 NOT NULL,
  resources_cpu_cores INTEGER DEFAULT 0 NOT NULL,
  resources_ram_bytes  INTEGER DEFAULT 0 NOT NULL,
  resources_gpu_count INTEGER DEFAULT 0 NOT NULL,
  resources_storage INTEGER DEFAULT 0 NOT NULL,
  resources_net_inbound INTEGER DEFAULT 0 NOT NULL,
  resources_net_outbound INTEGER DEFAULT 0 NOT NULL,
  resources_net_type INTEGER DEFAULT 0 NOT NULL,
  resources_properties TEXT,
  status INTEGER NOT NULL DEFAULT 1
);

CREATE UNIQUE INDEX IF NOT EXISTS orders_id_uindex on orders (id);