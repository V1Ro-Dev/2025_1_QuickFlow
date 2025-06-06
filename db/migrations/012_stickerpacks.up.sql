create table if not exists sticker_pack(
                                           id uuid primary key default gen_random_uuid(),
                                           name TEXT NOT NULL,
                                           creator_id uuid references "user"(id) on delete cascade,
                                           created_at TIMESTAMP DEFAULT NOW(),
                                           updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sticker (
                                       id UUID primary key default gen_random_uuid(),
                                       sticker_pack_id UUID REFERENCES sticker_pack(id) ON DELETE CASCADE,
                                       sticker_url text not null,
                                       created_at TIMESTAMP DEFAULT NOW()
);
