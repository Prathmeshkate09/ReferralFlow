import React from 'react'

export default function JobCard({ job, onOpenReferralModal, isUnlocked }) {
  return (
    <div className="bg-white border border-slate-200 rounded-xl p-5 shadow-sm hover:shadow-md transition flex flex-col justify-between">
      <div>
        <div className="flex items-start justify-between gap-3 mb-2">
          <div>
            <div className="flex items-center gap-2">
              <h3 className="font-semibold text-slate-900 text-base">{job.title}</h3>
              {job.is_verified && (
                <span className="inline-flex items-center text-[10px] bg-emerald-50 text-emerald-700 font-semibold px-2 py-0.5 rounded-full border border-emerald-200">
                  ✓ Verified
                </span>
              )}
            </div>
            <p className="text-sm font-medium text-indigo-600">{job.company}</p>
          </div>
          <span className="text-xs font-medium bg-slate-100 text-slate-700 px-2.5 py-1 rounded-md">
            {job.workplace}
          </span>
        </div>

        <div className="flex flex-wrap items-center gap-y-1 gap-x-3 text-xs text-slate-500 mb-3">
          <span>📍 {job.location}</span>
          <span>•</span>
          <span>💼 {job.experience}</span>
          <span>•</span>
          <span>💰 {job.stipend_range}</span>
        </div>

        <div className="flex flex-wrap gap-1.5 mb-4">
          {job.skills && job.skills.map((skill, idx) => (
            <span
              key={idx}
              className="bg-slate-50 text-slate-700 border border-slate-200 text-[11px] font-medium px-2 py-0.5 rounded"
            >
              {skill}
            </span>
          ))}
        </div>
      </div>

      <div className="pt-3 border-t border-slate-100 flex items-center justify-between">
        <span className="text-[11px] text-slate-400">
          Posted {new Date(job.created_at).toLocaleDateString()}
        </span>

        <button
          onClick={() => onOpenReferralModal(job)}
          disabled={!isUnlocked}
          className={`text-xs font-semibold px-4 py-2 rounded-lg transition ${
            isUnlocked
              ? 'bg-indigo-600 hover:bg-indigo-700 text-white shadow-sm'
              : 'bg-slate-100 text-slate-400 cursor-not-allowed border border-slate-200'
          }`}
        >
          {isUnlocked ? 'Request Referral →' : 'Unlock Profile to Refer'}
        </button>
      </div>
    </div>
  )
}
