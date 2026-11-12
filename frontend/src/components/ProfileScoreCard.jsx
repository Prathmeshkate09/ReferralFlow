import React from 'react'

export default function ProfileScoreCard({ profile, evaluation }) {
  if (!profile || !evaluation) return null

  const isUnlocked = evaluation.is_unlocked
  const score = evaluation.score

  return (
    <div className="bg-white border border-slate-200 rounded-xl p-5 shadow-sm mb-6">
      <div className="flex items-center justify-between mb-3">
        <div>
          <h3 className="font-semibold text-slate-900 text-sm">Profile Completeness Gate</h3>
          <p className="text-xs text-slate-500">Reach at least 80% to unlock referral requests</p>
        </div>
        <div className="text-right">
          <span className="text-lg font-bold text-slate-900">{score}%</span>
          <span className="text-xs text-slate-400"> / 100%</span>
        </div>
      </div>

      <div className="w-full bg-slate-100 rounded-full h-2.5 overflow-hidden mb-3">
        <div
          className={`h-2.5 rounded-full transition-all duration-500 ${
            isUnlocked ? 'bg-indigo-600' : 'bg-amber-500'
          }`}
          style={{ width: `${score}%` }}
        ></div>
      </div>

      <div className="flex flex-wrap items-center justify-between text-xs text-slate-600 gap-2">
        <div className="flex items-center gap-2">
          <span className="font-medium text-slate-900">{profile.full_name}</span>
          <span className="text-slate-400">•</span>
          <span className="text-slate-500">{profile.degree}</span>
        </div>
        {isUnlocked ? (
          <span className="inline-flex items-center gap-1 text-emerald-700 bg-emerald-50 px-2 py-0.5 rounded font-medium border border-emerald-200">
            ✓ Referrals Unlocked (Ready to Request)
          </span>
        ) : (
          <span className="inline-flex items-center gap-1 text-amber-700 bg-amber-50 px-2 py-0.5 rounded font-medium border border-amber-200">
            ⚠️ Complete profile to unlock referrals
          </span>
        )}
      </div>

      {evaluation.missing_items && evaluation.missing_items.length > 0 && (
        <div className="mt-3 pt-3 border-t border-slate-100">
          <span className="text-[11px] font-medium text-slate-500 uppercase tracking-wider">Missing for 100%:</span>
          <ul className="mt-1 list-disc list-inside text-xs text-slate-600 space-y-0.5">
            {evaluation.missing_items.map((item, idx) => (
              <li key={idx}>{item}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  )
}
