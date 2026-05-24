\connect finance_db;

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('cash', 'bank')),
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT accounts_user_id_fkey 
        FOREIGN KEY (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE,

    CONSTRAINT uq_user_account_type UNIQUE (user_id, type)
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL, -- 1. TAMBAHAN: Kolom user_id untuk pengamanan & performa get mutasi
    account_id BIGINT NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense')),
    amount NUMERIC(15, 2) NOT NULL,
    category VARCHAR(50),
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- FK ke tabel users
    CONSTRAINT fk_tx_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    -- FK ke tabel accounts
    CONSTRAINT fk_tx_account
        FOREIGN KEY (account_id)
        REFERENCES accounts(id)
        ON DELETE CASCADE
);

CREATE TABLE debts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    person_name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('debt', 'loan')), 
    amount NUMERIC(15, 2) NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'unpaid' CHECK (status IN ('unpaid', 'paid')),
    account_type VARCHAR(10) CHECK (account_type IN ('cash', 'bank')),
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    paid_at TIMESTAMP,
    
    CONSTRAINT fk_debt_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- INDEXING (Dioptimalkan untuk query pencarian & filter)
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_accounts_user_id ON accounts(user_id);

-- 2. TAMBAHAN INDEX: Sangat krusial untuk fitur 'Get All Transactions By User ID' atau Riwayat Mutasi
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_created_at ON transactions(created_at DESC); -- Mempercepat sorting transaksi terbaru

CREATE INDEX idx_debts_user_id ON debts(user_id);
CREATE INDEX idx_debts_status ON debts(status);
CREATE INDEX idx_debts_type ON debts(type);