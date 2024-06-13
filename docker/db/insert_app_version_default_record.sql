-- Create app version table
CREATE TABLE IF NOT EXISTS app_version (
    latest_build_version BIGINT UNSIGNED NOT NULL,
    minimum_build_version BIGINT UNSIGNED NOT NULL
);

-- Check if the app_version table have no record
INSERT INTO app_version (latest_build_version, minimum_build_version)
SELECT 1, 1
FROM DUAL
WHERE NOT EXISTS (SELECT 1 FROM app_version);