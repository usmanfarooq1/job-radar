CREATE TYPE task_state_enum AS ENUM ('running', 'paused');
CREATE TYPE task_type_enum AS ENUM ('linkedin');
CREATE TYPE job_application_status_enum AS ENUM (
    'bookmarked', 'applied', 'interviewing', 'rejected', 'negotiating', 'accepted'
);
CREATE TYPE job_apply_type_enum AS ENUM ('easy_apply', 'apply');
CREATE TYPE skill_type_enum AS ENUM (
    'backend', 'frontend', 'operations', 'deployment', 'tools'
);

CREATE TABLE tasks(
    task_id UUID NOT NULL,
    search_location varchar(50) NOT NULL,
    location_id varchar(25) NOT NULL,
    delay_in_seconds int NOT NULL,
    task_state task_state_enum NOT NULL,
    task_type task_type_enum NOT NULL,
    search_keyword text NOT NULL,
    distance_radius int NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY(task_id)
);

CREATE TABLE jobs(
    job_id UUID NOT NULL,
    title text NOT NULL,
    company text NOT NULL,
    description_text text NOT NULL,
    description_text_hash text NOT NULL,
    hiring_manager_name text NULL,
    job_location varchar(50)  NULL,
    location_id varchar(25) NOT NULL,
    delay_in_seconds int NOT NULL,
    job_application_status job_application_status_enum NOT NULL,
    match_percentages json NOT NULL,
    job_apply_type job_apply_type_enum NOT NULL,
    job_link text NOT NULL,
    notes text NULL,
    created_at TIMESTAMP NOT NULL,
    applied_at TIMESTAMP NULL,
    PRIMARY KEY(job_id)
);


CREATE TABLE skills(
    skill_id UUID NOT NULL,
    skill_name text NOT NULL,
    skill_type skill_type_enum NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(skill_id)
);

CREATE INDEX idx_jobs_job_location
ON jobs (job_location);

CREATE INDEX idx_jobs_location_id
ON jobs (location_id);