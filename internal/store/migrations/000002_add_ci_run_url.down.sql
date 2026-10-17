-- Remove ci_run_url column from deployments table

ALTER TABLE deployments DROP COLUMN IF EXISTS ci_run_url;
