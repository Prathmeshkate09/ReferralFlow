-- 001_init_schema.sql: Database schema for ReferralFlow

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(64) PRIMARY KEY,
    full_name VARCHAR(120) NOT NULL,
    email VARCHAR(180) UNIQUE NOT NULL,
    role VARCHAR(32) DEFAULT 'candidate',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS candidate_profiles (
    id VARCHAR(64) PRIMARY KEY,
    user_id VARCHAR(64) UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    phone VARCHAR(32),
    college VARCHAR(200),
    degree VARCHAR(120),
    graduation_year INT CHECK (graduation_year >= 2020),
    resume_url TEXT,
    github_url TEXT,
    linkedin_url TEXT,
    skills TEXT[] DEFAULT '{}',
    completion_score INT DEFAULT 0 CHECK (completion_score >= 0 AND completion_score <= 100),
    referral_unlocked BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS jobs (
    id VARCHAR(64) PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    company VARCHAR(150) NOT NULL,
    location VARCHAR(150) NOT NULL,
    workplace VARCHAR(50) DEFAULT 'Remote',
    experience VARCHAR(50) DEFAULT '0+ years',
    stipend_range VARCHAR(100),
    department VARCHAR(100),
    skills TEXT[] DEFAULT '{}',
    is_verified BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Composite index for fast discovery filtering on verified listings
CREATE INDEX IF NOT EXISTS idx_jobs_verified_created ON jobs(is_verified, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_jobs_workplace ON jobs(workplace);

CREATE TABLE IF NOT EXISTS referral_requests (
    id VARCHAR(64) PRIMARY KEY,
    job_id VARCHAR(64) REFERENCES jobs(id) ON DELETE CASCADE,
    candidate_id VARCHAR(64) REFERENCES candidate_profiles(id) ON DELETE CASCADE,
    pitch_note TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'SUBMITTED',
    referrer_id VARCHAR(64),
    referrer_name VARCHAR(150),
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    resolved_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_referral_candidate ON referral_requests(candidate_id);
CREATE INDEX IF NOT EXISTS idx_referral_status ON referral_requests(status);
