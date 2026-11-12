-- 001_seed_jobs.sql: Initial verified opportunities and candidates

INSERT INTO jobs (id, title, company, location, workplace, experience, stipend_range, department, skills, is_verified)
VALUES
('job-aura-01', 'Software Engineering Intern', 'Auraverse / Presio', 'Remote, India', 'Remote', '0+ years', '₹15,000 / month', 'Core Product', ARRAY['Go', 'React', 'PostgreSQL', 'REST APIs', 'Docker'], true),
('job-fintech-02', 'Backend Systems Intern (Go)', 'FinRoute Labs', 'Bengaluru, India', 'Hybrid', '0-1 years', '₹25,000 / month', 'Platform Engineering', ARRAY['Go', 'PostgreSQL', 'Kafka', 'Microservices'], true),
('job-cloud-03', 'Full Stack Developer Intern', 'NextScale AI', 'Remote, Global', 'Remote', '0+ years', '₹20,000 / month', 'Product Engineering', ARRAY['React', 'TypeScript', 'Go', 'PostgreSQL', 'Tailwind CSS'], true),
('job-infra-04', 'Junior Database & Backend Engineer', 'DataPulse Systems', 'Pune, India', 'Onsite', '0-1 years', '₹30,000 / month', 'Infrastructure', ARRAY['PostgreSQL', 'SQL Tuning', 'Go', 'Redis', 'Linux'], true)
ON CONFLICT (id) DO NOTHING;
