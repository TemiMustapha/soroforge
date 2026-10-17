-- Add ci_run_url to deployments table to record which CI pipeline run triggered
-- a deploy or upgrade. Outside CI this remains NULL.

ALTER TABLE deployments ADD COLUMN IF NOT EXISTS ci_run_url TEXT;
