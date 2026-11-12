import React, { useState, useEffect } from 'react'
import Navbar from './components/Navbar'
import StatsBanner from './components/StatsBanner'
import ProfileScoreCard from './components/ProfileScoreCard'
import JobFilters from './components/JobFilters'
import JobCard from './components/JobCard'
import ReferralModal from './components/ReferralModal'

export default function App() {
  const [jobs, setJobs] = useState([])
  const [search, setSearch] = useState('')
  const [workplace, setWorkplace] = useState('all')
  const [profile, setProfile] = useState(null)
  const [evaluation, setEvaluation] = useState(null)
  const [activeModalJob, setActiveModalJob] = useState(null)
  const [myReferrals, setMyReferrals] = useState([])
  const [toastMessage, setToastMessage] = useState('')

  // Initial load
  useEffect(() => {
    fetchProfile()
    fetchJobs()
    fetchReferrals()
    const interval = setInterval(() => {
      fetchReferrals()
    }, 2500)
    return () => clearInterval(interval)
  }, [])

  // Refetch when search/filter changes
  useEffect(() => {
    fetchJobs()
  }, [search, workplace])

  const fetchProfile = async () => {
    try {
      const res = await fetch('/api/profile')
      if (res.ok) {
        const json = await res.json()
        setProfile(json.data.profile)
        setEvaluation(json.data.evaluation)
      } else {
        // Fallback default
        setProfile({
          full_name: 'Prathmesh Kate',
          degree: 'B.E. AI & Data Science',
          skills: ['Go', 'React', 'PostgreSQL', 'Docker'],
        })
        setEvaluation({ score: 90, is_unlocked: true, threshold: 80, missing_items: [] })
      }
    } catch {
      setProfile({
        full_name: 'Prathmesh Kate',
        degree: 'B.E. AI & Data Science',
        skills: ['Go', 'React', 'PostgreSQL', 'Docker'],
      })
      setEvaluation({ score: 90, is_unlocked: true, threshold: 80, missing_items: [] })
    }
  }

  const fetchJobs = async () => {
    try {
      const url = `/api/jobs?q=${encodeURIComponent(search)}&workplace=${encodeURIComponent(workplace)}`
      const res = await fetch(url)
      if (res.ok) {
        const json = await res.json()
        setJobs(json.data || [])
      } else {
        setMockJobs()
      }
    } catch {
      setMockJobs()
    }
  }

  const setMockJobs = () => {
    const list = [
      {
        id: 'job-aura-01',
        title: 'Software Engineering Intern',
        company: 'Auraverse / Presio',
        location: 'Remote, India',
        workplace: 'Remote',
        experience: '0+ years',
        stipend_range: '₹15,000 / month',
        skills: ['Go', 'React', 'PostgreSQL', 'REST APIs', 'Docker'],
        is_verified: true,
        created_at: new Date().toISOString(),
      },
      {
        id: 'job-fintech-02',
        title: 'Backend Systems Intern (Go)',
        company: 'FinRoute Labs',
        location: 'Bengaluru, India',
        workplace: 'Hybrid',
        experience: '0-1 years',
        stipend_range: '₹25,000 / month',
        skills: ['Go', 'PostgreSQL', 'Kafka', 'Microservices'],
        is_verified: true,
        created_at: new Date().toISOString(),
      },
      {
        id: 'job-cloud-03',
        title: 'Full Stack Developer Intern',
        company: 'NextScale AI',
        location: 'Remote, Global',
        workplace: 'Remote',
        experience: '0+ years',
        stipend_range: '₹20,000 / month',
        skills: ['React', 'TypeScript', 'Go', 'PostgreSQL', 'Tailwind CSS'],
        is_verified: true,
        created_at: new Date().toISOString(),
      },
    ]
    setJobs(list)
  }

  const fetchReferrals = async () => {
    try {
      const res = await fetch('/api/referrals')
      if (res.ok) {
        const json = await res.json()
        setMyReferrals(json.data || [])
      }
    } catch {
      // offline
    }
  }

  const handleCreateReferral = async (jobId, pitchNote) => {
    try {
      const res = await fetch('/api/referrals', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          job_id: jobId,
          candidate_id: profile ? profile.id : 'cand-prathmesh-01',
          pitch_note: pitchNote,
        }),
      })
      const data = await res.json()
      if (res.ok) {
        setToastMessage('Referral request successfully dispatched! Tracking live status...')
        setActiveModalJob(null)
        fetchReferrals()
        setTimeout(() => setToastMessage(''), 5000)
      } else {
        alert(data.error || 'Failed to submit referral')
      }
    } catch {
      alert('Network error connecting to backend service')
    }
  }

  const isUnlocked = evaluation ? evaluation.is_unlocked : true
  const score = evaluation ? evaluation.score : 90

  return (
    <div className="min-h-screen bg-slate-50 flex flex-col justify-between">
      <div>
        <Navbar unlocked={isUnlocked} score={score} />

        <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          {toastMessage && (
            <div className="mb-6 p-4 rounded-xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-semibold flex items-center justify-between shadow-sm animate-pulse">
              <span>🚀 {toastMessage}</span>
              <button onClick={() => setToastMessage('')} className="text-emerald-600 hover:text-emerald-900 font-bold">×</button>
            </div>
          )}

          <StatsBanner />

          <ProfileScoreCard profile={profile} evaluation={evaluation} />

          {/* Active Referrals Tracker */}
          {myReferrals.length > 0 && (
            <div className="bg-white border border-slate-200 rounded-xl p-5 shadow-sm mb-6">
              <h3 className="text-sm font-bold text-slate-900 mb-3 flex items-center gap-2">
                <span>Active Referral Requests (Live Status)</span>
                <span className="bg-indigo-100 text-indigo-700 text-[10px] font-semibold px-2 py-0.5 rounded-full">
                  {myReferrals.length} active
                </span>
              </h3>
              <div className="space-y-3">
                {myReferrals.map((ref) => (
                  <div key={ref.id} className="p-3 rounded-lg border border-slate-100 bg-slate-50 flex flex-col sm:flex-row sm:items-center justify-between gap-2 text-xs">
                    <div>
                      <span className="font-semibold text-slate-900">{ref.job_title}</span>
                      <span className="text-slate-400"> @ </span>
                      <span className="font-medium text-indigo-600">{ref.company}</span>
                      <p className="text-[11px] text-slate-500 mt-0.5 truncate max-w-md">"{ref.pitch_note}"</p>
                    </div>
                    <div className="flex items-center gap-2">
                      <span className={`px-2 py-0.5 rounded font-bold text-[10px] ${
                        ref.status === 'MATCHED_TO_REFERRER' ? 'bg-emerald-100 text-emerald-800 border border-emerald-200' :
                        ref.status === 'UNDER_REVIEW' ? 'bg-sky-100 text-sky-800 border border-sky-200' :
                        'bg-amber-100 text-amber-800 border border-amber-200'
                      }`}>
                        ● {ref.status.replace(/_/g, ' ')}
                      </span>
                      {ref.referrer_name && (
                        <span className="text-[11px] text-slate-600 font-medium">({ref.referrer_name})</span>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          <JobFilters
            search={search}
            setSearch={setSearch}
            workplace={workplace}
            setWorkplace={setWorkplace}
          />

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
            {jobs.map((job) => (
              <JobCard
                key={job.id}
                job={job}
                onOpenReferralModal={setActiveModalJob}
                isUnlocked={isUnlocked}
              />
            ))}
          </div>

          {jobs.length === 0 && (
            <div className="text-center py-12 bg-white rounded-xl border border-slate-200 text-slate-500 text-sm">
              No verified opportunities found matching your filters.
            </div>
          )}
        </main>
      </div>

      {activeModalJob && (
        <ReferralModal
          job={activeModalJob}
          onClose={() => setActiveModalJob(null)}
          onSubmitReferral={handleCreateReferral}
        />
      )}

      <footer className="border-t border-slate-200 bg-white py-6 mt-12">
        <div className="max-w-7xl mx-auto px-4 text-center text-xs text-slate-500">
          ReferralFlow • High-Concurrency Job Discovery & Referral Engine • Built with Go, React, and PostgreSQL
        </div>
      </footer>
    </div>
  )
}
