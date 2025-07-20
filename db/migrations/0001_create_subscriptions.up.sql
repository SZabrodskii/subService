CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE subscriptions (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               service_name TEXT NOT NULL,
                               price INTEGER NOT NULL,
                               user_id UUID NOT NULL,
                               start_date DATE NOT NULL,
                               end_date DATE
);