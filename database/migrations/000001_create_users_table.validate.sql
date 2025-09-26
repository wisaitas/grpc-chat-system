-- Validation script for 000001_create_users_table
-- ตรวจสอบว่า table ถูกสร้างถูกต้องหรือไม่

-- ตรวจสอบว่า extension uuid-ossp มีอยู่
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp') THEN
        RAISE EXCEPTION 'Extension uuid-ossp does not exist';
    END IF;
END $$;

-- ตรวจสอบว่า table มีอยู่จริง
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables 
                   WHERE table_name = 'tbl_users') THEN
        RAISE EXCEPTION 'Table tbl_users does not exist';
    END IF;
END $$;

-- ตรวจสอบ columns ครบถ้วน
DO $$
BEGIN
    -- ตรวจสอบ column id
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'tbl_users' AND column_name = 'id' 
                   AND data_type = 'uuid') THEN
        RAISE EXCEPTION 'Column id with type uuid does not exist';
    END IF;
    
    -- ตรวจสอบ column email
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'tbl_users' AND column_name = 'email' 
                   AND data_type = 'character varying') THEN
        RAISE EXCEPTION 'Column email with type varchar does not exist';
    END IF;
    
    -- ตรวจสอบ column password
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'tbl_users' AND column_name = 'password' 
                   AND data_type = 'character varying') THEN
        RAISE EXCEPTION 'Column password with type varchar does not exist';
    END IF;
    
    -- ตรวจสอบ column created_at
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'tbl_users' AND column_name = 'created_at' 
                   AND data_type = 'timestamp without time zone') THEN
        RAISE EXCEPTION 'Column created_at with type timestamp does not exist';
    END IF;
    
    -- ตรวจสอบ column updated_at
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'tbl_users' AND column_name = 'updated_at' 
                   AND data_type = 'timestamp without time zone') THEN
        RAISE EXCEPTION 'Column updated_at with type timestamp does not exist';
    END IF;
    
    -- ตรวจสอบ constraints
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints 
                   WHERE table_name = 'tbl_users' AND constraint_type = 'PRIMARY KEY') THEN
        RAISE EXCEPTION 'Primary key constraint does not exist';
    END IF;
    
    -- ตรวจสอบ unique constraint บน email
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints 
                   WHERE table_name = 'tbl_users' AND constraint_type = 'UNIQUE') THEN
        RAISE EXCEPTION 'Unique constraint on email does not exist';
    END IF;
END $$;

-- ทดสอบ insert/delete เพื่อตรวจสอบ functionality
INSERT INTO tbl_users (email, password) VALUES ('test@example.com', 'testpass');
DELETE FROM tbl_users WHERE email = 'test@example.com';

SELECT 'OK' as result;
