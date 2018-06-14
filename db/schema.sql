CREATE TABLE documents (
  id UUID NOT NULL PRIMARY KEY,
  active_version_id UUID,
  title VARCHAR,
  description TEXT,
  slug VARCHAR NOT NULL,
  is_published BOOLEAN NOT NULL DEFAULT FALSE,
  metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
  tags JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

CREATE TYPE processing_status AS ENUM (
  'not_started'
  'retrieving'
  'uploaded'
  'complete'
  'failed'
  'failed_ignored'
  'processing'
  'core_complete'
);

CREATE TYPE measurement AS ENUM (
  'inches'
);

CREATE TYPE page_dimension AS (
  width FLOAT,
  height FLOAT,
  measurement measurement
);

CREATE TABLE versions (
  id UUID NOT NULL PRIMARY KEY,
  document_id UUID NOT NULL,
  status processing_status NOT NULL DEFAULT 'not_started',
  num_pages INTEGER,
  page_dimensions page_dimension[],
  image_resolutions INTEGER[] NOT NULL DEFAULT array[]::integer[],
  html_resolutions INTEGER[] NOT NULL DEFAULT array[]::integer[],
  errors TEXT[] NOT NULL DEFAULT array[]::text[],
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

CREATE TABLE files (
  id UUID NOT NULL PRIMARY KEY,
  version_id UUID NOT NULL,
  created_from_id UUID,
  filename VARCHAR NOT NULL,
  storage_key VARCHAR NOT NULL,
  mime_type VARCHAR NOT NULL,
  size_bytes INTEGER,
  source_url VARCHAR,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);