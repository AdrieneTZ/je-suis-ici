-- drop index
DROP INDEX IF EXISTS idx_checkins_location_name;
DROP INDEX IF EXISTS idx_checkins_user_did;
DROP INDEX IF EXISTS idx_checkins_created_at;
DROP INDEX IF EXISTS idx_checkins_user_id;

-- drop checkins table
DROP TABLE IF EXISTS checkins;

-- drop users table
DROP TABLE IF EXISTS users;